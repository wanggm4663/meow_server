package service

import (
	"github.com/gin-gonic/gin"
	"meow_server/app/business"
	"meow_server/app/util"
	"strconv"
)

type IndexService struct {
}

func (s IndexService) Cate1Tabs(c *gin.Context) {
	data, err := business.CategoryBusiness{}.GetCate1Tabs(c)
	util.Echo(c, data, err)
}

func (s IndexService) Cate2Tabs(c *gin.Context) {
	cate1Id := c.Param("cate1_id")
	dbCate1Id, _ := strconv.Atoi(cate1Id)
	data, err := business.CategoryBusiness{}.GetCate2Tabs(c, dbCate1Id)
	util.Echo(c, data, err)
}

func (s IndexService) ProductInfos(c *gin.Context) {
	cate1Id := c.Param("cate1_id")
	dbCate1Id, _ := strconv.Atoi(cate1Id)
	data, err := business.ProductBusiness{}.GetProductInfos(c, dbCate1Id)
	util.Echo(c, data, err)
}
