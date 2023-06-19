package main

import (
	"database/sql"
	"fmt"
	"log"
	"testing"

	"github.com/axgle/mahonia"
	_ "github.com/ibmdb/go_ibm_db"
	"github.com/saintfish/chardet"
)

func TestDb2gbktest2(t *testing.T) {

	// db - gbk 需设置环境变量 "DB2CODEPAGE=1386"
	dataSourceName := "HOSTNAME=localhost;DATABASE=testdb;PORT=50003;UID=db2inst1;PWD=db2inst1;AUTHENTICATION=SERVER;CurrenTSchema=DBSYNCTEST;DB2CODEPAGE=1386"

	// var db *sql.DB
	// var stmt *sql.Stmt
	// var rows *sql.Rows
	var err error

	db, err := sql.Open("go_ibm_db", dataSourceName)

	stmt, err := db.Prepare("select name from student")
	if err != nil {
		t.Fatal(err)
	}
	rows, _ := stmt.Query()

	// rows, err = db.Query("select name from student where id>?", 0)

	if err != nil {
		t.Fatal(err)
	}
	defer rows.Close()

	detector := chardet.NewTextDetector()

	for rows.Next() {

		count := 1
		var rowValues = make([]interface{}, count)
		var valuePointers = make([]interface{}, count)
		for i := range rowValues {
			valuePointers[i] = &rowValues[i]
		}
		err := rows.Scan(valuePointers...)
		if err != nil {
			log.Panic(err)
		}

		// enc := mahonia.NewDecoder("gbk")
		for i := range rowValues {
			var value interface{}
			rawValue := rowValues[i]
			b, ok := rawValue.([]byte)
			if ok {

				// 自动识别编码并做读取
				result, _ := detector.DetectBest(b)
				srcDecoder := mahonia.NewDecoder(result.Charset)
				value = srcDecoder.ConvertString(string(b))

				// value = string(b)
				// value = enc.ConvertString(string(b))
				// fmt.Printf("charset:%v, len:%v ", result.Charset, len(b))
				fmt.Printf("len:%v ", len(b))
			} else {
				value = rawValue
			}
			rowValues[i] = value
		}

		fmt.Println("===", rowValues)
	}
}
