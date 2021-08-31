/**************************************
 * @Author: mazhuang
 * @Date: 2021-08-30 16:33:08
 * @LastEditTime: 2021-08-31 10:15:34
 * @Description:
 **************************************/

package md

import (
	"fmt"
	"os"
)

const (
	H1 = "# "
	H2 = "## "
	H3 = "### "
	H4 = "#### "

	TableLineEnd = "|\n"
)

type MD struct {
	title string
	name  string
	f     *os.File
}

func Open(name, location string) (md *MD) {
	md = &MD{
		name:  location + string(os.PathSeparator) + name + ".md",
		title: name + ` 数据库表结构`,
	}
	mdFile, err := os.Create(md.name)
	if err != nil {
		fmt.Println("creat markdown file err:", err)
		os.Exit(1)
	}
	fmt.Printf("open markdown file\n")
	md.f = mdFile
	fmt.Printf("write markdown header ...\n")
	md.WriteChapter(H1 + md.title + "\n\n")
	return
}

func (md *MD) WriteChapter(chapter string) {
	if _, err := md.f.WriteString(chapter); err != nil {
		fmt.Println("write chapter err:", err)
	}
}

func (md *MD) Close() {
	md.f.Close()
	fmt.Println("write md done.")
}
