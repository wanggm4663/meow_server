package service

import (
	"github.com/gin-gonic/gin"
	"meow_server/app/model"
	"meow_server/app/util"
	"strconv"
)

type OrderService struct {
}

func (s OrderService) GetOrderList(c *gin.Context) {
	uid := c.Query("uid")
	status := c.Query("status")
	dbUid, _ := strconv.Atoi(uid)
	dbStatus, _ := strconv.Atoi(status)
	data, err := model.Order{}.GetsByUidAndStatus("*", dbUid, dbStatus)
	util.Echo(c, data, err)
}

func (s OrderService) GetOrderDetail(c *gin.Context) {
	id := c.Param("id")
	dbId, _ := strconv.Atoi(id)
	data, err := model.Order{}.GetById("*", dbId)
	util.Echo(c, data, err)
}

func (s OrderService) InsertOrder(c *gin.Context) {

}

func (s OrderService) UpdateOrder(c *gin.Context) {

}

func (s OrderService) DeleteOrder(c *gin.Context) {

}
