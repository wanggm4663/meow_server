package model

import (
	"fmt"
	"gopkg.in/ini.v1"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func InitDB(config *ini.File) *gorm.DB {
	var err error
	dbHost := config.Section("mysql").Key("db_host").String()
	dbPort := config.Section("mysql").Key("db_port").String()
	dbUser := config.Section("mysql").Key("db_user").String()
	dbPassword := config.Section("mysql").Key("db_password").String()
	dbName := config.Section("mysql").Key("db_name").String()
	DB, err = gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPassword, dbHost, dbPort, dbName)), &gorm.Config{})
	if err != nil {
		log.Fatal("连接数据库失败，", err)
	}
	return DB
}
