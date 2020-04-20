package main

import (
	"fmt"
	"m2p/core"

	"github.com/droundy/goopt"
)

var (
	sqlTable   = goopt.String([]string{"-t", "--table"}, "", "Table to build struct from")
	connString = goopt.String([]string{"-m", "--mysql"}, "", "mysql config")
	fileDir    = goopt.String([]string{"-o", "--out-file"}, "", "file dir,if not use 'output' ")
)

func init() {
	goopt.Description = func() string {
		return "m2p is tool that can automaticlly generate proto file."
	}
	goopt.Version = "0.1"
	goopt.Summary = `m2p --mysql user:password@tcp\(host:port\)/database\?charset=utf8 --table tableName --out-file ./`
	goopt.Parse(nil)
}

func main() {
	if *connString == "" {
		fmt.Println("mysql can not is empty")
		return
	}
	if *sqlTable != "" {
		core.Generator(*sqlTable, *fileDir)
	} else {
		fmt.Println("table can not is empty")
	}
}
