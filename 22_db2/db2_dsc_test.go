package db2

import (
	"fmt"
	"log"
	"testing"
	"time"

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

func TestDb2Dsc1(t *testing.T) {
	factory := dsc.NewManagerFactory()
	config := dsc.NewConfig("go_ibm_db",
		"HOSTNAME=[host];DATABASE=[database];PORT=[port];UID=[user];PWD=[pwd];AUTHENTICATION=SERVER;CurrentSchema=[schema]",
		"user:db2inst1,pwd:db2inst1,host:localhost,database:testdb,port:50001,schema:DBSYNCTEST")
	manager, err := factory.Create(config)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(manager.Config().DriverName)

	// insert 包含 decimal 类型字段
	destCon, err := manager.ConnectionProvider().Get()
	var record = make([]interface{}, 0)
	record = append(record, "刘德华222", 1, 1, "11.11", time.Now()) // FEE 是DECIMAL 类型，传参使用string方式则会丢失精度
	record = append(record, "刘德华111", 1, 1, 11.11, time.Now())   // 使用float传参就不会有问题
	result, err := manager.ExecuteOnConnection(destCon, "INSERT INTO STUDENT2(NAME,AGE,GRADES,FEE,MODIFIED) VALUES(?,?,?,?,?),(?,?,?,?,?)", record)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("insert result:%v\n", result)
}

func TestDb2Dsc2(t *testing.T) {
	factory := dsc.NewManagerFactory()
	config := dsc.NewConfig("go_ibm_db",
		"HOSTNAME=[host];DATABASE=[database];PORT=[port];UID=[user];PWD=[pwd];AUTHENTICATION=SERVER;CurrentSchema=[schema]",
		"user:db2inst1,pwd:db2inst1,host:localhost,database:testdb,port:50001,schema:DBSYNCTEST")
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
	columns, err := dialect.GetColumns(manager, datastore, "STUDENT")
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

	/*===dmlProvider test ===*/

	var table *dsc.TableDescriptor
	table = manager.TableDescriptorRegistry().Get("STUDENT3")
	dmlProvider := dsc.NewMapDmlProvider(table)
	fmt.Println(dmlProvider)
	// dmlProvider.Get(dsc.SQLTypeInsert, instance interface{})

	manager.ReadAllWithHandler("select * from student", nil, func(scanner dsc.Scanner) (toContinue bool, err error) {
		var record = make(map[string]interface{})

		err = scanner.Scan(&record) //从数据库读取一条数据

		log.Println("=1=", record)
		toContinue = true
		return
	})

	manager.ReadAllWithHandler("select * from student", nil, func(scanner dsc.Scanner) (toContinue bool, err error) {

		var rowValues = make([]interface{}, 6)
		var valuePointers = make([]interface{}, 6)
		for i := range rowValues {
			valuePointers[i] = &rowValues[i] // 将原切片中元素的指针取出
		}

		err = scanner.Scan(valuePointers...) //从数据库读取一条数据

		for i := range rowValues {
			var value interface{}
			rawValue := rowValues[i]
			b, ok := rawValue.([]byte) //byte，占用1个节字，就 8 个比特位（2^8 = 256，因此 byte 的表示范围 0->255），所以它和 uint8 类型本质上没有区别，它表示的是 ACSII 表中的一个字符
			if ok {
				value = string(b) //string 的本质，其实是一个 byte数组
			} else {
				value = rawValue
			}
			rowValues[i] = value
		}

		log.Println("=2=", rowValues)
		toContinue = true
		return
	})

	// var record = make(map[string]interface{})
	// manager.ReadAll(record, "select * from student", nil, nil)

}
