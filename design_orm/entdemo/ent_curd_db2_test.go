package main

import (
	"context"
	"entdemo/hzwent"
	"entdemo/hzwent/student"
	"log"
	"testing"
)

func TestCreate(t *testing.T) {
	ctx := context.Background()
	drivername := "go_ibm_db"
	con := "HOSTNAME=localhost;DATABASE=testdb;PORT=50001;UID=db2inst1;PWD=db2inst1;AUTHENTICATION=SERVER;CurrentSchema=ENTTEST"
	client, err := hzwent.Open(drivername, con)
	if err != nil {
		log.Fatalf("failed opening connection to %s: %v", drivername, err)
	}
	defer client.Close()

	// create one
	istu, err := client.Student.
		Create().
		SetName("张三").
		SetStuNo("001").
		SetAge(18).
		Save(ctx) // SaveX(ctx) 会panic
	if err != nil {
		log.Fatalf("failed creating user: %v", err)
	}
	log.Printf("student create one: id:%v  %v ", istu.ID, istu)

	// TODO db2 not support LastInsertId，可以正常插入，但无法获取自增id
	//// create many
	//istus := client.Student.CreateBulk(
	//	client.Student.Create().SetName("李四").SetStuNo("002").SetAge(19),
	//	client.Student.Create().SetName("王五").SetStuNo("003").SetAge(20),
	//).SaveX(ctx)
	//log.Printf("student create many: %v ", istus)

	//// create many Mapcreate
	//names := []string{"赵六", "孙七", "周八"}
	//istus = client.Student.MapCreateBulk(names, func(sc *hzwent.StudentCreate, i int) {
	//	sc.SetName(names[i]).SetStuNo(names[i]).SetAge(10 + i)
	//}).SaveX(ctx)
	//log.Printf("student create many Mapcreate: %v ", istus)
}

func TestUpdate(t *testing.T) {
	ctx := context.Background()
	drivername := "go_ibm_db"
	con := "HOSTNAME=localhost;DATABASE=testdb;PORT=50001;UID=db2inst1;PWD=db2inst1;AUTHENTICATION=SERVER;CurrentSchema=ENTTEST"
	client, err := hzwent.Open(drivername, con)
	if err != nil {
		log.Fatalf("failed opening connection to %s: %v", drivername, err)
	}
	defer client.Close()

	// create one
	// SQL: INSERT INTO student (STU_NO, NAME, AGE, VERSION) VALUES (?, ?, ?, ?)
	stu := client.Student.Create().SetName("王二小").SetStuNo("999").SetAge(8).SaveX(ctx)
	log.Printf("1: student create one: id:%v  %v ", stu.ID, stu)

	// update one
	// SQL: UPDATE student SET NAME = ? WHERE ID = ?
	ustu := stu.Update().
		SetName("王二中").
		SaveX(ctx)
	log.Printf("2: student update one: %v", ustu)

	// update by ID
	// SQL: UPDATE student SET NAME = ? WHERE ID = ?
	stuid := stu.ID
	ustu = client.Student.
		UpdateOneID(stuid).
		SetName("王二大").
		SaveX(ctx)
	log.Printf("3: student update by ID: %v", ustu)

	// updata by id (version)
	// SQL: UPDATE student SET NAME = ?, VERSION = COALESCE(student.VERSION, 0) + ? WHERE ID = ? AND student.VERSION = ?
	ustu, err = client.Student.
		UpdateOneID(ustu.ID).
		SetName("王二大大").
		AddVersion(1). // 乐观锁+1
		Where(
			student.Version(ustu.Version), // 乐观锁 panic: version mismatch
		).
		Save(ctx)
	switch {
	case hzwent.IsNotFound(err):
		log.Fatalf("student not found: %v", err)
	case err != nil:
		log.Fatalf("failed updating user: %v", err)
	}
	log.Printf("4: student update by ID: %v", ustu)

	// update by entity (version)
	// SQL: UPDATE student SET NAME = ?, VERSION = COALESCE(student.VERSION, 0) + ? WHERE ID = ? AND student.VERSION = ?
	ustu, err = client.Student.
		UpdateOne(ustu).
		SetName("王二大大大").
		AddVersion(1). // 乐观锁+1
		Where(
			student.Version(ustu.Version), // 乐观锁 panic: version mismatch
			//student.Version(3), // 乐观锁 err: version mismatch
		).
		Save(ctx)
	switch {
	case hzwent.IsNotFound(err):
		log.Fatalf("student not found: %v", err)
	case err != nil:
		log.Fatalf("failed updating user: %v", err)
	}
	log.Printf("5: student update by entity: %v", ustu)

	// update many
	// SQL: UPDATE student SET AGE = ? WHERE student.NAME = ? OR student.NAME = ?
	affectsize, err := client.Student.Update().
		Where(
			student.Or(
				student.Name("王二大大大"),
				student.Name("王二大大"),
			),
		).SetAge(99).
		Save(ctx)
	if err != nil {
		log.Fatalf("failed updating user: %v", err)
	}
	log.Printf("6: student update many: %v", affectsize)

	// update many like
	// SQL: UPDATE student SET AGE = ? WHERE student.NAME LIKE ?
	affectsize, err = client.Student.Update().
		Where(
			student.NameContains("王二"),
		).
		SetAge(100).
		Save(ctx)
	if err != nil {
		log.Fatalf("failed updating user: %v", err)
	}
	log.Printf("7: student update many like: %v", affectsize)

	// update many prefix
	// SQL: UPDATE student SET AGE = ? WHERE student.NAME LIKE ?
	affectsize, err = client.Student.Update().
		Where(
			student.NameHasPrefix("大大"), // 参数被拼为大大%
		).
		SetAge(101).
		Save(ctx)
	if err != nil {
		log.Fatalf("failed updating user: %v", err)
	}
	log.Printf("8: student update many prefix: %v", affectsize)
}

func TestUpsert(t *testing.T) {
	//	ctx := context.Background()
	drivername := "go_ibm_db"
	con := "HOSTNAME=localhost;DATABASE=testdb;PORT=50001;UID=db2inst1;PWD=db2inst1;AUTHENTICATION=SERVER;CurrentSchema=ENTTEST"
	client, err := hzwent.Open(drivername, con)
	if err != nil {
		log.Fatalf("failed opening connection to %s: %v", drivername, err)
	}
	defer client.Close()

	// upsert one
	// SQL: INSERT INTO student (STU_NO, NAME, AGE, VERSION, ID) VALUES (?, ?, ?, ?, ?)
	//client.Student.Create().
	//	SetID(1).SetName("张三三").SetStuNo("001").SetAge(18).
	//	OnConflict().
	//	UpdateNewValues().
	//	Exec(ctx)

}

func TestQuery(t *testing.T) {
	ctx := context.Background()
	drivername := "go_ibm_db"
	con := "HOSTNAME=localhost;DATABASE=testdb;PORT=50001;UID=db2inst1;PWD=db2inst1;AUTHENTICATION=SERVER;CurrentSchema=ENTTEST"
	client, err := hzwent.Open(drivername, con)
	if err != nil {
		log.Fatalf("failed opening connection to %s: %v", drivername, err)
	}
	defer client.Close()

	// query LT
	// SQL: SELECT student.ID, student.STU_NO, student.NAME, student.AGE, student.VERSION FROM student WHERE student.ID < ?
	qstus, err := client.Student.Query().
		Where(
			student.IDLT(10),
		).
		All(ctx)
	if err != nil {
		log.Fatalf("failed querying users: %v", err)
	}
	log.Printf("1: student query LT: ")
	for _, stu := range qstus {
		log.Printf("    %v", stu)
	}

	// query field selection
	var v []struct {
		Name string
		Age  int
	}
	err = client.Student.Query().
		Select(student.FieldName, student.FieldAge).Scan(ctx, &v)
	if err != nil {
		log.Fatalf("failed querying users: %v", err)
	}
	log.Printf("2: student query field selection: %v", v)

}
