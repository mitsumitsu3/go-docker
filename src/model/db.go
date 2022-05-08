package model

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func init() {

	connect := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_DATABASE"))

	DB, err = gorm.Open(mysql.Open(connect), &gorm.Config{})
	if err != nil {
		log.Fatalln(connect + "database can't connect")
	}
	DB.AutoMigrate(&User{})
}
