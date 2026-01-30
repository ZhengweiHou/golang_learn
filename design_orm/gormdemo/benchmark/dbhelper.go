package benchmark

import (
	"log"

	ibmdb "github.com/ZhengweiHou/gorm_ibmdb"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// GetCrudDb 获取数据库客户端
// dbType: "mysql" 或 "db2"
func GetCrudDb(dsn string, dbType string) *gorm.DB {
	var dialector gorm.Dialector
	
	switch dbType {
	case "mysql":
		dialector = mysql.Open(dsn)
	case "db2":
		dialector = ibmdb.Open(dsn)
	default:
		log.Fatalf("unsupported database type: %s, supported types: mysql, db2", dbType)
	}

	conf := &gorm.Config{
		PrepareStmt: false,
		//		DryRun:      true, // 生成SQL而不执行
		// Logger: logger.New(log.New(log.Writer(), "\r\n", log.LstdFlags), logger.Config{LogLevel: logger.Info}),
	}

	db, err := gorm.Open(dialector, conf)
	if err != nil {
		log.Fatalf("open err: %v", err)
	}

	return db
}
