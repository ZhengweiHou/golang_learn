package service

import (
	"context"
	"errors"
	"fmt"
	"wiredemo/internal/model"
	"wiredemo/internal/repository"
)

type IHzwService interface {
	CreateHzw(ctx context.Context, hzw *model.Hzw) (*model.Hzw, error)
	CreateHzwTxTest(ctx context.Context, hzw *model.Hzw) ([]*model.Hzw, error)
	GetHzw(ctx context.Context, key int) (*model.Hzw, error)
	CreateHzwWithTx(ctx context.Context, hzw *model.Hzw) (rhzws []*model.Hzw, err error)
}

type HzwService struct {
	*BaseService
	hzwDao repository.IHzwDao
}

// func NewHzwService(dao repository.IHzwDao) *HzwService {
func NewHzwService(bs *BaseService, dao repository.IHzwDao) IHzwService {
	return &HzwService{
		BaseService: bs,
		hzwDao:      dao,
	}
}

func (s *HzwService) CreateHzw(ctx context.Context, hzw *model.Hzw) (rhzw *model.Hzw, err error) {
	//return s.hzwDao.InsertOne(ctx, hzw)
	//db.TM.Transaction(
	s.tm.Transaction(
		ctx,
		func(ctx context.Context) error {
			rhzw, err = s.hzwDao.InsertOne(ctx, hzw)
			return err
		})
	return
}

func (s *HzwService) GetHzw(ctx context.Context, key int) (hzw *model.Hzw, err error) {
	//db.TM.Transaction(
	s.tm.Transaction(
		ctx,
		func(ctx context.Context) error {
			hzw, err = s.hzwDao.FindBykey(ctx, key)
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

func (s *HzwService) CreateHzwTxTest(ctx context.Context, hzw *model.Hzw) (hzws []*model.Hzw, err error) {
	hzws2 := []*model.Hzw{}

	//db.TM.Transaction(
	s.tm.Transaction(
		ctx,
		func(ctx context.Context) error {
			hzw1, err := s.hzwDao.InsertOne(ctx, hzw.Clone())
			hzws2 = append(hzws2, hzw1)
			s.logger.Info(fmt.Sprintf("hzw1:%v", hzw1))

			s.tm.Transaction(ctx, func(ctx context.Context) error {
				hzw2 := hzw.Clone()
				hzw2.Name = "222"
				hzw2.Age = 222
				hzw2, err := s.hzwDao.InsertOne(ctx, hzw2)
				if err != nil {
					return err
				}
				hzws2 = append(hzws2, hzw2)
				s.logger.Info(fmt.Sprintf("hzw2:%v", hzw2))

				return errors.New("rollback hzw2") // 模拟异常，回滚hzw2
				//return nil
			})

			s.tm.Transaction(ctx, func(ctx context.Context) error {
				hzw3 := hzw.Clone()
				hzw3.Name = "333"
				hzw3.Age = 333
				hzw3, err := s.hzwDao.InsertOne(ctx, hzw3)
				if err != nil {
					return err
				}
				hzws2 = append(hzws2, hzw3)
				s.logger.Info(fmt.Sprintf("hzw3:%v", hzw3))

				return nil
			})

			return err
		})

	return hzws2, nil
}

func (s *HzwService) CreateHzwWithTx(ctx context.Context, hzw *model.Hzw) (rhzws []*model.Hzw, err error) {
	rhzws = []*model.Hzw{}

	ctx, txcall := s.tm.WithTransaction(ctx)
	defer func() {
		txcall(err) // 异常传给txcall，使用匿名函数来捕获当前方法返回err的值
	}()

	hzw1 := hzw.Clone()
	hzw1.ID = 0 // 操作一：不报错
	rhzw, err := s.hzwDao.InsertOne(ctx, hzw1)
	if err != nil {
		return nil, err
	}
	rhzws = append(rhzws, rhzw)

	rhzw, err = s.hzwDao.InsertOne(ctx, hzw.Clone())
	if err != nil {
		return nil, err
	}
	rhzws = append(rhzws, rhzw)
	s.logger.Info(fmt.Sprintf("hzws:%v", rhzws))

	return
}
