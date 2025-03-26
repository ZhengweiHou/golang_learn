package service

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"wiredemo/internal/repository/dao"
	"wiredemo/internal/repository/model"
)

type IHzwService interface {
	CreateHzw(ctx context.Context, hzw *model.Hzw) (*model.Hzw, error)
	CreateHzwTxTest(ctx context.Context, hzw *model.Hzw) ([]*model.Hzw, error)
	GetHzw(ctx context.Context, key int) (*model.Hzw, error)
	// CreateHzwWithTx(ctx context.Context, hzw *model.Hzw) (rhzws []*model.Hzw, err error)
}

type HzwService struct {
	*BaseService
	hzwDao dao.IHzwDao
}

// func NewHzwService(dao repository.IHzwDao) *HzwService {
func NewHzwService(bs *BaseService, dao dao.IHzwDao) IHzwService {
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
	switch hzw.Age {
	case 1:
		// 事务测试-函数封包方式
		return s.withTxFunc(ctx, hzw)
	case 2:
		// 事务测试-闭包回调方式
		return s.withTxCall(ctx, hzw)
	case 3:
		// 事务测试-闭包回调方式-多协程
		return s.withTxCallGoroutine(ctx, hzw)
	default:
		// 事务测试-闭包回调方式-多协程
		return s.withTxCallGoroutine(ctx, hzw)
	}

	return nil, nil
}

// 事务测试-函数封包方式
func (s *HzwService) withTxFunc(ctx context.Context, hzw *model.Hzw) (hzws []*model.Hzw, err error) {
	s.logger.Info("======事务测试-函数封包方式====")
	hzws2 := []*model.Hzw{}

	//db.TM.Transaction(
	// 事务1
	s.tm.Transaction(ctx, func(ctx context.Context) error {
		hzw1, err := s.hzwDao.InsertOne(ctx, hzw.Clone())
		hzws2 = append(hzws2, hzw1)
		s.logger.Info(fmt.Sprintf("hzw1:%v", hzw1))
		// 嵌套事务1
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

		// 嵌套事务2
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

// 事务测试-闭包回调方式
func (s *HzwService) withTxCall(ctx context.Context, hzw *model.Hzw) (rhzws []*model.Hzw, err error) {
	s.logger.Info("======事务测试-闭包回调方式====")
	rhzws = []*model.Hzw{}

	ctx, txcall := s.tm.WithTransaction(ctx)
	defer func() {
		txcall(err) // 异常传给txcall，使用匿名函数来捕获当前方法返回err的值
	}()

	// 无主键正常插入
	hzw1 := hzw.Clone()
	hzw1.Id = 0 // 操作一：不报错
	rhzw, err := s.hzwDao.InsertOne(ctx, hzw1)
	if err != nil {
		return nil, err
	}
	rhzws = append(rhzws, rhzw)

	// 存在主键插入，可能发生主键冲突异常，导致事务回滚
	hzw2 := hzw.Clone()
	hzw2.Id = 1
	rhzw, err = s.hzwDao.InsertOne(ctx, hzw2)
	if err != nil {
		return nil, err
	}
	rhzws = append(rhzws, rhzw)
	s.logger.Info(fmt.Sprintf("hzws:%v", rhzws))

	return

}

// 事务测试-闭包回调方式-多协程
func (s *HzwService) withTxCallGoroutine(ctx context.Context, hzw *model.Hzw) (hzws []*model.Hzw, rerr error) {
	s.logger.Info("======事务测试-闭包回调方式-多协程====")
	hzws = []*model.Hzw{}
	ctx, cancelfunc := context.WithCancel(ctx) // 使用cancel方式通知ctx
	ctx, txcall := s.tm.WithTransaction(ctx)
	defer func() {
		txcall(rerr) // 异常传给txcall，使用匿名函数来捕获当前方法返回err的值
	}()

	// 事务1操作1
	hzw1 := hzw.Clone()
	hzw1.Id = 0 // 操作一：不报错
	rhzw, rerr := s.hzwDao.InsertOne(ctx, hzw1)
	if rerr != nil {
		return nil, rerr
	}
	hzws = append(hzws, rhzw)

	// == 并发操作 ==
	concurrency := 3
	var swg, rwg sync.WaitGroup
	swg.Add(1)
	rwg.Add(concurrency + 1)

	go func() {
		defer rwg.Done()
		_hzw := hzw.Clone()
		_hzw.Id = 1 // 可能发生主键冲突异常
		_hzw.Name = fmt.Sprintf("goroutine_%v", "err")
		swg.Wait()

		rhzw, _err := s.hzwDao.InsertOne(ctx, _hzw)
		if _err != nil {
			//			time.Sleep(time.Millisecond * 100)
			rerr = _err  // 异常时父函数返回异常赋值
			cancelfunc() // 异常时通知ctx
			return
		}
		hzws = append(hzws, rhzw)
	}()

	for i := 0; i < concurrency; i++ {
		go func(i int) {
			defer rwg.Done()
			_hzw := hzw.Clone()
			_hzw.Name = fmt.Sprintf("goroutine_%d", i)
			swg.Wait()

			rhzw, _err := s.hzwDao.InsertOne(ctx, _hzw)
			if _err != nil {
				rerr = _err  // 异常时父函数返回异常赋值
				cancelfunc() // 异常时通知ctx
				return
			}
			hzws = append(hzws, rhzw)

		}(i)
	}

	swg.Done()
	// 等待所有goroutine完成
	rwg.Wait()

	//	fmt.Printf("err:%v\n", rerr)
	//	return hzws, rerr
	return
}
