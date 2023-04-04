// 自动生成模板Area
package area

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
)

// Area 结构体
type Area struct {
      global.GVA_MODEL
      AreaId  *int `json:"areaId" form:"areaId" gorm:"column:area_id;comment:;"`
      ParentId  *int `json:"parentId" form:"parentId" gorm:"column:parent_id;comment:;"`
      Name  string `json:"name" form:"name" gorm:"column:name;comment:;size:120;"`
}


// TableName Area 表名
func (Area) TableName() string {
  return "t_area"
}

