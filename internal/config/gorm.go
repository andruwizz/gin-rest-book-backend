package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDatabase() *gorm.DB {
	host := os.Getenv("DATABASE_HOST")
	port, _ := strconv.Atoi(os.Getenv("DATABASE_PORT"))
	user := os.Getenv("DATABASE_USER")
	pass := os.Getenv("DATABASE_PASSWORD")
	name := os.Getenv("DATABASE_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, pass, host, port, name)

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
