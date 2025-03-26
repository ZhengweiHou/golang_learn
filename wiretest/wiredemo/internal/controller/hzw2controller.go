package controller

import (
	"context"
	"strconv"
	"wiredemo/internal/repository/model"
	"wiredemo/internal/service"

	"github.com/gin-gonic/gin"
)

// NewHzw2Controller hzw ctl构造函数
func NewHzw2Controller(userService service.IHzw2Service) *Hzw2Controller {
	return &Hzw2Controller{
		userService: userService,
	}
}

// Hzw2Controller hzw controller
type Hzw2Controller struct {
	userService service.IHzw2Service
}

// SaveHzw2 创建Hzw2
// @Summary 创建Hzw2
// @Description 创建新的Hzw2对象
// @Tags hzw
// @Accept  json
// @Produce  json
// @Param hzw body model.Hzw2 true "Hzw2信息"
// @Success 200 {object} model.Hzw2
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /hzw2 [put]
func (ctl *Hzw2Controller) SaveHzw2(ginc *gin.Context) {
	var hzw model.Hzw2
	if err := ginc.ShouldBindJSON(&hzw); err != nil {
		ginc.JSON(400, gin.H{"error": err.Error()})
		return
	}
	ctx := context.Background()
	result, err := ctl.userService.CreateHzw2(ctx, &hzw)
	if err != nil {
		ginc.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ginc.JSON(200, result)
}

// QueryById 根据ID查询Hzw2
// @Summary 根据ID查询Hzw2
// @Description 根据ID查询Hzw2对象
// @Tags hzw
// @Accept  json
// @Produce  json
// @Param id query int true "Hzw2 ID"
// @Success 200 {object} model.Hzw2
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /hzw2 [get]
func (ctl *Hzw2Controller) QueryById(ginc *gin.Context) {
	idstr, _ := ginc.GetQuery("id")
	id, err := strconv.Atoi(idstr)
	if err != nil {
		ginc.JSON(400, gin.H{"error": "invalid id"})
		return
	}

	ctx := context.Background()
	hzw, err := ctl.userService.GetHzw2(ctx, id)
	if err != nil {
		ginc.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ginc.JSON(200, hzw)
}

/*
// SaveHzw2TxTest 创建Hzw2 事务测试
// @Summary 创建Hzw2
// @Description 创建新的Hzw2对象
// @Tags hzw
// @Accept  json
// @Produce  json
// @Param hzw body model.Hzw2 true "Hzw2信息"
// @Success 200 {object} model.Hzw2
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /hzwtxtest [put]
func (ctl *Hzw2Controller) SaveHzw2TxTest(ginc *gin.Context) {
	var hzw model.Hzw2
	if err := ginc.ShouldBindJSON(&hzw); err != nil {
		ginc.JSON(400, gin.H{"error": err.Error()})
		return
	}
	ctx := context.Background()
	result, err := ctl.userService.CreateHzw2TxTest(ctx, &hzw)
	if err != nil {
		ginc.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ginc.JSON(200, result)
}
*/

/*
// SaveHzw2WithTx 创建Hzw2 事务测试
// @Summary 创建Hzw2
// @Description 创建新的Hzw2对象
// @Tags hzw
// @Accept  json
// @Produce  json
// @Param hzw body model.Hzw2 true "Hzw2信息"
// @Success 200 {object} model.Hzw2
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /hzwwithtx [put]
func (ctl *Hzw2Controller) SaveHzw2WithTx(ginc *gin.Context) {
	var hzw model.Hzw2
	if err := ginc.ShouldBindJSON(&hzw); err != nil {
		ginc.JSON(400, gin.H{"error": err.Error()})
		return
	}
	ctx := context.Background()
	result, err := ctl.userService.CreateHzw2WithTx(ctx, &hzw)
	if err != nil {
		ginc.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ginc.JSON(200, result)

}
*/
