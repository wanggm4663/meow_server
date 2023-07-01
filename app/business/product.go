package business

import (
	"github.com/gin-gonic/gin"
	"meow_server/app/model"
	"meow_server/app/model/res"
)

type ProductBusiness struct {
}

func (s ProductBusiness) GetProductInfos(c *gin.Context, cate1Id int) (*res.ProductInfos, error) {
	data := new(res.ProductInfos)
	cate2List, _ := CategoryBusiness{}.GetCate2Tabs(c, cate1Id)
	productList, _ := model.Product{}.GetsByCondition(c, "*", map[string][]int{"cate1_id": []int{cate1Id}})
	productMap := make(map[int][]*model.Product)
	if len(productList) > 0 {
		for _, v := range productList {
			productMap[v.Cate2Id] = append(productMap[v.Cate2Id], v)
		}
	}
	for _, v := range cate2List {
		productInfo := &res.ProductInfo{
			Cate2Id:   v.Id,
			Cate2Name: v.Name,
			List:      productMap[v.Id],
		}
		data.List = append(data.List, productInfo)
	}
	return data, nil
}

func (s ProductBusiness) GetProductDetail(c *gin.Context, id int) (*res.ProductDetail, error) {
	data := new(res.ProductDetail)
	product, err := model.Product{}.GetByCondition(c, "*", map[string][]int{"id": []int{id}})
	if product.Id > 0 {
		data.ProductId = product.Id
		data.Cate1Id = product.Cate1Id
		data.Cate2Id = product.Cate2Id
		data.Price = product.Price
		data.Title = product.Title
		data.Img = product.Img
	}
	return data, err
}
