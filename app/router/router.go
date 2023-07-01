package router

import (
	"github.com/gin-gonic/gin"
	"meow_server/app/service"
)

func InitRouter(router *gin.Engine) {
	index := router.Group("/index")
	{
		index.GET("/cate1Tabs", func(c *gin.Context) {
			service.IndexService{}.Cate1Tabs(c)
		})
		index.GET("/cate2Tabs/:cate1_id", func(c *gin.Context) {
			service.IndexService{}.Cate2Tabs(c)
		})
		index.GET("/productInfos/:cate1_id", func(c *gin.Context) {
			service.IndexService{}.ProductInfos(c)
		})
	}
	product := router.Group("/product")
	{
		product.GET("/detail/:id", func(c *gin.Context) {
			service.ProductService{}.Detail(c)
		})
	}
	cart := router.Group("/cart")
	{
		cart.GET("/list", func(c *gin.Context) {
			service.CartService{}.List(c)
		})
	}
	order := router.Group("/order")
	{
		order.GET("/list", func(c *gin.Context) {
			service.OrderService{}.GetOrderList(c)
		})
		order.GET("/detail/:id", func(c *gin.Context) {
			service.OrderService{}.GetOrderDetail(c)
		})
		order.POST("/insert", func(c *gin.Context) {
		})
		order.POST("/update", func(c *gin.Context) {
		})
		order.DELETE("/delete", func(c *gin.Context) {
		})
	}
	address := router.Group("/address")
	{
		address.GET("/list", func(c *gin.Context) {
			service.AddressService{}.GetAddressList(c)
		})
		address.GET("/detail/:id", func(c *gin.Context) {
			service.AddressService{}.GetAddressDetail(c)
		})
		address.POST("/insert", func(c *gin.Context) {
		})
		address.POST("/update", func(c *gin.Context) {
		})
		address.DELETE("/delete", func(c *gin.Context) {
		})
	}
}
