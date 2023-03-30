package main

import (
	"database/sql"
	"log"
	"strings"

	_ "github.com/ibmdb/go_ibm_db"
)

func main() {
	con := "HOSTNAME=localhost;DATABASE=testdb;PORT=50000;UID=db2inst1;PWD=db2inst1;AUTHENTICATION=SERVER;SCHEMA=DBSYNCTEST"
	db, _ := sql.Open("go_ibm_db", con)

	args := []interface{}{1}
	var sqlBuilder strings.Builder
	sqlBuilder.WriteString("select * from DBSYNCTEST.student where id in (")
	for i := 1; i < 1<<15; i++ {
		args = append(args, 1)
		sqlBuilder.WriteString("?,")
	}
	sqlBuilder.WriteString("?)")
	sqlstr := sqlBuilder.String()
	log.Println("argslen:%d sqlstr count'?':%d", len(args), strings.Count(sqlstr, "?"))
	rs, err := db.Exec(sqlstr, args...)
	if err != nil {
		log.Panic(rs, err)
	}

	db.Close()
}
