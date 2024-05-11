package db2

import (
	"database/sql"
	"fmt"

	_ "github.com/ibmdb/go_ibm_db"

	// _ "bitbucket.org/phiggins/go-db2-cli"
	"testing"
)

// export DB2HOME=/home/houzw/document/golang_project/pkg/mod/github.com/ibmdb/clidriver
// export IBM_DB_HOME=$DB2HOME
// export CGO_CFLAGS=-I$DB2HOME/include
// export CGO_LDFLAGS=-L$DB2HOME/lib
// export LD_LIBRARY_PATH=$DB2HOME/lib

func Test1(m *testing.T) {
	con := "HOSTNAME=host;DATABASE=name;PORT=number;UID=username;PWD=password"
	db, err := sql.Open("go_ibm_db", con)
	// db, err := sql.Open("db2-cli", con)
	if err != nil {
		fmt.Println(err)
	}
	db.Close()
}

func TestDb1(t *testing.T) {
	// A库连接信息
	aConnStr := "HOSTNAME=localhost;DATABASE=testdb;PORT=50001;UID=db2inst1;PWD=db2inst1;AUTHENTICATION=SERVER;CurrentSchema=DBSYNCTEST"
	// B库连接信息
	bConnStr := "HOSTNAME=localhost;DATABASE=testdb;PORT=50001;UID=db2inst1;PWD=db2inst1;AUTHENTICATION=SERVER;CurrentSchema=DBSYNCTEST"
	// bConnStr := "DATABASE=B;HOSTNAME=hostname;PORT=50000;PROTOCOL=TCPIP;UID=username;PWD=password;"

	// A库连接
	aDb, err := sql.Open("go_ibm_db", aConnStr)
	if err != nil {
		panic(err)
	}
	defer aDb.Close()

	// B库连接
	bDb, err := sql.Open("go_ibm_db", bConnStr)
	if err != nil {
		panic(err)
	}
	defer bDb.Close()

	// 查询A库T表数据
	rows, err := aDb.Query("SELECT * FROM T")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	// 插入B库T2表数据
	for rows.Next() {
		var a, b, c string
		err := rows.Scan(&a, &b, &c)
		if err != nil {
			panic(err)
		}
		_, err = bDb.Exec("INSERT INTO T2 (A, B, C) VALUES (?, ?, ?)", a, b, c)
		if err != nil {
			panic(err)
		}
	}

	fmt.Println("数据转移完成！")

}

func TestDb2(t *testing.T) {
	// A库连接信息，字符集为GBK
	aConnStr := "HOSTNAME=localhost;DATABASE=testdb;PORT=50001;UID=db2inst1;PWD=db2inst1;AUTHENTICATION=SERVER;CurrentSchema=DBSYNCTEST"
	// B库连接信息，字符集为UTF-8
	bConnStr := "HOSTNAME=localhost;DATABASE=testdb;PORT=50001;UID=db2inst1;PWD=db2inst1;AUTHENTICATION=SERVER;CurrentSchema=DBSYNCTEST"

	// A库连接
	aDb, err := sql.Open("go_ibm_db", aConnStr)
	if err != nil {
		panic(err)
	}
	defer aDb.Close()

	// B库连接
	bDb, err := sql.Open("go_ibm_db", bConnStr)
	if err != nil {
		panic(err)
	}
	defer bDb.Close()

	// 查询A库T表数据
	rows, err := aDb.Query("SELECT * FROM STUDENT")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	columns, err := rows.Columns()
	// 插入B库T2表数据
	for rows.Next() {
		var args []interface{}
		var values []string
		for i, _ := range columns {
			fmt.Println(i)
			args = append(args, new(interface{}))
		}
		err := rows.Scan(args...)
		if err != nil {
			panic(err)
		}

		// 将A库数据转换为UTF-8编码
		for _, arg := range args {
			s, ok := arg.(*interface{})
			if ok {
				values = append(values, fmt.Sprintf("%v", *s))
			}
		}

		// 插入B库T2表数据
		_, err = bDb.Exec("INSERT INTO STUDENT2 VALUES ("+GetPlaceholder(len(args))+")", values)
		// _, err = bDb.Exec("INSERT INTO STUDENT2 VALUES ("+GetPlaceholder(len(args))+")", values...)
		if err != nil {
			panic(err)
		}
	}

	fmt.Println("数据转移完成！")
}

// 获取占位符
func GetPlaceholder(n int) string {
	s := ""
	for i := 0; i < n; i++ {
		if i == 0 {
			s = "?"
		} else {
			s += ",?"
		}
	}
	return s
}

func TestGetPlaceholder(t *testing.T) {
	fmt.Println(GetPlaceholder(3))
}
