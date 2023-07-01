package model

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

type Cart struct {
	Id        int       `gorm:"column:id" db:"id" json:"id" form:"id"`
	Uid       int       `gorm:"column:uid" db:"uid" json:"uid" form:"uid"`                             //用户id
	ProductId int       `gorm:"column:product_id" db:"product_id" json:"product_id" form:"product_id"` //商品id
	Count     int       `gorm:"column:count" db:"count" json:"count" form:"count"`                     //商品数量
	CreatedAt time.Time `gorm:"column:created_at" db:"created_at" json:"created_at" form:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" db:"updated_at" json:"updated_at" form:"updated_at"`
}

func (Cart) TableName() string {
	return "cart"
}

func (Cart) GetsByUid(c *gin.Context, field string, uid int) ([]*Cart, error) {
	var data []*Cart
	if uid <= 0 {
		return data, nil
	}
	where := make(map[string]string)
	where["uid"] = strconv.Itoa(uid)
	return data, DB.Model(Cart{}).Select(field).Where(where).Find(&data).Error
}
