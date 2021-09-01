/**************************************
 * @Author: mazhuang
 * @Date: 2021-09-01 16:19:16
 * @LastEditTime: 2021-09-01 18:36:08
 * @Description:
 **************************************/

package sql

import (
	"database/sql"
	"fmt"
	"sql2md/md"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type SQLite struct {
	db *gorm.DB
}

func newSQLite() *SQLite {
	return &SQLite{}
}

func (s *SQLite) Connect(dsn string) (err error) {
	s.db, err = gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("SQLite open err", err)
		return
	}
	fmt.Println("connect to SQLite ok")
	return
}

func (s *SQLite) Debug() {
	s.db = s.db.Debug()
}

func (s *SQLite) FindTables() (ts []Tables, err error) {
	var names []string
	err = s.db.Table("sqlite_master").Where("type = 'table'").Pluck("tbl_name", &names).Error
	ts = NewTables(names)
	return
}

func (s *SQLite) FindColumns(tableName string) (cs []md.Columns, err error) {
	var tableInfos []TableInfo
	err = s.db.Raw("PRAGMA table_info('" + tableName + "');").Scan(&tableInfos).Error
	cs = make([]md.Columns, len(tableInfos))
	for i, t := range tableInfos {
		cs[i] = md.Columns{
			OrdinalPosition: t.Cid,
			TName:           tableName,
			Name:            t.Name,
			Type:            sql.NullString{String: t.Type},
			Key:             t.Key(),
			IsNull:          t.IsNull(),
			Default:         sql.NullString{String: t.DfltValue},
		}
	}
	return
}

type TableInfo struct {
	Cid       int
	Name      string
	Type      string
	Notnull   int
	DfltValue string
	PK        int
}

func (t TableInfo) Key() string {
	if t.PK == 1 {
		return "PRI"
	}
	return ""
}

func (t TableInfo) IsNull() string {
	if t.Notnull == 1 {
		return "No"
	}
	return "YES"
}
