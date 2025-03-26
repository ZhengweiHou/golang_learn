package aicgormdb

import (
	"time"
	//	"wiredemo/pkg/log"

	"github.com/spf13/viper"
	"gorm.io/driver/ibmdb"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// func NewDB(conf *viper.Viper, l *zap.Logger) *gorm.DB {
func NewDB(conf *viper.Viper, l logger.Interface) *gorm.DB {
	var (
		db  *gorm.DB
		err error
	)

	//logger := NewGormLog(l)
	logger := l
	driver := conf.GetString("data.db.user.driver")
	dsn := conf.GetString("data.db.user.dsn")

	// GORM doc: https://gorm.io/docs/connecting_to_the_database.html
	switch driver {
	case "ibmdb":
		db, err = gorm.Open(ibmdb.Open(dsn), &gorm.Config{ // 数据库不可用会报异常
			Logger: logger,
		})
	case "mysql":
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
			Logger: logger,
		})
	default:
		panic("unknown db driver")
	}
	if err != nil {
		panic(err)
	}
	db = db.Debug()

	// Connection Pool config
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)
	return db
}
