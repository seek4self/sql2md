/**************************************
 * @Author: mazhuang
 * @Date: 2021-08-30 15:12:42
 * @LastEditTime: 2021-09-01 17:42:23
 * @Description:
 **************************************/

package sql

import (
	"fmt"
	"sql2md/md"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MySQL struct {
	db     *gorm.DB
	dbName string // Target database name
}

func newMySQL(dbName string) *MySQL {
	return &MySQL{dbName: dbName}
}

// Connect to mysql with dsn `user:pass@(host:port)`
func (m *MySQL) Connect(dsn string) (err error) {
	dsn = fmt.Sprintf("%s/%s?charset=utf8mb4&parseTime=True&loc=Local", dsn, "information_schema")
	m.db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("MySQL open err", err)
		return
	}
	fmt.Println("connect to mysql ok")
	return
}

func (m *MySQL) Debug() {
	m.db = m.db.Debug()
}

func (m *MySQL) FindTables() (tableList []Tables, err error) {
	err = m.db.Table("tables").Select("table_name, table_comment").Where("table_schema = ?", m.dbName).Find(&tableList).Error
	return
}

func (m *MySQL) FindColumns(tableName string) (columnList []md.Columns, err error) {
	err = m.db.Table("columns").Select("table_name, ordinal_position, column_default, column_type, column_name, column_key, is_nullable, extra, column_comment").
		Where("table_name = ? AND table_schema = ?", tableName, m.dbName).Order("ordinal_position ASC").Find(&columnList).Error
	return
}
