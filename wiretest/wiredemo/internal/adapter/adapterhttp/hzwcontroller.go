package adapterhttp

import (
	"context"
	"strconv"
	"wiredemo/internal/repository/model"
	"wiredemo/internal/service"

	"github.com/gin-gonic/gin"
)

// NewHzwController hzw ctl构造函数
func NewHzwController(userService service.IHzwService) *HzwController {
	return &HzwController{
		userService: userService,
	}
}

// HzwController hzw controller
type HzwController struct {
	userService service.IHzwService
}

// SaveHzw 创建Hzw
// @Summary 创建Hzw
// @Description 创建新的Hzw对象
// @Tags hzw
// @Accept  json
// @Produce  json
// @Param hzw body model.Hzw true "Hzw信息"
// @Success 200 {object} model.Hzw
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /hzw [put]
func (ctl *HzwController) SaveHzw(ginc *gin.Context) {
	var hzw model.Hzw
	if err := ginc.ShouldBindJSON(&hzw); err != nil {
		ginc.JSON(400, gin.H{"error": err.Error()})
		return
	}
	ctx := context.Background()
	result, err := ctl.userService.CreateHzw(ctx, &hzw)
	if err != nil {
		ginc.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ginc.JSON(200, result)
}

// QueryById 根据ID查询Hzw
// @Summary 根据ID查询Hzw
// @Description 根据ID查询Hzw对象
// @Tags hzw
// @Accept  json
// @Produce  json
// @Param id query int true "Hzw ID"
// @Success 200 {object} model.Hzw
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /hzw [get]
func (ctl *HzwController) QueryById(ginc *gin.Context) {
	idstr, _ := ginc.GetQuery("id")
	id, err := strconv.Atoi(idstr)
	if err != nil {
		ginc.JSON(400, gin.H{"error": "invalid id"})
		return
	}

	ctx := context.Background()
	hzw, err := ctl.userService.GetHzw(ctx, id)
	if err != nil {
		ginc.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ginc.JSON(200, hzw)
}

// SaveHzwTxTest 创建Hzw 事务测试
// @Summary 创建Hzw
// @Description 创建新的Hzw对象
// @Tags hzw
// @Accept  json
// @Produce  json
// @Param hzw body model.Hzw true "Hzw信息"
// @Success 200 {object} model.Hzw
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /hzwtxtest [put]
func (ctl *HzwController) SaveHzwTxTest(ginc *gin.Context) {
	var hzw model.Hzw
	if err := ginc.ShouldBindJSON(&hzw); err != nil {
		ginc.JSON(400, gin.H{"error": err.Error()})
		return
	}
	ctx := context.Background()
	result, err := ctl.userService.CreateHzwTxTest(ctx, &hzw)
	if err != nil {
		ginc.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ginc.JSON(200, result)
}

/*
// SaveHzwWithTx 创建Hzw 事务测试
// @Summary 创建Hzw
// @Description 创建新的Hzw对象
// @Tags hzw
// @Accept  json
// @Produce  json
// @Param hzw body model.Hzw true "Hzw信息"
// @Success 200 {object} model.Hzw
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /hzwwithtx [put]
func (ctl *HzwController) SaveHzwWithTx(ginc *gin.Context) {
	var hzw model.Hzw
	if err := ginc.ShouldBindJSON(&hzw); err != nil {
		ginc.JSON(400, gin.H{"error": err.Error()})
		return
	}
	ctx := context.Background()
	result, err := ctl.userService.CreateHzwWithTx(ctx, &hzw)
	if err != nil {
		ginc.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ginc.JSON(200, result)

}
*/
