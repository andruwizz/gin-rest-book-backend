package config

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDatabase(env *Env) *gorm.DB {
	host := env.DatabaseHost
	port := env.DatabasePort
	user := env.DatabaseUser
	pass := env.DatabasePassword
	name := env.DatabaseName

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, pass, host, port, name)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Database connection failed: %v", err)
	}

	_, err = db.DB()
	if err != nil {
		log.Fatalf("Database connection failed: %v", err)
	}

	return db
}
