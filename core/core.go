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

type SqlFieldDesc struct {
	COLUMN_NAME    string
	COLUMN_COMMENT string
	COLUMN_TYPE    string
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

func GetBetweenStr(str, start, end string) string {
	n := strings.Index(str, start)
	n += len(start)
	if n == -1 {
		n = 0
	}
	str = string([]byte(str)[n:])
	m := strings.Index(str, end)
	if m == -1 {
		m = len(str)
	}
	str = string([]byte(str)[:m])
	return str
}

//获取mysql结构的slic
func GetMysqlStruct(connString string, tableName string) ([]SqlFieldDesc, error) {
	var slic = make([]SqlFieldDesc, 0)
	var sqlFieldDesc = new(SqlFieldDesc)
	db := GetDB(connString)
	defer db.Close()

	tableSchema := GetBetweenStr(connString, ")/", "?")
	fmt.Println("table schema is " + tableSchema)

	querySql := "select COLUMN_NAME, COLUMN_COMMENT, COLUMN_TYPE from information_schema.columns where table_schema ='" + tableSchema + "' and table_name = '" + tableName + "' ;"
	rows, err := db.Query(querySql)
	defer rows.Close()
	if err != nil {
		return slic, err
	}
	if err != nil {
		return slic, err
	}
	for rows.Next() {
		//定义变量接收查询数据
		err := rows.Scan(&sqlFieldDesc.COLUMN_NAME, &sqlFieldDesc.COLUMN_COMMENT, &sqlFieldDesc.COLUMN_TYPE)
		if err != nil {
			return slic, err
		}
		slic = append(slic, *sqlFieldDesc)
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

// 生成驼峰
// tag 大驼峰 1， 小驼峰 0
func generatorCamelName(str string, tag int) (name string) {
	parts := strings.Split(str, "_")
	if tag == 1 {
		for _, v := range parts {
			name += FirstToUpper(v)
		}
	} else {
		for k, v := range parts {
			if k != 0 {
				name += FirstToUpper(v)
			} else {
				name += v
			}
		}
	}
	return
}

//生成proto文件
func Generator(connString string, tableName string, fileDir string, packageName string) {
	// 大驼峰表名
	UCamelTableName := generatorCamelName(tableName, 1)
	SCamelTableName := generatorCamelName(tableName, 0)

	var fileString = templates.ProtoTpl
	//获取mysql结构
	fieldSlic, err := GetMysqlStruct(connString, tableName)
	if err != nil {
		fmt.Println(err.Error())
	}
	//整理参数
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
func ConvertMysqlTypeToProtoType(fieldSlic []SqlFieldDesc) string {
	var schema string
	//处理变量
	for i := 0; i < len(fieldSlic); i++ {
		numStr := strconv.Itoa(i + 1)
		if strings.Index(fieldSlic[i].COLUMN_TYPE, "bigint") > -1 {
			if strings.Index(fieldSlic[i].COLUMN_TYPE, "unsigned") > -1 {
				schema += "    uint64 "
			} else {
				schema += "    int64 "
			}
		} else if strings.Index(fieldSlic[i].COLUMN_TYPE, "int") > -1 {
			if strings.Index(fieldSlic[i].COLUMN_TYPE, "unsigned") > -1 {
				schema += "    uint32 "
			} else {
				schema += "    int32 "
			}
		} else if strings.Index(fieldSlic[i].COLUMN_TYPE, "text") > -1 {
			schema += "    string "
		} else if strings.Index(fieldSlic[i].COLUMN_TYPE, "char") > -1 {
			schema += "    string "
		} else if strings.Index(fieldSlic[i].COLUMN_TYPE, "enum") > -1 {
			schema += "    string "
		} else if strings.Index(fieldSlic[i].COLUMN_TYPE, "blob") > -1 {
			schema += "    string "
		} else if strings.Index(fieldSlic[i].COLUMN_TYPE, "float") > -1 {
			schema += "    float "
		} else if strings.Index(fieldSlic[i].COLUMN_TYPE, "double") > -1 {
			schema += "    double "
		} else if strings.Index(fieldSlic[i].COLUMN_TYPE, "date") > -1 {
			schema += "    string "
		} else if strings.Index(fieldSlic[i].COLUMN_TYPE, "time") > -1 {
			schema += "    string "
		} else {
			schema += "    string "
		}
		schema += generatorCamelName(fieldSlic[i].COLUMN_NAME, 0) + " = " + numStr + "; // " + fieldSlic[i].COLUMN_COMMENT
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
