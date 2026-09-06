package model

import (
	"time"

	"github.com/andruwizz/gin-rest-book-backend/internal/entity"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Book struct {
	Id        string `gorm:"type:uuid;primaryKey"`
	Title     string `gorm:"type:varchar"`
	Author    string `gorm:"type:varchar"`
	CreatedAt int64  `gorm:"type:bigint"`
	UpdatedAt int64  `gorm:"type:bigint"`
}

func (m *Book) BeforeCreate(tx *gorm.DB) error {
	newUUID := uuid.New()
	currentTime := time.Now().UnixMilli()
	tx.Statement.SetColumn("id", newUUID)
	tx.Statement.SetColumn("CreatedAt", currentTime)
	tx.Statement.SetColumn("UpdatedAt", currentTime)

	return nil
}

func (m *Book) BeforeUpdate(tx *gorm.DB) error {
	currentTime := time.Now().UnixMilli()
	tx.Statement.SetColumn("UpdatedAt", currentTime)

	return nil
}

func FromBookEntity(e *entity.Book) *Book {
	return &Book{
		Id:        e.Id,
		Title:     e.Title,
		Author:    e.Author,
		CreatedAt: e.CreatedAt,
		UpdatedAt: e.UpdatedAt,
	}
}

func (m *Book) ToEntity() *entity.Book {
	return &entity.Book{
		Id:        m.Id,
		Title:     m.Title,
		Author:    m.Author,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}
}
