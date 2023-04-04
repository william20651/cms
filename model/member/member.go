// 自动生成模板Member
package member

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// Member 结构体
type Member struct {
	global.GVA_MODEL
	Name          string    `json:"name" form:"name" gorm:"column:name;comment:;size:50;"`
	Realname      string    `json:"realname" form:"realname" gorm:"column:realname;comment:真实姓名;size:50;"`
	School1       string    `json:"school1" form:"school1" gorm:"column:school1;comment:毕业学校1;size:255;"`
	School2       string    `json:"school2" form:"school2" gorm:"column:school2;comment:;size:255;"`
	School3       string    `json:"school3" form:"school3" gorm:"column:school3;comment:;size:255;"`
	School4       string    `json:"school4" form:"school4" gorm:"column:school4;comment:;size:255;"`
	Born          time.Time `json:"born" form:"born" gorm:"type:datetime(0);column:born;comment:出生年月;size:6;"`
	Zodiac        uint      `json:"zodiac" form:"zodiac" gorm:"column:zodiac;comment:生肖;size:3;"`
	Constellation uint      `json:"constellation" form:"constellation" gorm:"column:constellation;comment:星座;size:3;"`
	Gender        uint      `json:"gender" form:"gender" gorm:"column:gender;comment:0男1女2保密;size:3;"`
	Degree        uint      `json:"degree" form:"degree" gorm:"column:degree;comment:学历,0高中以下，1中专2大专，3本科4硕士5博士;size:3;"`
	Marital       uint      `json:"marital" form:"marital" gorm:"column:marital;comment:0未婚1离异2丧偶;size:3;"`
	Vip           uint      `json:"vip" form:"vip" gorm:"column:vip;comment:0免费1付费;size:3;"`
	Job           uint      `json:"job" form:"job" gorm:"column:job;comment:职业;size:3;"`
	Stat          uint      `json:"stat" form:"stat" gorm:"column:stat;comment:0正常1结束（已找到）-1del;size:3;"`
	Phone         string    `json:"phone" form:"phone" gorm:"column:phone;comment:;size:255;"`
	Wechat        string    `json:"wechat" form:"wechat" gorm:"column:wechat;comment:;size:255;"`
	Photo         string    `json:"photo" form:"photo" gorm:"column:photo;comment:头像;size:255;"`
	Img1          string    `json:"img1" form:"img1" gorm:"column:img1;comment:;size:255;"`
	Img2          string    `json:"img2" form:"img2" gorm:"column:img2;comment:;size:255;"`
	Img3          string    `json:"img3" form:"img3" gorm:"column:img3;comment:;size:255;"`
	Hight         string    `json:"hight" form:"hight" gorm:"column:hight;comment:身高厘米;size:6;"`
	Weight        string    `json:"weight" form:"weight" gorm:"column:weight;comment:体重kg;size:6;"`
	NativePlace   uint      `json:"nativePlace" form:"nativePlace" gorm:"column:native_place;comment:籍贯;size:4;"`
	HasHouse      uint      `json:"hasHouse" form:"hasHouse" gorm:"column:has_house;comment:是否有房;size:3;"`
	HasCar        uint      `json:"hasCar" form:"hasCar" gorm:"column:has_car;comment:是否有车;size:3;"`
	Ext           string    `json:"ext" form:"ext" gorm:"column:ext;comment:房车补充;size:255;"`
	Interests     string    `json:"interests" form:"interests" gorm:"column:interests;comment:兴趣爱好;size:255;"`
	Requires      string    `json:"requires" form:"requires" gorm:"column:requires;comment:择偶要求;size:255;"`
	Desc          string    `json:"desc" form:"desc" gorm:"column:desc;comment:个人描述;size:255;"`
}

// TableName Member 表名
func (Member) TableName() string {
	return "t_member"
}
