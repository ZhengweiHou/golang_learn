package service

import (
	"log/slog"

	db "aic.com/pkg/aicgormdb"
)

type BaseService struct {
	// logger *log.Logger
	logger *slog.Logger
	tm     db.TransactionManager
}

func NewService(
	tm db.TransactionManager,
	logger *slog.Logger,
) *BaseService {
	return &BaseService{
		logger: logger,
		tm:     tm,
	}
}
