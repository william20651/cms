package global

import (
	"time"

	"gorm.io/gorm"
)

type GVA_MODEL struct {
	ID        uint           `gorm:"primarykey"`       // 主键ID
	CreatedAt time.Time      `gorm:"type:datetime(0)"` // 创建时间
	UpdatedAt time.Time      `gorm:"type:datetime(0)"` // 更新时间
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`   // 删除时间
}
