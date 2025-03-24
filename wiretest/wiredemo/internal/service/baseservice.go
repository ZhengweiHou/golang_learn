package service

import (
	"wiredemo/pkg/db"
	"wiredemo/pkg/log"
)

type BaseService struct {
	logger *log.Logger
	tm     db.TransactionManager
}

func NewService(
	tm db.TransactionManager,
	logger *log.Logger,
) *BaseService {
	return &BaseService{
		logger: logger,
		tm:     tm,
	}
}
