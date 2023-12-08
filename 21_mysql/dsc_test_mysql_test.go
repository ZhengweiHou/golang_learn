package main

// import (
// 	"fmt"
// 	"log"
// 	"testing"

// 	_ "github.com/go-sql-driver/mysql"

// 	"github.com/viant/dsc"
// 	// "hzw/dsc"
// )

// // func Test_mysql1(t *testing.T) {
// // 	DB, _ := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/dbsynctest?parseTime=true")
// // 	err := DB.Ping()
// // 	if err != nil {
// // 		panic(err)
// // 	}
// // }

// // func main() {

// func TestXxx(t *testing.T) {
// 	factory := dsc.NewManagerFactory()
// 	config := dsc.NewConfig("mysql", "[user]:[pwd]@[url]", "user:root,pwd:root,url:tcp(localhost:3306)/dbsynctest?parseTime=true")
// 	manager, err := factory.Create(config)
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	fmt.Println(manager.Config().DriverName)

// 	dialect := dsc.GetDatastoreDialect(manager.Config().DriverName)
// 	datastore, err := dialect.GetCurrentDatastore(manager)
// 	if err != nil {
// 		panic(err.Error())
// 	}

// 	// var columns []dsc.Column
// 	columns, err := dialect.GetColumns(manager, datastore, "student")
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	fmt.Println(columns)

// 	manager.ReadAllWithHandler("select * from student", nil, func(scanner dsc.Scanner) (toContinue bool, err error) {
// 		var record = make(map[string]interface{})

// 		err = scanner.Scan(&record) //从数据库读取一条数据

// 		log.Println("===", record)
// 		toContinue = true
// 		return
// 	})

// 	// var record = make(map[string]interface{})
// 	// manager.ReadAll(record, "select * from student", nil,)

// }
