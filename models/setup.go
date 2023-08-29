package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func KonekDB() {
	db, err := gorm.Open(mysql.Open("root:admin@tcp(localhost:3306)/jwt_login"))
	if err != nil {
		panic(err.Error())
	}

	db.AutoMigrate(&User{})

	DB = db
}
