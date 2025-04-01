/*
//go:build wireinject
// +build wireinject

*/

package repository

import (
	"wiredemo/internal/repository/dao"

	"github.com/google/wire"
)

var RepositoryWireSet = wire.NewSet(
	dao.NewHzw2Dao,
	dao.NewHzwDao,
)
