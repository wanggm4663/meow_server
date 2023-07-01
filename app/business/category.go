package business

import (
	"github.com/gin-gonic/gin"
	"meow_server/app/model"
)

type CategoryBusiness struct {
}

func (s CategoryBusiness) GetCate1Tabs(c *gin.Context) ([]*model.Category, error) {
	data, err := model.Category{}.GetsByCondition(c, "*", map[string][]int{"pid": []int{0}})
	return data, err
}

func (s CategoryBusiness) GetCate2Tabs(c *gin.Context, cate1Id int) ([]*model.Category, error) {
	data, err := model.Category{}.GetsByCondition(c, "*", map[string][]int{"pid": []int{cate1Id}})
	return data, err
}
