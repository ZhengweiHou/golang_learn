package dao

import (
	"log"

	ibmdb "github.com/ZhengweiHou/gorm_ibmdb"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	condb2   = "HOSTNAME=localhost;DATABASE=testdb;PORT=50001;UID=db2inst1;PWD=db2inst1;AUTHENTICATION=SERVER;CurrentSchema=ENTTEST"
	conmysql = "root:root@tcp(localhost:3306)/test?parseTime=True"
)

var globaldb *gorm.DB

func InitDb() {
	dialector := mysqlDialector()
	// dialector := db2Dialector()

	conf := &gorm.Config{
		PrepareStmt: true,
		Logger:      logger.New(log.New(log.Writer(), "\r\n", log.LstdFlags), logger.Config{LogLevel: logger.Info}),
	}

	var err error
	globaldb, err = gorm.Open(dialector, conf)
	if err != nil {
		log.Fatalf("open err: %v", err)
	}
}

func mysqlDialector() gorm.Dialector {
	return mysql.Open(conmysql)
}

func db2Dialector() gorm.Dialector {
	return ibmdb.Open(condb2)
}
