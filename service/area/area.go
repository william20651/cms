package area

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/area"
	areaReq "github.com/flipped-aurora/gin-vue-admin/server/model/area/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type AreaService struct {
}

// CreateArea 创建Area记录
// Author [piexlmax](https://github.com/piexlmax)
func (areaService *AreaService) CreateArea(area area.Area) (err error) {
	err = global.GVA_DB.Create(&area).Error
	return err
}

// DeleteArea 删除Area记录
// Author [piexlmax](https://github.com/piexlmax)
func (areaService *AreaService) DeleteArea(area area.Area) (err error) {
	err = global.GVA_DB.Delete(&area).Error
	return err
}

// DeleteAreaByIds 批量删除Area记录
// Author [piexlmax](https://github.com/piexlmax)
func (areaService *AreaService) DeleteAreaByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]area.Area{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateArea 更新Area记录
// Author [piexlmax](https://github.com/piexlmax)
func (areaService *AreaService) UpdateArea(area area.Area) (err error) {
	err = global.GVA_DB.Save(&area).Error
	return err
}

// GetArea 根据id获取Area记录
// Author [piexlmax](https://github.com/piexlmax)
func (areaService *AreaService) GetArea(id uint) (area area.Area, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&area).Error
	return
}

// GetAreaInfoList 分页获取Area记录
// Author [piexlmax](https://github.com/piexlmax)
func (areaService *AreaService) GetAreaInfoList(info areaReq.AreaSearch) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&area.Area{})
	var areas []area.Area
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&areas).Error
	return areas, total, err
}

//GetAreaInfoListTop 获取parent_id=0的区域

func (areaService *AreaService) GetAreaInfoListTop() (list interface{}, total int64, err error) {

	// 创建db
	db := global.GVA_DB.Model(&area.Area{})
	var areas []area.Area
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Where("parent_id = ?", 0).Find(&areas).Error
	return areas, total, err
}
