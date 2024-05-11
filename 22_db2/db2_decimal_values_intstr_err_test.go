package db2

import (
	"database/sql"
	"fmt"

	_ "github.com/ibmdb/go_ibm_db"

	"testing"
)

/*
CREATE TABLE TEST.TB1  (

	"DECIMAL1" DECIMAL(10,4)

)
*/
func TestDecimalStrValueErr(m *testing.T) {
	con := "HOSTNAME=localhost;DATABASE=testdb;PORT=50003;UID=db2inst1;PWD=db2inst1;AUTHENTICATION=SERVER;CurrentSchema=TEST"
	db, _ := sql.Open("go_ibm_db", con)
	sqlstr := "INSERT INTO TB1 (DECIMAL1) VALUES(?)"

	dbExec(db, sqlstr, []any{"1"})
	dbExec(db, sqlstr, []any{"111"})
	dbExec(db, sqlstr, []any{"1111"})
	dbExec(db, sqlstr, []any{"1.0"})
	dbExec(db, sqlstr, []any{"1.1"})
	dbExec(db, sqlstr, []any{"1.00"})
	dbExec(db, sqlstr, []any{1})
	dbExec(db, sqlstr, []any{"1.000"})
	dbExec(db, sqlstr, []any{"22e-1"})
	dbExec(db, sqlstr, []any{"2e1"})
	dbExec(db, sqlstr, []any{nil})
	dbExec(db, sqlstr, []any{"0.11111"})

	db.Close()
}

func dbExec(db *sql.DB, sqlstr string, args []any) {
	rs, err := db.Exec(sqlstr, args...)
	if err != nil {
		fmt.Printf("%-10v %s\n", args, err.Error())
	} else {
		affect, _ := rs.RowsAffected()
		fmt.Printf("%-10v affect:%d\n", args, affect)
	}
}
