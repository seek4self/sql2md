/**************************************
 * @Author: mazhuang
 * @Date: 2021-09-01 11:32:15
 * @LastEditTime: 2021-09-01 17:31:51
 * @Description:
 **************************************/

package sql

import "sql2md/md"

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
