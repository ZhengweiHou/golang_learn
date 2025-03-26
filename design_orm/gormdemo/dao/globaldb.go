package dao

import (
	"log"

	"gorm.io/driver/ibmdb"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var conhzw = "HOSTNAME=localhost;DATABASE=testdb;PORT=50001;UID=db2inst1;PWD=db2inst1;AUTHENTICATION=SERVER;CurrentSchema=ENTTEST"

var globaldb *gorm.DB

func init() {
	dialector := ibmdb.Open(conhzw)

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
