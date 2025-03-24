/*
//go:build wireinject
// +build wireinject
*/

package db

import "github.com/google/wire"

var DbWireSet = wire.NewSet(
	NewDB,                 // gorm DB
	NewRepository,         // Repository
	NewTransactionManager, // 事务管理器
)
