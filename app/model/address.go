package model

import (
	"strconv"
	"time"
)

type Address struct {
	Id          int       `gorm:"column:id" db:"id" json:"id" form:"id"`
	Uid         int       `gorm:"column:uid" db:"uid" json:"uid" form:"uid"`                                     //用户id
	UserName    string    `gorm:"column:user_name" db:"user_name" json:"user_name" form:"user_name"`             //收件人
	PhoneNumber int       `gorm:"column:phone_number" db:"phone_number" json:"phone_number" form:"phone_number"` //手机号码
	Province    int       `gorm:"column:province" db:"province" json:"province" form:"province"`                 //省
	City        int       `gorm:"column:city" db:"city" json:"city" form:"city"`                                 //市
	Area        int       `gorm:"column:area" db:"area" json:"area" form:"area"`                                 //区
	Content     string    `gorm:"column:content" db:"content" json:"content" form:"content"`                     //详细地址
	Status      int8      `gorm:"column:status" db:"status" json:"status" form:"status"`                         //状态 0 删除 1 正常
	CreatedAt   time.Time `gorm:"column:created_at" db:"created_at" json:"created_at" form:"created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at" db:"updated_at" json:"updated_at" form:"updated_at"`
}

func (Address) TableName() string {
	return "address"
}

func (Address) GetById(field string, id int) (*Address, error) {
	data := new(Address)
	if id <= 0 {
		return data, nil
	}
	where := make(map[string]string, 0)
	where["status"] = "1"
	where["id"] = strconv.Itoa(id)
	return data, DB.Model(Address{}).Select(field).Where(where).Find(&data).Error
}

func (Address) GetsByUid(field string, uid int) ([]*Address, error) {
	var data []*Address
	if uid <= 0 {
		return data, nil
	}
	where := make(map[string]string, 0)
	where["status"] = "1"
	where["uid"] = strconv.Itoa(uid)
	return data, DB.Model(Address{}).Select(field).Where(where).Find(&data).Error
}
