package config

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDatabase() *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", "db_user", "db_password", "db_host", 3306, "db_name")

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
