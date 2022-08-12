package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDB() *gorm.DB {
	db, err := gorm.Open(mysql.New(mysql.Config{DSN: "admin:admin@tcp(127.0.0.1:3306)/example?charset=utf8mb4"}))
	if err != nil {
		panic(err)
	}

	return db
}
