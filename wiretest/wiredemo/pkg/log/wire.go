//go:build wireinject
// +build wireinject

/*
 */

package log

import (
	"github.com/google/wire"
)

var LogWireSet = wire.NewSet(
	NewZapLog,
	//	NewLog,
	NewZapHandler,
	NewZapSlog,
)
