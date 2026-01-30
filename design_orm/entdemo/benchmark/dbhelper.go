package benchmark

import (
	"log"

	"entdemo/benchmark/hzwent"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/ibmdb/go_ibm_db"
)

// GetCrudDb 获取数据库客户端
// dbType: "mysql" 或 "db2"
func GetCrudDb(dsn string, dbType string) *hzwent.Client {
	var driverName string
	
	switch dbType {
	case "mysql":
		driverName = "mysql"
	case "db2":
		driverName = "go_ibm_db"
	default:
		log.Fatalf("unsupported database type: %s, supported types: mysql, db2", dbType)
	}

	// 使用 Ent 的 Open 函数创建客户端
	client, err := hzwent.Open(driverName, dsn)
	if err != nil {
		log.Fatalf("failed to create ent client with driver %s: %v", driverName, err)
	}

	return client
}
