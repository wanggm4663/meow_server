package model

import (
	"github.com/gin-gonic/gin"
	"time"
)

type Category struct {
	Id        int       `gorm:"column:id" db:"id" json:"id" form:"id"`
	Pid       int       `gorm:"column:pid" db:"pid" json:"pid" form:"pid"`     //上级分类id
	Name      string    `gorm:"column:name" db:"name" json:"name" form:"name"` //类目名称
	NameEn    string    `gorm:"column:name_en" db:"name_en" json:"name_en" form:"name_en"`
	Status    int8      `gorm:"column:status" db:"status" json:"status" form:"status"` //状态 0 删除 1 正常
	CreatedAt time.Time `gorm:"column:created_at" db:"created_at" json:"created_at" form:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" db:"updated_at" json:"updated_at" form:"updated_at"`
}

func (Category) TableName() string {
	return "category"
}

func (Category) GetsByCondition(c *gin.Context, field string, cond map[string][]int) ([]*Category, error) {
	var data []*Category
	db := DB.Model(Category{}).Select(field)
	if cond["pid"] != nil {
		if len(cond["pid"]) > 1 {
			db.Where("pid in ?", cond["pid"])
		} else {
			db.Where("pid = ?", cond["pid"][0])
		}
	}
	db.Where("status = ?", 1)
	return data, db.Debug().Find(&data).Error
}
