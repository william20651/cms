package service

import (
	"github.com/flipped-aurora/gin-vue-admin/server/service/area"
	"github.com/flipped-aurora/gin-vue-admin/server/service/example"
	"github.com/flipped-aurora/gin-vue-admin/server/service/member"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
	"github.com/flipped-aurora/gin-vue-admin/server/service/tesss"
	"github.com/flipped-aurora/gin-vue-admin/server/service/test"
)

type ServiceGroup struct {
	SystemServiceGroup  system.ServiceGroup
	ExampleServiceGroup example.ServiceGroup
	MemberServiceGroup  member.ServiceGroup
	TesssServiceGroup   tesss.ServiceGroup
	AreaServiceGroup    area.ServiceGroup
	TestServiceGroup    test.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
