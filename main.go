package main

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/ini.v1"
	"log"
	"meow_server/app/model"
	"meow_server/app/router"
)

func main() {
	r := gin.Default()
	config, err := ini.Load("./app/config/my.ini")
	if err != nil {
		log.Fatal("读取配置文件失败，", err)
	}
	router.InitRouter(r)
	model.InitDB(config)
	r.Run(":8000")
}
