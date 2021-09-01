/**************************************
 * @Author: mazhuang
 * @Date: 2021-09-01 11:32:15
 * @LastEditTime: 2021-09-01 18:35:13
 * @Description:
 **************************************/

package sql

import (
	"sql2md/md"
)

type Tables struct {
	Name    string `gorm:"column:table_name"`
	Comment string `gorm:"column:table_comment"`
}

type DB interface {
	Connect(dsn string) error
	Debug()
	FindTables() ([]Tables, error)
	FindColumns(tableName string) ([]md.Columns, error)
}

func NewSQL(typ, dbName string) DB {
	switch typ {
	case "sqlite":
		return newSQLite()
	case "mysql":
		return newMySQL(dbName)
	}
	return nil
}

func NewTables(names []string) []Tables {
	ts := make([]Tables, len(names))
	for i, n := range names {
		ts[i] = Tables{Name: n}
	}
	return ts
}
