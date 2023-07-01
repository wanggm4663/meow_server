package service

import (
	"github.com/gin-gonic/gin"
	"meow_server/app/business"
	"meow_server/app/util"
)

type CartService struct {
}

func (s CartService) List(c *gin.Context) {
	uid := 1
	data, err := business.CartBusiness{}.GetCartList(c, uid)
	util.Echo(c, data, err)
}
