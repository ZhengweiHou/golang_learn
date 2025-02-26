package db2

import (
	"database/sql"
	"fmt"
	"testing"
	"time"

	"github.com/ibmdb/go_ibm_db"
)

// db2 幻读测试
/*
CREATE TABLE "TEST"."PHANTOM_TEST"  (
  "ID" INTEGER NOT NULL GENERATED BY DEFAULT AS IDENTITY,
  "NAME" VARCHAR(100) NOT NULL ,
  "AGE" INTEGER ,
  "MODIFIED" TIMESTAMP NOT NULL GENERATED BY DEFAULT FOR EACH ROW ON UPDATE AS ROW CHANGE TIMESTAMP,
  PRIMARY KEY(ID)
  );

CREATE INDEX IDX_MODIFIED ON "TEST"."PHANTOM_TEST" ("MODIFIED");
*/
func TestPhantomRead1(t *testing.T) {
	dataSourceName := "HOSTNAME=localhost;DATABASE=testdb;PORT=50003;UID=db2inst1;PWD=db2inst1;AUTHENTICATION=SERVER;CurrentSchema=TEST;DB2CODEPAGE=1208"
	db, _ := sql.Open("go_ibm_db", dataSourceName)
	defer db.Close()

	var size int
	var id int
	var name string
	var age int
	var modified time.Time
	for i := 0; i < 10; i++ {
		size = 0
		rows, _ := db.Query("select id,name,age,modified from PHANTOM_TEST WHERE MODIFIED >'2024-01-28' with ur")
		for rows.Next() {
			size++
			rows.Scan(&id, &name, &age, &modified)
			time.Sleep(time.Microsecond * 10)
			// fmt.Printf("id:%d, name:%s, age:%d modified:%v\n", id, name, age, modified)
		}
		fmt.Printf("sizw:%d\n", size) // TODO 幻读未复现
		time.Sleep(time.Second)
	}
}

func TestPRMergetest1(t *testing.T) {
	mergeSQL := `MERGE INTO TEST.PHANTOM_TEST d USING (select ID, NAME, AGE, MODIFIED FROM TEST.PT_2) t 
	ON (d.ID = t.ID AND d.NAME = t.NAME) 
	WHEN MATCHED THEN 
	UPDATE SET d.NAME = t.NAME, d.AGE = t.AGE, d.MODIFIED = t.MODIFIED 
	WHEN NOT MATCHED THEN 
	INSERT (d.ID, d.NAME, d.AGE, d.MODIFIED) VALUES (t.ID, t.NAME, t.AGE, t.MODIFIED)`

	// delDupSQL := `DELETE FROM TEST.PT_2 t
	// WHERE ROWID <> (
    // SELECT MIN(ROWID)
    // FROM TEST.PT_2 AS t2
    // WHERE t2.ID = t.ID AND t2.NAME = t.NAME
    // )`

	delDupSQL := `
	DELETE FROM (
		SELECT 1 FROM (
			SELECT 1,ROW_NUMBER() OVER (PARTITION BY id,name) rn FROM TEST.PT_2
		) WHERE rn>1
	)
	`

	dataSourceName := "HOSTNAME=localhost;DATABASE=testdb;PORT=50003;UID=db2inst1;PWD=db2inst1;AUTHENTICATION=SERVER;CurrentSchema=TEST;DB2CODEPAGE=1208"
	db, _ := sql.Open("go_ibm_db", dataSourceName)
	defer db.Close()

	result, err := db.Exec(mergeSQL)
	if err != nil {

		db2stat788 := false
		if dbErr, ok := err.(*go_ibm_db.Error); ok {
			if dbErr.Diag != nil {
				for _, diag := range dbErr.Diag {
					if diag.NativeError == -788 {
						fmt.Printf("State:%v\n", diag.NativeError)
						db2stat788 = true // merge时发生主键重复错误
						break
					}
				}
			}
		}

		if db2stat788 {
			result, _ = db.Exec(delDupSQL)

			delaf, _ := result.RowsAffected()
			fmt.Printf("del temp affect size %d\n", delaf)
			result, err = db.Exec(mergeSQL)
			if err != nil {
				t.Fatal(err)
			}
		}
	}
	af, err := result.RowsAffected()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("affect:%d\n", af)
}
