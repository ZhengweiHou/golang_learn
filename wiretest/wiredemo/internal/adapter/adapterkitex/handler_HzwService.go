package adapterkitex

import (
	"context"
	"fmt"
	"log/slog"
	hzw "wiredemo/api/kitex/hzw"
	"wiredemo/internal/repository/model"
	"wiredemo/internal/service"
)

// HzwKitexCtl implements the last service interface defined in the IDL.
//
//	实现 wiredemo/api/kitex/hzw.HzwService
type HzwKitexCtl struct {
	hzwsvs service.IHzwService
}

func NewHzwKitexCtl(hzwsvs service.IHzwService) hzw.HzwService { // 实现 wiredemo/api/kitex/hzw.HzwService，TODO 构造器是返回接口还是实现
	s := &HzwKitexCtl{
		hzwsvs: hzwsvs,
	}
	return s
}

// CreateHzw implements the HzwServiceImpl interface.
func (s *HzwKitexCtl) CreateHzw(ctx context.Context, hzwDto *hzw.HzwDto) (resp *hzw.HzwDto, err error) {
	slog.Info("CreateHzw")
	mhzw := HzwDto2ModeHzw(hzwDto)
	rhzw, err := s.hzwsvs.CreateHzw(ctx, mhzw)
	if err != nil {
		return nil, err
	}
	fmt.Println(rhzw.CreatedAt.UnixMilli())
	resp = ModeHzw2HzwDto(rhzw)
	return resp, nil
}

// GetHzw implements the HzwServiceImpl interface.
func (s *HzwKitexCtl) GetHzw(ctx context.Context, id int64) (resp *hzw.HzwDto, err error) {
	// 示例实现 - 实际应根据业务需求修改
	if id <= 0 {
		return nil, fmt.Errorf("invalid id")
	}
	return &hzw.HzwDto{
		Id:   id,
		Name: "示例名称",
	}, nil
}

// CreateHzwTxTest implements the HzwServiceImpl interface.
func (s *HzwKitexCtl) CreateHzwTxTest(ctx context.Context, hzwDto *hzw.HzwDto) (resp *hzw.HzwDto, err error) {
	// 示例实现 - 实际应根据业务需求修改
	resp = &hzw.HzwDto{
		Id:   hzwDto.Id,
		Name: hzwDto.Name,
	}
	return resp, nil
}

func HzwDto2ModeHzw(hzwDto *hzw.HzwDto) *model.Hzw {
	if hzwDto == nil {
		return nil
	}

	return &model.Hzw{
		Id:      int32(hzwDto.Id),
		Name:    hzwDto.Name,
		Age:     int(hzwDto.Age),
		Version: hzwDto.Version,
		//CreatedAt: time.Unix(hzwDto.CreatedAt, 0), // TODO: 时间转换，默认0会转换为1970.1.1 8:0:0
		//UpdatedAt: time.UnixMilli(hzwDto.UpdatedAt),
		//Time1:     time.Unix(hzwDto.Time1, 0),
		//Time2:     time.Unix(hzwDto.Time2, 0),
		//Time3:     time.Unix(hzwDto.Time3, 0),
		Decimal1: hzwDto.Decimal1,
	}
}

func ModeHzw2HzwDto(mhzw *model.Hzw) *hzw.HzwDto {
	if mhzw == nil {
		return nil
	}

	return &hzw.HzwDto{
		Id:        int64(mhzw.Id),
		Name:      mhzw.Name,
		Age:       int32(mhzw.Age),
		Version:   mhzw.Version,
		CreatedAt: mhzw.CreatedAt.Unix(),
		UpdatedAt: mhzw.UpdatedAt.Unix(),
		Time1:     mhzw.Time1.Unix(),
		Time2:     mhzw.Time2.Unix(),
		Time3:     mhzw.Time3.Unix(),
		Decimal1:  mhzw.Decimal1,
	}
}
