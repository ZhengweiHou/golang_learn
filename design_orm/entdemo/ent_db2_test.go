package main

import (
	"context"
	"database/sql"
	"entdemo/hzwent"
	"entdemo/hzwent/student"
	"fmt"
	"log"
	"testing"
	"time"

	_ "github.com/ibmdb/go_ibm_db"
)

func TestEnt(t *testing.T) {
	ctx := context.Background()
	drivername := "go_ibm_db"
	con := "HOSTNAME=localhost;DATABASE=testdb;PORT=50001;UID=db2inst1;PWD=db2inst1;AUTHENTICATION=SERVER;CurrentSchema=ENTTEST"
	client, err := hzwent.Open(drivername, con)
	if err != nil {
		log.Fatalf("failed opening connection to %s: %v", drivername, err)
	}
	defer client.Close()

	// Run the auto migration tool.
	//	if err := client.Schema.Create(ctx); err != nil {
	//		log.Fatalf("failed creating schema resources: %v", err)
	//	}
	//
	//	stu1, err := CreateStudent(ctx, client)
	//	fmt.Printf("stu:%v", stu1)

	// insert
	stu1, err := CreateStudent(ctx, client)
	if err != nil {
		log.Fatalf("failed create: %v", err)
	}
	fmt.Printf("stu:%v", stu1)

	// query
	stu, err := QueryStudent(ctx, client)
	if err != nil {
		log.Fatalf("failed query: %v", err)
	}
	fmt.Printf("stu:%v\n", stu)

	// update

}

func TestEntDb2(t *testing.T) {
	ctx := context.Background()
	drivername := "go_ibm_db"
	con := "HOSTNAME=localhost;DATABASE=testdb;PORT=50001;UID=db2inst1;PWD=db2inst1;AUTHENTICATION=SERVER;CurrentSchema=ENTTEST"
	client, err := hzwent.Open(drivername, con)
	if err != nil {
		log.Fatalf("failed opening connection to %s: %v", drivername, err)
	}
	defer client.Close()

	// insert
	istu, err := client.Student.
		Create().
		SetName("张三").
		SetStuNo("001").
		SetAge(18).
		Save(ctx)
	if err != nil {
		log.Fatalf("failed creating user: %v", err)
	}
	log.Printf("student was created: id:%v  %v ", istu.ID, istu)

	// query
	stuid := istu.ID
	qstu, err := client.Student.
		Query().
		Where(student.ID(stuid)).
		// `Only` fails if no user found,
		// or more than 1 user returned.
		Only(ctx)
	if err != nil {
		log.Fatalf("failed querying user: %v", err)
	}
	log.Printf("student returned: %v", qstu)

	// update with query
	usty := qstu.Update().SetName("小天狼").SaveX(ctx)
	log.Printf("student updated with query: %v", usty)

	// update one
	ustu, err := client.Student.
		UpdateOne(qstu).
		SetAge(20).
		Save(ctx)
	if err != nil {
		log.Fatalf("failed updating user: %v", err)
	}
	log.Printf("student updated one: %v", ustu)

	// update by id
	ustu = client.Student.UpdateOneID(stuid).SetName("hzw").SaveX(ctx)
	log.Printf("student updated by id: %v", ustu)

	// delete
	//err = client.Student.DeleteOneID(ustu.ID).Exec(ctx)
	//if err != nil {
	//	log.Fatalf("failed delete user: %v", err)
	//}
	//log.Printf("student was deleted: %v", ustu)

	// insert batch

}

func TestDb2(t *testing.T) {
	ctx := context.Background()
	drivername := "go_ibm_db"
	con := "HOSTNAME=localhost;DATABASE=testdb;PORT=50001;UID=db2inst1;PWD=db2inst1;AUTHENTICATION=SERVER;CurrentSchema=ENTTEST"

	db, err := sql.Open(drivername, con)
	if err != nil {
		log.Fatalf("failed opening connection to %s: %v", drivername, err)
	}
	defer db.Close()

	insertsql := "INSERT INTO student (STU_NO, NAME, AGE) VALUES ('222', '222', 22)"
	_, err = db.ExecContext(ctx, insertsql)
	if err != nil {
		log.Fatalf("failed insert: %v", err)
	}

	id := 0
	getIdQuery := "SELECT IDENTITY_VAL_LOCAL() FROM SYSIBM.SYSDUMMY1"
	err = db.QueryRowContext(ctx, getIdQuery).Scan(&id)
	if err != nil {
		log.Fatalf("failed get id: %v", err)
	}

	fmt.Printf("id:%d\n", id)

}

func TestDb2_2(t *testing.T) {
	ctx := context.Background()
	drivername := "go_ibm_db"
	con := "HOSTNAME=localhost;DATABASE=testdb;PORT=50001;UID=db2inst1;PWD=db2inst1;AUTHENTICATION=SERVER;CurrentSchema=ENTTEST"

	db, err := sql.Open(drivername, con)
	db2, err := sql.Open(drivername, con)
	defer db.Close()
	defer db2.Close()

	insertsql := "INSERT INTO student (STU_NO, NAME, AGE) VALUES ('222', '222', 22)"
	_, err = db.ExecContext(ctx, insertsql)
	_, err = db2.ExecContext(ctx, insertsql)
	_, err = db.ExecContext(ctx, insertsql)

	time.Sleep(1 * time.Second)

	getIdQuery := "SELECT IDENTITY_VAL_LOCAL() FROM SYSIBM.SYSDUMMY1"
	id2 := 0
	err = db2.QueryRowContext(ctx, getIdQuery).Scan(&id2)

	time.Sleep(1 * time.Second)

	id := 0
	err = db.QueryRowContext(ctx, getIdQuery).Scan(&id)

	if err != nil {
		log.Fatalf("failed get id: %v", err)
	}

	fmt.Printf("id:%d\n", id)
	fmt.Printf("id2:%d\n", id2)

}

func BenchmarkEnt(b *testing.B) {
	ctx := context.Background()
	drivername := "go_ibm_db"
	con := "HOSTNAME=localhost;DATABASE=testdb;PORT=50001;UID=db2inst1;PWD=db2inst1;AUTHENTICATION=SERVER;CurrentSchema=ENTTEST"
	client, err := hzwent.Open(drivername, con)
	if err != nil {
		log.Fatalf("failed opening connection to %s: %v", drivername, err)
	}
	defer client.Close()

	// insert
	istu, err := client.Student.
		Create().
		SetName("张三").
		SetStuNo("001").
		SetAge(18).
		Save(ctx)
	if err != nil {
		log.Fatalf("failed creating user: %v", err)
	}
	log.Printf("student was created: id:%v  %v ", istu.ID, istu)

	err = client.Student.DeleteOneID(istu.ID).Exec(ctx)
	if err != nil {
		log.Fatalf("failed delete user: %v", err)
	}
	log.Printf("student was deleted: %v", istu)

}
