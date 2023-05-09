package main

import (
	"database/sql"
	"fmt"
	"log"
	"testing"
	"time"

	_ "github.com/ibmdb/go_ibm_db"
	"github.com/shopspring/decimal"
	// _ "bitbucket.org/phiggins/go-db2-cli"
)

// export DB2HOME=/home/houzw/document/golang_project/pkg/mod/github.com/ibmdb/clidriver
// export IBM_DB_HOME=$DB2HOME
// export CGO_CFLAGS=-I$DB2HOME/include
// export CGO_LDFLAGS=-L$DB2HOME/lib
// export LD_LIBRARY_PATH=$DB2HOME/lib

func TestMain1(t *testing.T) {

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

func TestMain2(t *testing.T) {
	con := "HOSTNAME=localhost;DATABASE=testdb;PORT=50001;UID=db2inst1;PWD=db2inst1;AUTHENTICATION=SERVER;CurrentSchema=DBSYNCTEST"
	db, _ := sql.Open("go_ibm_db", con)

	var record = make([]interface{}, 0)

	sql := "INSERT INTO STUDENT2(NAME,AGE,GRADES,FEE,MODIFIED,ID) VALUES('111',1,1.1,?,?,111)"
	record = append(record, 1.1, time.Now())

	result, err := db.Exec(sql, record...)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("insert result:%v\n", result)
}

func TestMain3(t *testing.T) {
	con := "HOSTNAME=localhost;DATABASE=testdb;PORT=50001;UID=db2inst1;PWD=db2inst1;AUTHENTICATION=SERVER;CurrentSchema=DBSYNCTEST"
	db, _ := sql.Open("go_ibm_db", con)

	sql := "INSERT INTO STUDENT2(NAME,AGE,GRADES,FEE,MODIFIED,ID) VALUES('111',1,?,11.11,?,111)"
	stmt, err := db.Prepare(sql)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(decimal.NewFromFloat(1.10), time.Now())
	if err != nil {
		log.Fatal(err)
	}
}

func TestMain4(t *testing.T) {
	con := "HOSTNAME=localhost;DATABASE=testdb;PORT=50001;UID=db2inst1;PWD=db2inst1;AUTHENTICATION=SERVER;CurrentSchema=DBSYNCTEST"
	db, _ := sql.Open("go_ibm_db", con)
	// var d, _ = decimal.NewFromString("1.10")

	stmt, err := db.Prepare("INSERT INTO STUDENT2(NAME,AGE,GRADES,FEE,MODIFIED,ID) VALUES('111',1,?,?,?,111)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(1.11, 11.11, time.Now())
	if err != nil {
		log.Fatal(err)
	}

}
