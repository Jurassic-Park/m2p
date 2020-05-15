package core

import (
	"database/sql"
	"fmt"
	"github.com/Jurassic-Park/m2p/templates"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

type SqlField struct {
	Field   string
	Type    string
	Null    string
	Key     string
	Default sql.NullString
	Extra   string
}

func GetDB(connString string) *sql.DB {
	var err error
	db, err := sql.Open("mysql", connString)
	db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(50)
	if err != nil {
		fmt.Printf("connect mysql fail ! [%s]", err)
	} else {
		fmt.Println("connect to mysql success")
	}
	return db
}

//获取mysql结构的slic
func GetMysqlStruct(connString string, tableName string) ([]SqlField, error) {
	var slic = make([]SqlField, 0)
	var sqlField = new(SqlField)
	db := GetDB(connString)
	defer db.Close()
	rows, err := db.Query("desc " + tableName)
	defer rows.Close()
	if err != nil {
		return slic, err
	}
	if err != nil {
		return slic, err
	}
	for rows.Next() {
		//定义变量接收查询数据
		err := rows.Scan(&sqlField.Field, &sqlField.Type, &sqlField.Null, &sqlField.Key, &sqlField.Default, &sqlField.Extra)
		if err != nil {
			return slic, err
		}
		slic = append(slic, *sqlField)
	}
	return slic, err
}

//处理名称
func DealServerName(tableName string) string {
	var slic = make([]string, 0)
	var serverName string
	slic = strings.Split(tableName, "_")
	for i := 0; i < len(slic); i++ {
		serverName += FirstToUpper(slic[i])
	}
	return serverName
}

//首字母大写其他小写
func FirstToUpper(s string) string {
	return strings.ToUpper(s[0:1]) + strings.ToLower(s[1:])
}

//生成proto文件
func Generator(connString string, tableName string, fileDir string, packageName string) {
	// 大驼峰表名
	parts := strings.Split(tableName, "_")
	UCamelTableName := ""
	SCamelTableName := ""
	for k, v := range parts {
		if k != 0 {
			SCamelTableName += FirstToUpper(v)
		} else {
			SCamelTableName += v
		}
		UCamelTableName += FirstToUpper(v)
	}

	var fileString = templates.ProtoTpl
	//获取mysql结构
	fieldSlic, err := GetMysqlStruct(connString, tableName)
	if err != nil {
		fmt.Println(err.Error())
	}
	//整理参数
	// format := map[string]string{
	// 	"{{TableSchema}}": ConvertMysqlTypeToProtoType(fieldSlic),
	// 	"{{TableName}}":   tableName,
	// 	"{{ServerName}}":  DealServerName(tableName),
	// }
	format := map[string]string{
		"{{TableSchema}}":     ConvertMysqlTypeToProtoType(fieldSlic),
		"{{UCamelTableName}}": UCamelTableName,
		"{{SCamelTableName}}": SCamelTableName,
		"{{PackageName}}":     packageName,
	}
	//替换关键字
	for k, v := range format {
		fileString = strings.ReplaceAll(fileString, k, v)
	}
	//生成文件
	// if fileDir != "" {
	// 	fileDir = filepath.Join(fileDir, "protos", tableName)
	// } else {
	// 	fileDir = filepath.Join("output", "protos", tableName)
	// }
	// 当前有相同文件不更新
	fileName := tableName + ".proto"
	if ok, err := PathExists(fileDir + "/" + fileName); err == nil && ok {
		fmt.Println("目录下存在相同文件:" + fileDir + "/" + fileName)
		return
	}
	WriteFile(fileDir, fileName, fileString, 0755)
}
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

//convert mysql type to proto type
func ConvertMysqlTypeToProtoType(fieldSlic []SqlField) string {
	var schema string
	//处理变量
	for i := 0; i < len(fieldSlic); i++ {
		numStr := strconv.Itoa(i + 1)
		if strings.Index(fieldSlic[i].Type, "bigint") > -1 {
			if strings.Index(fieldSlic[i].Type, "unsigned") > -1 {
				schema += "    uint64 "
			} else {
				schema += "    int64 "
			}
		} else if strings.Index(fieldSlic[i].Type, "int") > -1 {
			if strings.Index(fieldSlic[i].Type, "unsigned") > -1 {
				schema += "    uint32 "
			} else {
				schema += "    int32 "
			}
		} else if strings.Index(fieldSlic[i].Type, "text") > -1 {
			schema += "    string "
		} else if strings.Index(fieldSlic[i].Type, "char") > -1 {
			schema += "    string "
		} else if strings.Index(fieldSlic[i].Type, "enum") > -1 {
			schema += "    string "
		} else if strings.Index(fieldSlic[i].Type, "blob") > -1 {
			schema += "    string "
		} else if strings.Index(fieldSlic[i].Type, "float") > -1 {
			schema += "    float "
		} else if strings.Index(fieldSlic[i].Type, "double") > -1 {
			schema += "    double "
		} else if strings.Index(fieldSlic[i].Type, "date") > -1 {
			schema += "    string "
		} else if strings.Index(fieldSlic[i].Type, "time") > -1 {
			schema += "    string "
		} else {
			schema += "    string "
		}
		schema += fieldSlic[i].Field + " = " + numStr + ";"
		if i < len(fieldSlic)-1 {
			schema += "\n"
		}
	}
	return schema
}

//写入文件
func WriteFile(fileDir string, fileName string, file string, mode os.FileMode) error {
	_, err := os.Stat(fileDir)
	if err != nil {
		err = os.MkdirAll(fileDir, mode)
		if err != nil {
			log.Fatalln(err.Error() + ": " + fileDir)
		}
	}
	fn := filepath.Join(fileDir, fileName)
	err = ioutil.WriteFile(fn, []byte(file), mode)
	if err != nil {
		log.Fatalln(err.Error() + ": " + fn)
	}
	fmt.Println("success create :" + fn)
	return err
}
