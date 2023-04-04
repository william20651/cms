package member

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/member"
	memberReq "github.com/flipped-aurora/gin-vue-admin/server/model/member/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type MemberApi struct {
}

var memberService = service.ServiceGroupApp.MemberServiceGroup.MemberService

// CreateMember 创建Member
// @Tags Member
// @Summary 创建Member
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body member.Member true "创建Member"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /member/createMember [post]
func (memberApi *MemberApi) CreateMember(c *gin.Context) {
	var member member.Member
	_ = c.ShouldBindJSON(&member)
	if err := memberService.CreateMember(member); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteMember 删除Member
// @Tags Member
// @Summary 删除Member
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body member.Member true "删除Member"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /member/deleteMember [delete]
func (memberApi *MemberApi) DeleteMember(c *gin.Context) {
	var member member.Member
	_ = c.ShouldBindJSON(&member)
	if err := memberService.DeleteMember(member); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteMemberByIds 批量删除Member
// @Tags Member
// @Summary 批量删除Member
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Member"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /member/deleteMemberByIds [delete]
func (memberApi *MemberApi) DeleteMemberByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := memberService.DeleteMemberByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateMember 更新Member
// @Tags Member
// @Summary 更新Member
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body member.Member true "更新Member"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /member/updateMember [put]
func (memberApi *MemberApi) UpdateMember(c *gin.Context) {
	var member member.Member

	_ = c.ShouldBindJSON(&member)
	// global.GVA_LOG.Info("member is  ",
	// 	zap.String("hight", member.Hight),
	// 	zap.String("weight", member.Weight),
	// 	zap.String("job", strconv.Itoa(int(member.Job))),
	// )
	if err := memberService.UpdateMember(member); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindMember 用id查询Member
// @Tags Member
// @Summary 用id查询Member
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query member.Member true "用id查询Member"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /member/findMember [get]
func (memberApi *MemberApi) FindMember(c *gin.Context) {
	var member member.Member
	_ = c.ShouldBindQuery(&member)
	if remember, err := memberService.GetMember(member.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"remember": remember}, c)
	}
}

// GetMemberList 分页获取Member列表
// @Tags Member
// @Summary 分页获取Member列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query memberReq.MemberSearch true "分页获取Member列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /member/getMemberList [get]
func (memberApi *MemberApi) GetMemberList(c *gin.Context) {
	var pageInfo memberReq.MemberSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list, total, err := memberService.GetMemberInfoList(pageInfo); err != nil {
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
