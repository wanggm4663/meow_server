package res

import "meow_server/app/model"

type ProductInfos struct {
	List []*ProductInfo
}

type ProductInfo struct {
	Cate2Id   int              `json:"cate_2_id"`
	Cate2Name string           `json:"cate_2_name"`
	List      []*model.Product `json:"list"`
}

type ProductDetail struct {
	ProductId int    `json:"product_id"`
	Cate1Id   int    `json:"cate_1_id"`
	Cate2Id   int    `json:"cate_2_id"`
	Price     int    `json:"price"`
	Title     string `json:"title"`
	Img       string `json:"img"`
}
