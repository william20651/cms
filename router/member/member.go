package member

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type MemberRouter struct {
}

// InitMemberRouter 初始化 Member 路由信息
func (s *MemberRouter) InitMemberRouter(Router *gin.RouterGroup) {
	memberRouter := Router.Group("member").Use(middleware.OperationRecord())
	memberRouterWithoutRecord := Router.Group("member")
	var memberApi = v1.ApiGroupApp.MemberApiGroup.MemberApi
	{
		memberRouter.POST("createMember", memberApi.CreateMember)   // 新建Member
		memberRouter.DELETE("deleteMember", memberApi.DeleteMember) // 删除Member
		memberRouter.DELETE("deleteMemberByIds", memberApi.DeleteMemberByIds) // 批量删除Member
		memberRouter.PUT("updateMember", memberApi.UpdateMember)    // 更新Member
	}
	{
		memberRouterWithoutRecord.GET("findMember", memberApi.FindMember)        // 根据ID获取Member
		memberRouterWithoutRecord.GET("getMemberList", memberApi.GetMemberList)  // 获取Member列表
	}
}
