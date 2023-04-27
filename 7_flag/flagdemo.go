package main

import (
	"flag"
	"fmt"
)

var dbsource string        // 源数据库配置
var dbdest string          // 目标数据库配置
var sourceselectsql string // 源数据筛选sql
var desttable string       // 目标表名
var writerthreads int      // 写库线程数
var batchsize int          // 批写库时的单批数据量
var bool1 bool

func init() {
	flag.StringVar(&dbsource, "dbsource", "HOSTNAME=localhost;DATABASE=testdb;PORT=50000;UID=db2inst1;PWD=db2inst1", "源数据库配置")
	flag.StringVar(&dbdest, "dbdest", "HOSTNAME=localhost;DATABASE=testdb;PORT=50000;UID=db2inst1;PWD=db2inst1", "目标数据库配置")
	flag.StringVar(&sourceselectsql, "sourceselectsql", "select * from DBSYNCTEST.student", "源数据筛选sql")
	flag.StringVar(&desttable, "desttable", "DBSYNCTEST.student2", "目标表名")
	flag.IntVar(&writerthreads, "writerthreads", 1, "写库线程数")
	flag.IntVar(&batchsize, "batchsize", 10000, "批写库时的单批数据量")
	flag.BoolVar(&bool1, "bool1", false, "bool类型测试")
}

func main() {
	flag.Parse()

	fmt.Println("hello flag")
}
