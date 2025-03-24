package repository

//import (
//	"log"
//
//	"gorm.io/driver/ibmdb"
//	"gorm.io/driver/mysql"
//	"gorm.io/gorm"
//	"gorm.io/gorm/logger"
//)
//
//var conhzwdb2 = "HOSTNAME=localhost;DATABASE=testdb;PORT=50001;UID=db2inst1;PWD=db2inst1;AUTHENTICATION=SERVER;CurrentSchema=ENTTEST"
//var conhzwmysql = "root:root@tcp(localhost:3306)/test?parseTime=True"
//var hzwlog = log.New(log.Writer(), "\r\n", log.LstdFlags)
//var logconf = logger.Config{LogLevel: logger.Info}
//
//func NewDb() *gorm.DB {
//	return getHzwDB2Db()
//}
//
//func getHzwDB2Db() *gorm.DB {
//	dialector := ibmdb.Open(conhzwdb2)
//
//	conf := &gorm.Config{
//		PrepareStmt: true,
//		//		DryRun:      true, // 生成SQL而不执行
//		Logger: logger.New(hzwlog, logconf),
//	}
//
//	db, err := gorm.Open(dialector, conf)
//	if err != nil {
//		log.Fatalf("open err: %v", err)
//	}
//
//	return db
//}
//func getHzwMysqlDb() *gorm.DB {
//	dialector := mysql.Open(conhzwmysql)
//	conf := &gorm.Config{
//		PrepareStmt: true,
//		//		DryRun:      true, // 生成SQL而不执行
//		Logger: logger.New(hzwlog, logconf),
//	}
//	db, err := gorm.Open(dialector, conf)
//	if err != nil {
//		log.Fatalf("open err: %v", err)
//	}
//	return db
//}
//
