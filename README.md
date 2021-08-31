# sql2md

a tool to export sql table schema to markdown table

## Usage

```bash
Usage of ./sql2md:
  -h string
        mysql host, default: 127.0.0.1 (default "127.0.0.1")
  -P int
        mysql port, default: 3306 (default 3306)
  -n string
        mysql database name, default: mysql (default "mysql")
  -o string
        markdown output location, default: mysql (default ".")
  -p string
        mysql password, default: root (default "root")
  -t string
        mysql tables, support ',' separator for filter, default all tables
  -u string
        mysql username, default: root (default "root")
  -v    show version and exit, default: false
  -d    show sql debug log, default: false
```

example:

```bash
$ ./sql2md -p 123456 -n bmi_vdms 
connect to mysql ok
open markdown file
write markdown header ...
1/5 creating table authorization ...
2/5 creating table casbin_rule ...
3/5 creating table d_video ...
4/5 creating table s_video ...
5/5 creating table video ...
write md done.
```

## Todo list

[ ] add sqlite support  
[ ] add mongodb support
