package config

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDatabase() *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", "user", "password", "127.0.0.1", 3306, "app")

	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		log.Fatalf("Database connection failed: %v", err)
	}

	_, err = db.DB()
	if err != nil {
		log.Fatalf("Database connection failed: %v", err)
	}

	return db
}
