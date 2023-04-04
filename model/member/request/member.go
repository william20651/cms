package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/member"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type MemberSearch struct{
    member.Member
    request.PageInfo
}
