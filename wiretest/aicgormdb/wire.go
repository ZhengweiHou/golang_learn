/*
//go:build wireinject
// +build wireinject

*/

package aicgormdb

import (
	"github.com/google/wire"
)

var DbWireSet = wire.NewSet(
	NewDB,                 // gorm DB
	NewRepository,         // Repository
	NewTransactionManager, // 事务管理器
	//NewSLogGormLog,
	NewZapGormLog,
)
