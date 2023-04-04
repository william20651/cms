package v1

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/area"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/example"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/member"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/system"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/tesss"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/test"
)

type ApiGroup struct {
	SystemApiGroup  system.ApiGroup
	ExampleApiGroup example.ApiGroup
	MemberApiGroup  member.ApiGroup
	TesssApiGroup   tesss.ApiGroup
	AreaApiGroup    area.ApiGroup
	TestApiGroup    test.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
