package main

import (
	"fmt"

	// _ "github.com/go-sql-driver/mysql"
	_ "github.com/ibmdb/go_ibm_db"

	"github.com/viant/dsc"
	// "hzw/dsc"
)

// func Test_db21() {
// 	DB, _ := sql.Open("go_ibm_db", "HOSTNAME=localhost;DATABASE=testdb;PORT=50000;UID=db2inst1;PWD=db2inst1;AUTHENTICATION=SERVER;SCHEMA=DBSYNCTEST")
// 	err := DB.Ping()
// 	if err != nil {
// 		panic(err)
// 	}
// }

func main() {
	factory := dsc.NewManagerFactory()
	config := dsc.NewConfig("go_ibm_db",
		"HOSTNAME=[host];DATABASE=[database];PORT=[port];UID=[user];PWD=[pwd];AUTHENTICATION=SERVER;CurrentSchema=[schema]",
		"user:db2inst1,pwd:db2inst1,host:localhost,database:testdb,port:50000,schema:DBSYNCTEST")
	manager, err := factory.Create(config)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(manager.Config().DriverName)

	dialect := dsc.GetDatastoreDialect(manager.Config().DriverName)
	datastore, err := dialect.GetCurrentDatastore(manager)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(datastore)

	dbs, err := dialect.GetDatastores(manager)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(dbs)

	// var columns []dsc.Column
	columns, err := dialect.GetColumns(manager, datastore, "STUDENT3")
	if err != nil {
		panic(err.Error())
	}
	for _, column := range columns {
		l, hasl := column.Length()
		p, s, hasps := column.DecimalSize()
		fmt.Printf("name:%-10s type:%-10s hasl:%-5t l:%-3d hasps:%-5t p:%-2d s:%-2d\n",
			column.Name(), column.DatabaseTypeName(),
			hasl, l,
			hasps, p, s)
	}

	keyNames := dialect.GetKeyName(manager, "DBSYNCTEST", "STUDENT3")
	fmt.Println(keyNames)

	ddl, err := dialect.ShowCreateTable(manager, "STUDENT3") // TODO 待处理
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(ddl)

	// manager.ReadAllWithHandler("select * from student", nil, func(scanner dsc.Scanner) (toContinue bool, err error) {
	// 	var record = make(map[string]interface{})

	// 	err = scanner.Scan(&record) //从数据库读取一条数据

	// 	log.Println("===", record)
	// 	toContinue = true
	// 	return
	// })

	// var record = make(map[string]interface{})
	// manager.ReadAll(record, "select * from student", nil,)

}
