package db2

import (
	"database/sql"
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	"github.com/axgle/mahonia"
	_ "github.com/ibmdb/go_ibm_db"
	"golang.org/x/text/encoding/simplifiedchinese"
)

func TestSelect1(t *testing.T) {

	startT := time.Now()

	dataSourceName := "HOSTNAME=localhost;DATABASE=testdb;PORT=50001;UID=db2inst1;PWD=db2inst1;AUTHENTICATION=SERVER;CurrentSchema=DBSYNCTEST"
	// type Db *sql.DB
	var db *sql.DB
	var stmt *sql.Stmt
	var rows *sql.Rows

	// 创建数据库操作对象
	db, _ = sql.Open("go_ibm_db", dataSourceName)

	tx, _ := db.Begin() // 开启事务
	tx.Commit()         // 提交事务
	// tx.Rollback()       // 回滚事务

	//
	stmt, _ = db.Prepare("select * from student")

	// 执行查询 TODO 此处是否已将所有数据查询出来
	rows, _ = stmt.Query()

	// 通过结果集获取字段名
	fmt.Println("================")
	cols, _ := rows.Columns()
	for i := 0; i < len(cols); i++ {
		fmt.Print("  ", cols[i])
	}
	fmt.Println("\n================")

	defer rows.Close()
	// scan获取记录
	for rows.Next() {

		count := len(cols)
		var record = make(map[string]interface{})

		var rowValues = make([]interface{}, count)
		var valuePointers = make([]interface{}, count)
		for i := range rowValues {
			valuePointers[i] = &rowValues[i] // 将原切片中元素的指针取出
		}
		err := rows.Scan(valuePointers...) // 参数必须是指针pointer
		// err := rows.Scan(&t, &x, &m, &n)
		if err != nil {
			log.Panic(err)
		}
		enc := mahonia.NewDecoder("gbk")
		// 格式化数据
		for i := range rowValues {
			var value interface{}
			rawValue := rowValues[i]
			b, ok := rawValue.([]byte) //byte，占用1个节字，就 8 个比特位（2^8 = 256，因此 byte 的表示范围 0->255），所以它和 uint8 类型本质上没有区别，它表示的是 ACSII 表中的一个字符
			if ok {
				// value = string(b) //string 的本质，其实是一个 byte数组
				value = enc.ConvertString(string(b))
			} else {
				value = rawValue
			}
			rowValues[i] = value
		}

		// 将数据放入map中
		for i, col := range cols {
			record[col] = rowValues[i]
		}

		fmt.Println("===", record)
		tc := time.Since(startT)
		fmt.Printf("time cost = %v\n", tc)
	}

}

// test db2 select

func TestSelect2(t *testing.T) {
	dataSourceName := "HOSTNAME=localhost;DATABASE=testdb;PORT=50003;UID=db2inst1;PWD=db2inst1;AUTHENTICATION=SERVER;CurrentSchema=TEST;DB2CODEPAGE=1208"
	db, _ := sql.Open("go_ibm_db", dataSourceName)
	rows, _ := db.Query("select * from TT2")

	var id int
	var name string
	for rows.Next() {
		rows.Scan(&id, &name)
		fmt.Printf("id:%v, name:%v\n", id, name)
	}
}

func TestSelect3(t *testing.T) {
	fmt.Printf("DB2CODEPAGE:%s\n", os.Getenv("DB2CODEPAGE"))
	dataSourceName := "HOSTNAME=localhost;DATABASE=testdb;PORT=50003;UID=db2inst1;PWD=db2inst1;CurrentSchema=TEST;"
	db, _ := sql.Open("go_ibm_db", dataSourceName)
	rows, _ := db.Query("select id,name from student3 where id in (1,2,3)")
	var id int
	var name string
	for rows.Next() {
		rows.Scan(&id, &name)
		hexnameStr := hex.EncodeToString([]byte(name))
		fmt.Printf("hexname:%s,id:%d,name:%s\n", hexnameStr, id, name)
	}
}
func TestSelect4(t *testing.T) {
	fmt.Printf("DB2CODEPAGE:%s\n", os.Getenv("DB2CODEPAGE"))
	dataSourceName := "HOSTNAME=localhost;DATABASE=testdb;PORT=50003;UID=db2inst1;PWD=db2inst1;CurrentSchema=TEST;"
	db, _ := sql.Open("go_ibm_db", dataSourceName)
	rows, _ := db.Query("select id,name from student3 where id in (1,2,3)")
	var id int
	var name string
	for rows.Next() {
		rows.Scan(&id, &name)
		hexnameStr := hex.EncodeToString([]byte(name))
		fmt.Printf("hexname:%s,id:%d,name:%s\n", hexnameStr, id, name)

		// tname, _ := io.ReadAll(transform.NewReader(bytes.NewBuffer([]byte(name)), unicode.UTF8.NewEncoder()))
		// tname, _ := io.ReadAll(transform.NewReader(bytes.NewBuffer([]byte(name)), simplifiedchinese.GBK.NewDecoder()))
		tname, _ := simplifiedchinese.GBK.NewDecoder().Bytes([]byte(name))
		thexnameStr := hex.EncodeToString([]byte(tname))
		fmt.Printf("hexname:%s,id:%d,name:%s\n", thexnameStr, id, tname)
	}
}
