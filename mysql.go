/**************************************
 * @Author: mazhuang
 * @Date: 2021-08-30 15:12:42
 * @LastEditTime: 2021-08-31 10:53:16
 * @Description:
 **************************************/

package main

import (
	"fmt"
	"sql2md/md"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type Tables struct {
	Name    string `gorm:"column:table_name"`
	Comment string `gorm:"column:table_comment"`
}

var DB *gorm.DB

func connect() {
	dsn := fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, pass, host, port, "information_schema")
	config := &gorm.Config{
		NamingStrategy:                           schema.NamingStrategy{SingularTable: true},
		DisableForeignKeyConstraintWhenMigrating: true,
	}
	db, err := gorm.Open(mysql.Open(dsn), config)
	if err != nil {
		fmt.Println("MySQL open err", err)
	}

	if debug {
		db = db.Debug()
	}
	DB = db
	fmt.Println("connect to mysql ok")
}

func findTables() (tableList []Tables, err error) {
	err = DB.Table("tables").Select("table_name, table_comment").Where("table_schema = ?", dbName).Find(&tableList).Error
	return
}

func findColumns(tableName string) (columnList []md.Columns, err error) {
	err = DB.Table("columns").Select("table_name, ordinal_position, column_default, column_type, column_name, column_key, is_nullable, extra, column_comment").
		Where("table_name = ? AND table_schema = ?", tableName, dbName).Order("ordinal_position ASC").Find(&columnList).Error
	return
}
