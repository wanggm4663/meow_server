package business

import (
	"github.com/gin-gonic/gin"
	"meow_server/app/model"
	"meow_server/app/model/res"
)

type CartBusiness struct {
}

func (s CartBusiness) GetCartList(c *gin.Context, uid int) (*res.CartList, error) {
	data := new(res.CartList)
	carts, err := model.Cart{}.GetsByUid(c, "product_id,count", uid)
	if err != nil || len(carts) == 0 {
		return data, err
	}
	productIds := make([]int, 0, len(carts))
	for _, v := range carts {
		productIds = append(productIds, v.ProductId)
	}
	productList, _ := model.Product{}.GetsByCondition(c, "id,title,price,img", map[string][]int{"id": productIds})
	if len(productList) > 0 {
		productMap := make(map[int]*model.Product)
		for _, v := range productList {
			productMap[v.Id] = v
		}
		price := 0
		for _, v := range carts {
			cur := &res.Cart{
				ProductId: v.ProductId,
				Count:     v.Count,
				Title:     productMap[v.ProductId].Title,
				Img:       productMap[v.ProductId].Img,
			}
			curPrice := v.Count * productMap[v.ProductId].Price
			price += curPrice
			data.List = append(data.List, cur)
		}
		data.Price = price
	}
	return data, err
}
