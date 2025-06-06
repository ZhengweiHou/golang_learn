package service

import (
	"context"
	"wiredemo/internal/repository/dao"
	"wiredemo/internal/repository/model"
)

type IHzw2Service interface {
	CreateHzw2(ctx context.Context, hzw *model.Hzw2) (*model.Hzw2, error)
	GetHzw2(ctx context.Context, key int) (*model.Hzw2, error)
}

type Hzw2Service struct {
	*BaseService
	hzw2Dao dao.IHzw2Dao
}

func NewHzw2Service(bs *BaseService, dao dao.IHzw2Dao) IHzw2Service {
	return &Hzw2Service{
		BaseService: bs,
		hzw2Dao:     dao,
	}
}

func (s *Hzw2Service) CreateHzw2(ctx context.Context, hzw *model.Hzw2) (rhzw *model.Hzw2, err error) {
	//db.TM.Transaction(
	s.tm.Transaction(
		ctx,
		func(ctx context.Context) error {
			affect, err := s.hzw2Dao.InsertOne(ctx, hzw)
			s.logger.Info("affect", affect)
			return err
		})
	return hzw, nil
}

func (s *Hzw2Service) GetHzw2(ctx context.Context, key int) (hzw *model.Hzw2, err error) {
	//db.TM.Transaction(
	s.tm.Transaction(
		ctx,
		func(ctx context.Context) error {
			hzw, err = s.hzw2Dao.FindByPrimaryKey(ctx, uint(key))
			return err
		})
	//db.TM.Transaction2(
	//	ctx,
	//	func() error {
	//		hzw, err = s.hzwDao.FindBykey(ctx, key)
	//		return err
	//	})
	return
}
