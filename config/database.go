package config

import (
	"rest-api-golang-pemula/structs"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {

	dsn := "root:amaterasu@tcp(127.0.0.1:3306)/hactiv8-rest?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn))

	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&structs.Person{})

	return db

}
