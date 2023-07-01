package res

type CartList struct {
	List  []*Cart `json:"list"`
	Price int     `json:"price"`
}

type Cart struct {
	ProductId int    `json:"product_id"`
	Count     int    `json:"count"`
	Title     string `json:"title"`
	Img       string `json:"img"`
}
