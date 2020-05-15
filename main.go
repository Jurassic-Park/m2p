package main

import (
	"fmt"
	"github.com/Jurassic-Park/m2p/core"
	"github.com/droundy/goopt"
)

var (
	sqlTable    = goopt.String([]string{"-t", "--table"}, "", "Table to build struct from")
	connString  = goopt.String([]string{"-m", "--mysql"}, "", "mysql config")
	fileDir     = goopt.String([]string{"-o", "--out-file"}, "", "file dir,if not use 'output' ")
	packageName = goopt.String([]string{"-p", "--package-name"}, "", "package ,if not use 'output' ")
)

func init() {
	goopt.Description = func() string {
		return "m2p is tool that can automaticlly generate proto file."
	}
	goopt.Version = "0.1"
	goopt.Summary = `m2p --mysql user:password@tcp\(host:port\)/database\?charset=utf8 --table tableName --out-file ./ --packageName aa`
	goopt.Parse(nil)
}

func main() {
	if *connString == "" {
		fmt.Println("mysql connect can not is empty")
		return
	}
	if *sqlTable == "" {
		fmt.Println("table can not is empty")
		return
	}
	if *packageName == "" {
		fmt.Println("package name can not is empty")
		return
	}
	core.Generator(*connString, *sqlTable, *fileDir, *packageName)
}
