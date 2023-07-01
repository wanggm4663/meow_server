package model

import (
	"github.com/gin-gonic/gin"
	"time"
)

type Product struct {
	Id        int       `gorm:"column:id" db:"id" json:"id" form:"id"`
	Cate1Id   int       `gorm:"column:cate1_id" db:"cate1_id" json:"cate1_id" form:"cate1_id"` //一级类目id
	Cate2Id   int       `gorm:"column:cate2_id" db:"cate2_id" json:"cate2_id" form:"cate2_id"` //二级类目id
	Title     string    `gorm:"column:title" db:"title" json:"title" form:"title"`             //商品名称
	Price     int       `gorm:"column:price" db:"price" json:"price" form:"price"`             //价格
	Img       string    `gorm:"column:img" db:"img" json:"img" form:"img"`                     //头图
	Status    int8      `gorm:"column:status" db:"status" json:"status" form:"status"`         //状态 0 删除 1正常
	CreatedAt time.Time `gorm:"column:created_at" db:"created_at" json:"created_at" form:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" db:"updated_at" json:"updated_at" form:"updated_at"`
}

func (Product) TableName() string {
	return "product"
}

func (Product) GetsByCondition(c *gin.Context, field string, cond map[string][]int) ([]*Product, error) {
	var data []*Product
	db := DB.Model(Product{}).Select(field)
	if cond["id"] != nil {
		if len(cond["id"]) > 1 {
			db.Where("id in ?", cond["id"])
		} else {
			db.Where("id = ?", cond["id"][0])
		}
	}
	if cond["cate1_id"] != nil {
		if len(cond["cate1_id"]) > 1 {
			db.Where("cate1_id in ?", cond["cate1_id"])
		} else {
			db.Where("cate1_id = ?", cond["cate1_id"][0])
		}
	}
	db.Where("status = ?", 1)
	return data, db.Debug().Find(&data).Error
}

func (Product) GetByCondition(c *gin.Context, field string, cond map[string][]int) (*Product, error) {
	var data *Product
	db := DB.Model(Product{}).Select(field)
	if cond["id"] != nil {
		if len(cond["id"]) > 1 {
			db.Where("id in ?", cond["id"])
		} else {
			db.Where("id = ?", cond["id"][0])
		}
	}
	db.Where("status = ?", 1)
	return data, db.Debug().Take(&data).Error
}
