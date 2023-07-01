package service

import (
	"github.com/gin-gonic/gin"
	"meow_server/app/business"
	"meow_server/app/util"
	"strconv"
)

type ProductService struct {
}

func (s ProductService) Detail(c *gin.Context) {
	id := c.Param("id")
	dbId, _ := strconv.Atoi(id)
	data, err := business.ProductBusiness{}.GetProductDetail(c, dbId)
	util.Echo(c, data, err)
}
