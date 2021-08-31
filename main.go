/**************************************
 * @Author: mazhuang
 * @Date: 2021-08-30 14:41:41
 * @LastEditTime: 2021-08-31 10:09:47
 * @Description:
 **************************************/

package main

import (
	"flag"
	"fmt"
	"os"
	"sql2md/md"
)

var (
	host    = "127.0.0.1"
	port    = 3306
	user    = "root"
	pass    = "root"
	dbName  = "mysql"
	output  = "."
	tables  = ""
	version = false
	debug   = false
)

func init() {
	flag.StringVar(&host, "h", host, fmt.Sprintf("mysql host, default: %s", host))
	flag.IntVar(&port, "P", port, fmt.Sprintf("mysql port, default: %d", port))
	flag.StringVar(&user, "u", user, fmt.Sprintf("mysql username, default: %s", user))
	flag.StringVar(&pass, "p", pass, fmt.Sprintf("mysql password, default: %s", pass))
	flag.StringVar(&dbName, "n", dbName, fmt.Sprintf("mysql database name, default: %s", dbName))
	flag.StringVar(&tables, "t", tables, fmt.Sprintf("mysql tables, default: %s", tables))
	flag.StringVar(&output, "o", output, fmt.Sprintf("markdown output location, default: %s", dbName))
	flag.BoolVar(&version, "v", version, fmt.Sprintf("show version and exit, default: %v", version))
	flag.BoolVar(&debug, "d", debug, fmt.Sprintf("show sql debug log, default: %v", debug))
	flag.Parse()
}

func main() {
	connect()
	tables, err := findTables()
	if err != nil {
		fmt.Println("find tables err", err)
		os.Exit(1)
	}
	mdFile := md.Open(dbName, output)
	defer mdFile.Close()
	for i, t := range tables {
		fmt.Printf("%d/%d table %s making ...\n", i, len(tables), t.Name)
		columns, err := findColumns(t.Name)
		if err != nil {
			fmt.Printf("find table <%s> columns err: %v\n", t.Name, err)
			continue
		}
		tableContent := columns[0].TableHeader()
		for _, c := range columns {
			tableContent += c.TableLine()
		}
		tableContent += "\n"
		mdFile.WriteChapter(tableContent)
	}
}
