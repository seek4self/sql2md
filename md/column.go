/**************************************
 * @Author: mazhuang
 * @Date: 2021-08-30 17:45:01
 * @LastEditTime: 2021-08-31 10:04:22
 * @Description:
 **************************************/

package md

import (
	"database/sql"
	"fmt"
	"reflect"
	"strings"
)

type Columns struct {
	TName           string         `gorm:"column:table_name" md:"head"`
	OrdinalPosition int            `gorm:"column:ordinal_position" md:"序号"`
	Name            string         `gorm:"column:column_name" md:"列名"`
	Type            sql.NullString `gorm:"column:column_type" md:"类型"`
	Key             string         `gorm:"column:column_key" md:"主键"`
	IsNull          string         `gorm:"column:is_nullable" md:"为空"`
	Default         sql.NullString `gorm:"column:column_default" md:"默认值"`
	Extra           sql.NullString `gorm:"column:extra" md:"额外信息"`
	Comment         string         `gorm:"column:column_comment" md:"注释"`
}

func (c Columns) TableHeader() (header string) {
	t := reflect.TypeOf(c)
	v := reflect.ValueOf(c)
	splitLine := ""
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		tag := field.Tag.Get("md")
		if tag == "head" {
			header += H2 + v.Field(i).String() + "\n\n"
			continue
		}
		header += "| " + tag + " "
		splitLine += "| -- "
	}
	header += TableLineEnd
	splitLine += TableLineEnd
	header += splitLine
	return header
}

func (c Columns) TableLine() string {
	return fmt.Sprintf("| %d | `%s` | %s | %s | %s | %s | %s | %s |\n",
		c.OrdinalPosition,
		c.Name,
		c.Type.String,
		c.Key,
		c.IsNull,
		c.Default.String,
		c.Extra.String,
		strings.ReplaceAll(strings.ReplaceAll(c.Comment, "|", "\\|"), "\n", " "),
	)
}
