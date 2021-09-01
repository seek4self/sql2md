/**************************************
 * @Author: mazhuang
 * @Date: 2021-08-30 16:33:08
 * @LastEditTime: 2021-09-01 11:18:51
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
	path  string
	f     *os.File
}

func Open(name, location string) (md *MD) {
	md = &MD{
		path:  location + string(os.PathSeparator) + name + ".md",
		title: name + ` 数据库表结构`,
	}
	fmt.Printf("open markdown file ...\n")
	mkdir(location)
	mdFile, err := os.Create(md.path)
	if err != nil {
		fmt.Println("create markdown file err:", err)
		os.Exit(1)
	}
	md.f = mdFile
	return
}

func (md *MD) WriteHeader() {
	fmt.Printf("write markdown header ...\n")
	md.WriteChapter(H1 + md.title + "\n\n")
}

func (md *MD) WriteChapter(chapter string) {
	if _, err := md.f.WriteString(chapter); err != nil {
		fmt.Println("write chapter err:", err)
	}
}

func (md *MD) Close() {
	md.f.Close()
	fmt.Printf("write %s done.\n", md.path)
}

func mkdir(dir string) {
	_, err := os.Stat(dir)
	if err == nil {
		return
	}
	if os.IsNotExist(err) {
		err = os.Mkdir(dir, 0776)
	}
	if err == nil {
		return
	}
	fmt.Printf("mkdir %s err %v\n", dir, err)
	os.Exit(1)
}
