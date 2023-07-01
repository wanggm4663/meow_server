package model

import (
	"strconv"
	"time"
)

type Order struct {
	Id          int       `gorm:"column:id" db:"id" json:"id" form:"id"`
	Uid         int       `gorm:"column:uid" db:"uid" json:"uid" form:"uid"`                                     //用户id
	ProductId   int       `gorm:"column:product_id" db:"product_id" json:"product_id" form:"product_id"`         //商品id
	Price       int       `gorm:"column:price" db:"price" json:"price" form:"price"`                             //总金额
	RealPrice   int       `gorm:"column:real_price" db:"real_price" json:"real_price" form:"real_price"`         //实际支付金额
	OrderNumber string    `gorm:"column:order_number" db:"order_number" json:"order_number" form:"order_number"` //订单编号
	Content     string    `gorm:"column:content" db:"content" json:"content" form:"content"`                     //备注
	Status      int8      `gorm:"column:status" db:"status" json:"status" form:"status"`                         //状态 0 删除 1 待支付 2 支付中 3 已支付 4 已完成
	CreatedAt   time.Time `gorm:"column:created_at" db:"created_at" json:"created_at" form:"created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at" db:"updated_at" json:"updated_at" form:"updated_at"`
}

func (Order) TableName() string {
	return "order"
}

func (Order) GetsByUidAndStatus(field string, uid, status int) ([]*Order, error) {
	var data []*Order
	if uid <= 0 {
		return data, nil
	}
	where := make(map[string]string, 0)
	where["uid"] = strconv.Itoa(uid)
	if status != 0 {
		where["status"] = strconv.Itoa(status)
	}
	return data, DB.Model(Order{}).Select(field).Where(where).Find(&data).Error
}

func (Order) GetById(field string, id int) (*Order, error) {
	data := new(Order)
	if id <= 0 {
		return data, nil
	}
	where := make(map[string]string, 0)
	where["status"] = "1"
	where["id"] = strconv.Itoa(id)
	return data, DB.Model(Order{}).Select(field).Where(where).Find(&data).Error
}
