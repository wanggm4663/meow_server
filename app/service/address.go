package service

import (
	"github.com/gin-gonic/gin"
	"meow_server/app/model"
	"meow_server/app/util"
	"strconv"
)

type AddressService struct {
}

func (s AddressService) GetAddressList(c *gin.Context) {
	uid := c.Query("uid")
	dbUid, _ := strconv.Atoi(uid)
	data, err := model.Address{}.GetsByUid("*", dbUid)
	util.Echo(c, data, err)
}

func (s AddressService) GetAddressDetail(c *gin.Context) {
	id := c.Param("id")
	dbId, _ := strconv.Atoi(id)
	data, err := model.Address{}.GetById("*", dbId)
	util.Echo(c, data, err)
}

func (s AddressService) InsertAddress(c *gin.Context) {

}

func (s AddressService) UpdateAddress(c *gin.Context) {

}

func (s AddressService) DeleteAddress(c *gin.Context) {

}
