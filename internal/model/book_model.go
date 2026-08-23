package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Book struct {
	Id        string `gorm:"type:uuid;primaryKey" json:"id"`
	Title     string `gorm:"type:varchar" json:"title"`
	Author    string `gorm:"type:varchar" json:"author"`
	CreatedAt string `gorm:"type:bigint" json:"created_at"`
	UpdatedAt string `gorm:"type:bigint" json:"updated_at"`
}

func (b *Book) BeforeCreate(tx *gorm.DB) error {
	uuid := uuid.New()
	time := time.Now().UnixMilli()
	tx.Statement.SetColumn("id", uuid)
	tx.Statement.SetColumn("CreatedAt", time)
	tx.Statement.SetColumn("UpdatedAt", time)

	return nil
}

func (b *Book) BeforeUpdate(tx *gorm.DB) error {
	time := time.Now().UnixMilli()
	tx.Statement.SetColumn("UpdatedAt", time)

	return nil
}

type BookRequest struct {
	Title  string `form:"title" binding:"required,alphanumspace"`
	Author string `form:"author" binding:"required,alphanumspace"`
}

type BookResponse struct {
	Id        string `json:"id"`
	Title     string `json:"title"`
	Author    string `json:"author"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
