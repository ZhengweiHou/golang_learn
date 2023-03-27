package main

import (
	"database/sql"
	"fmt"

	_ "github.com/ibmdb/go_ibm_db"
	// _ "bitbucket.org/phiggins/go-db2-cli"
)

// export DB2HOME=/home/houzw/document/golang_project/pkg/mod/github.com/ibmdb/clidriver
// export IBM_DB_HOME=$DB2HOME
// export CGO_CFLAGS=-I$DB2HOME/include
// export CGO_LDFLAGS=-L$DB2HOME/lib
// export LD_LIBRARY_PATH=$DB2HOME/lib

func main() {
	// con := "HOSTNAME=localhost;DATABASE=testdb;PORT=50000;UID=db2inst1;PWD=db2inst1"
	con := "HOSTNAME=192.168.104.223;DATABASE=testdb;PORT=50000;UID=db2inst1;PWD=db2inst1"
	fmt.Println(con)
	db, err := sql.Open("go_ibm_db", con)
	// db, err := sql.Open("db2-cli", con)
	if err != nil {
		fmt.Println(err)
	}
	err = db.Ping()
	if err != nil {
		fmt.Println(err)
	}

	db.Close()
}
