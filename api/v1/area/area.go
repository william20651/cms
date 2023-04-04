package area

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/area"
	areaReq "github.com/flipped-aurora/gin-vue-admin/server/model/area/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type AreaApi struct {
}

var areaService = service.ServiceGroupApp.AreaServiceGroup.AreaService

// CreateArea 创建Area
// @Tags Area
// @Summary 创建Area
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body area.Area true "创建Area"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /area/createArea [post]
func (areaApi *AreaApi) CreateArea(c *gin.Context) {
	var area area.Area
	_ = c.ShouldBindJSON(&area)
	if err := areaService.CreateArea(area); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteArea 删除Area
// @Tags Area
// @Summary 删除Area
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body area.Area true "删除Area"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /area/deleteArea [delete]
func (areaApi *AreaApi) DeleteArea(c *gin.Context) {
	var area area.Area
	_ = c.ShouldBindJSON(&area)
	if err := areaService.DeleteArea(area); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteAreaByIds 批量删除Area
// @Tags Area
// @Summary 批量删除Area
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Area"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /area/deleteAreaByIds [delete]
func (areaApi *AreaApi) DeleteAreaByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := areaService.DeleteAreaByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateArea 更新Area
// @Tags Area
// @Summary 更新Area
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body area.Area true "更新Area"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /area/updateArea [put]
func (areaApi *AreaApi) UpdateArea(c *gin.Context) {
	var area area.Area
	_ = c.ShouldBindJSON(&area)
	if err := areaService.UpdateArea(area); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindArea 用id查询Area
// @Tags Area
// @Summary 用id查询Area
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query area.Area true "用id查询Area"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /area/findArea [get]
func (areaApi *AreaApi) FindArea(c *gin.Context) {
	var area area.Area
	_ = c.ShouldBindQuery(&area)
	if rearea, err := areaService.GetArea(area.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rearea": rearea}, c)
	}
}

// GetAreaList 分页获取Area列表
// @Tags Area
// @Summary 分页获取Area列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query areaReq.AreaSearch true "分页获取Area列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /area/getAreaList [get]
func (areaApi *AreaApi) GetAreaList(c *gin.Context) {
	var pageInfo areaReq.AreaSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list, total, err := areaService.GetAreaInfoList(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}

// GetAreaListTop 获取parent_id=0 的Area列表
// @Tags Area
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query areaReq.AreaSearch true "获取parent_id=0 的Area列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /area/getAreaListTop [get]
func (areaApi *AreaApi) GetAreaListTop(c *gin.Context) {
	if list, total, err := areaService.GetAreaInfoListTop(); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:  list,
			Total: total,
		}, "获取成功", c)
	}
}
