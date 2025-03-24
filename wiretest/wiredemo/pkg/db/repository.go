package db

import (
	"context"
	"database/sql"
	"wiredemo/pkg/log"

	"gorm.io/gorm"
)

const CtxTxKey = "AicTxKey"

// TM TransactionManager
// var TM TransactionManager // TODO TransactionManager

type Repository struct {
	db *gorm.DB
	//tm     TransactionManager // 循环注入
	logger *log.Logger
}

type TransactionManager interface {
	Transaction(ctx context.Context, fn func(ctx context.Context) error) error
	WithTransaction(ctx context.Context, opts ...*sql.TxOptions) (context.Context, func(error))
}

func NewRepository(
	logger *log.Logger,
	db *gorm.DB,
	//tm TransactionManager, // 循环注入
) *Repository {
	return &Repository{
		db:     db,
		logger: logger,
		//tm:     tm,
	}
}

func NewTransactionManager(repository *Repository) TransactionManager {
	//TM = repository // 保留一个全局对象，方便事务操作
	return repository
}

func (r *Repository) DB(ctx context.Context) *gorm.DB {
	// 若上下文开启了事务则返回上下文事务
	v := ctx.Value(CtxTxKey)
	if v != nil {
		if tx, ok := v.(*gorm.DB); ok {
			r.logger.Info("已开启事务")
			return tx
		}
	}
	r.logger.Info("新事务")
	// 若未开启事务则返回新db，事务行为为默认方式
	return r.db.WithContext(ctx)
}

// 开启事务处理
func (r *Repository) Transaction(ctx context.Context, fn func(ctx context.Context) error) error {
	// TODO 存在当前事务时，处理事务传播逻辑
	// 将传入的业务处理fn包装到gorm的Transaction中处理
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		ctx = context.WithValue(ctx, CtxTxKey, tx) // 创建子ctx，置入事务对象，后续数据库操作要使用该事务对象
		r.logger.Info("开启事务")
		return fn(ctx)
	})
}

func (r *Repository) WithTransaction(ctx context.Context, opts ...*sql.TxOptions) (context.Context, func(error)) {
	tx := r.db.Begin(opts...)
	txctx := context.WithValue(ctx, CtxTxKey, tx)
	r.logger.Info("开启事务")
	return txctx, func(err error) {

		if err != nil {
			r.logger.Info("事务回滚")
			tx.Rollback()
		} else {
			select {
			case <-ctx.Done():
				r.logger.Info("事务回滚")
				tx.Rollback()
				return
			default:
				r.logger.Info("事务提交")
				tx.Commit()
			}
		}
	}
}
