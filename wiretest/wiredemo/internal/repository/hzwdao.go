package repository

import (
	"context"
	"wiredemo/internal/model"
	"wiredemo/pkg/db"
)

// IHzwDao 接口
type IHzwDao interface {
	InsertOne(ctx context.Context, hzw *model.Hzw) (*model.Hzw, error)
	FindBykey(ctx context.Context, id int) (*model.Hzw, error)
}

type HzwDao struct {
	*db.Repository
}

// func NewHzwDao(db *gorm.DB) *HzwDao {
//
//	func NewHzwDao(db *gorm.DB) IHzwDao {
//		return &HzwDao{
//			db: db,
//		}
//	}
func NewHzwDao(r *db.Repository) IHzwDao {
	return &HzwDao{
		Repository: r,
	}
}

func (dao *HzwDao) InsertOne(ctx context.Context, hzw *model.Hzw) (*model.Hzw, error) {
	//result := dao.db.Create(hzw)
	result := dao.DB(ctx).Create(hzw) // ctx中可能已存在当前事务实例
	if result.Error != nil {
		return nil, result.Error
	}
	return hzw, nil
}

func (dao *HzwDao) FindBykey(ctx context.Context, key int) (*model.Hzw, error) {
	hzw := &model.Hzw{}
	//result := dao.db.Find(hzw, key)
	result := dao.DB(ctx).Find(hzw, key)
	if result.Error != nil {
		return nil, result.Error
	}
	return hzw, nil
}
