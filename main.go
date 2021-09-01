/**************************************
 * @Author: mazhuang
 * @Date: 2021-08-30 14:41:41
 * @LastEditTime: 2021-09-01 18:09:11
 * @Description:
 **************************************/

package main

import (
	"flag"
	"fmt"
	"os"
	"path"
	"strings"

	"sql2md/md"
	"sql2md/sql"
)

var (
	host    = "127.0.0.1"
	port    = 3306
	user    = "root"
	pass    = "root"
	dbName  = "mysql"
	output  = "."
	tables  = ""
	sqlite  = ""
	version = false
	debug   = false
)

func init() {
	flag.StringVar(&host, "h", host, "mysql host")
	flag.IntVar(&port, "P", port, "mysql port")
	flag.StringVar(&user, "u", user, "mysql username")
	flag.StringVar(&pass, "p", pass, "mysql password")
	flag.StringVar(&dbName, "n", dbName, "mysql database name")
	flag.StringVar(&tables, "t", tables, "mysql tables, support ',' separator for filter, default all tables")
	flag.StringVar(&sqlite, "s", sqlite, "sqlite database path")
	flag.StringVar(&output, "o", output, "markdown output location")
	flag.BoolVar(&version, "v", version, "show version and exit")
	flag.BoolVar(&debug, "d", debug, "show sql debug log")
	flag.Parse()
}

func main() {
	fmt.Println("sql2md version v1.0.1")
	if version {
		return
	}
	var db sql.DB
	dsn := sqlite
	if sqlite != "" {
		_, dbName = path.Split(sqlite)
		db = sql.NewSQL("sqlite", dbName)
	} else {
		db = sql.NewSQL("mysql", dbName)
		dsn = fmt.Sprintf("%s:%s@(%s:%d)", user, pass, host, port)
	}
	if err := db.Connect(dsn); err != nil {
		return
	}
	if debug {
		db.Debug()
	}

	var (
		tableList    []sql.Tables
		tableContent string
		err          error
	)
	if tables != "" {
		tableNames := strings.Split(tables, ",")
		for _, n := range tableNames {
			tableList = append(tableList, sql.Tables{Name: n})
		}
	} else {
		tableList, err = db.FindTables()
		if err != nil {
			fmt.Println("find tables err", err)
			os.Exit(1)
		}
	}

	mdFile := md.Open(dbName, output)
	defer mdFile.Close()
	mdFile.WriteHeader()
	for i, t := range tableList {
		fmt.Printf("%d/%d creating table %s ...\n", i+1, len(tableList), t.Name)
		columns, err := db.FindColumns(t.Name)
		if err != nil {
			fmt.Printf("find table <%s> columns err: %v\n", t.Name, err)
			continue
		}
		tableContent = columns[0].TableHeader(t.Comment)
		for _, c := range columns {
			tableContent += c.TableLine()
		}
		tableContent += "\n"
		mdFile.WriteChapter(tableContent)
	}
}
