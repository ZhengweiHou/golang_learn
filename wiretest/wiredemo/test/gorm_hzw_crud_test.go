package test

import (
	"encoding/json"
	"fmt"
	"log"
	"testing"
	"wiredemo/internal/repository/model"

	"gorm.io/driver/ibmdb"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var conhzwdb2 = "HOSTNAME=localhost;DATABASE=testdb;PORT=50001;UID=db2inst1;PWD=db2inst1;AUTHENTICATION=SERVER;CurrentSchema=ENTTEST"
var conhzwmysql = "root:root@tcp(localhost:3306)/test?parseTime=True"
var hzwlog = logger.New(log.New(log.Writer(), "\r\n", log.LstdFlags), logger.Config{LogLevel: logger.Info})

func getHzwDb() *gorm.DB {
	//	return getHzwMysqlDb()
	return getHzwDB2Db()
}
func getHzwDB2Db() *gorm.DB {
	dialector := ibmdb.Open(conhzwdb2)

	conf := &gorm.Config{
		PrepareStmt: true,
		//		DryRun:      true, // 生成SQL而不执行
		Logger: hzwlog,
	}

	db, err := gorm.Open(dialector, conf)
	if err != nil {
		log.Fatalf("open err: %v", err)
	}

	return db
}
func getHzwMysqlDb() *gorm.DB {
	dialector := mysql.Open(conhzwmysql)
	conf := &gorm.Config{
		PrepareStmt: true,
		//		DryRun:      true, // 生成SQL而不执行
		Logger: hzwlog,
	}
	db, err := gorm.Open(dialector, conf)
	if err != nil {
		log.Fatalf("open err: %v", err)
	}
	return db
}

// 新增
func TestHzwCreate(t *testing.T) {
	db := getHzwDb()

	// ========= Create =========
	hzw := model.Hzw{
		//hzw := Hzw{
		Name:     "hzw",
		Age:      20,
		Decimal1: 999999.99999999, // 数据库
	}
	// ---- DryRun show Create sql ----
	//stmt := db.Session(&gorm.Session{DryRun: true}).Create(&hzw).Statement
	//log.Printf("Create SQL: %v\n", stmt.SQL.String())
	//log.Printf("Create SQL Vars: %v\n", stmt.Vars)

	fmt.Println(hzw.CreatedAt.Unix())
	// ---- Create ----
	// TODO CreateAt什么时候赋值的
	result := db.Create(&hzw)
	if result.Error != nil {
		log.Fatalf("Create failed: %v", result.Error)
	}
	jstr, err := json.Marshal(hzw)
	if err != nil {
		log.Fatalf("json marshal failed: %v", err)
	}
	fmt.Printf("Affected: %v Created hzw: %s\n", result.RowsAffected, jstr)

}
