# sql2md

a tool to export sql table schema to markdown table

## Install

[release pkg](https://github.com/seek4self/sql2md/releases)

## Usage

```bash
Usage of ./sql2md:
  -P int
        mysql port (default 3306)
  -h string
        mysql host (default "127.0.0.1")
  -n string
        mysql database name (default "mysql")
  -o string
        markdown output location (default ".")
  -p string
        mysql password (default "root")
  -s string
        sqlite database path
  -t string
        mysql tables, support ',' separator for filter, default all tables
  -u string
        mysql username (default "root")
  -d    show sql debug log
  -v    show version and exit
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

markdown output:

```markdown
# bmi_vdms 数据库表结构

## authorization

| 序号 | 列名          | 类型        | 主键 | 为空 | 默认值 | 额外信息       | 注释              |
| ---- | ------------- | ----------- | ---- | ---- | ------ | -------------- | ----------------- |
| 1    | `id`          | bigint(20)  | PRI  | NO   |        | auto_increment |                   |
| 2    | `code`        | varchar(32) |      | YES  |        |                | 授权码            |
| 3    | `is_use`      | tinyint(4)  |      | YES  | 0      |                | 1:注册，0：未注册 |
| 4    | `location_id` | bigint(20)  | UNI  | YES  |        |                |                   |

## casbin_rule

| 序号 | 列名    | 类型         | 主键 | 为空 | 默认值 | 额外信息       | 注释                   |
| ---- | ------- | ------------ | ---- | ---- | ------ | -------------- | ---------------------- |
| 1    | `id`    | int(11)      | PRI  | NO   |        | auto_increment |                        |
| 2    | `ptype` | varchar(100) | MUL  | YES  |        |                |                        |
| 3    | `v0`    | varchar(100) |      | YES  |        |                | 角色(id)               |
| 4    | `v1`    | varchar(100) |      | YES  |        |                | p:请求 path / g:父角色 |
| 5    | `v2`    | varchar(100) |      | YES  |        |                | 请求 method            |
| 6    | `v3`    | varchar(100) |      | YES  |        |                | 注释                   |
| 7    | `v4`    | varchar(100) |      | YES  |        |                |                        |
| 8    | `v5`    | varchar(100) |      | YES  |        |                |                        |
```

table preview:

> ## authorization
>
>| 序号 | 列名 | 类型 | 主键 | 为空 | 默认值 | 额外信息 | 注释 |
>| -- | -- | -- | -- | -- | -- | -- | -- |
>| 1 | `id` | bigint(20) | PRI | NO |  | auto_increment |  |
>| 2 | `code` | varchar(32) |  | YES |  |  | 授权码 |
>| 3 | `is_use` | tinyint(4) |  | YES | 0 |  | 1:注册，0：未注册 |
>| 4 | `location_id` | bigint(20) | UNI | YES |  |  |  |

## Todo list

[ ] add sqlite support  
[ ] add mongodb support
