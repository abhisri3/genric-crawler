package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var urlDSN = "root:root@tcp(localhost:3306)/insuranceDB?charset=utf8mb4&parseTime=True&loc=Local"
var err error



func DataMigration() {
	DB, err = gorm.Open(mysql.Open(urlDSN),&gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("Connection error")
	}
	DB.AutoMigrate(&Search{})
	DB.AutoMigrate(&InsuranceData{})
}
