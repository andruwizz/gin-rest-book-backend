package model

import (
	"math/rand"
	"time"

	"github.com/oklog/ulid"
	"gorm.io/gorm"
)

type Book struct {
	Id        string `gorm:"type:varchar;primaryKey" json:"id"`
	Title     string `gorm:"type:varchar" json:"title"`
	Author    string `gorm:"type:varchar" json:"author"`
	CreatedAt string `gorm:"type:bigint" json:"created_at"`
	UpdatedAt string `gorm:"type:bigint" json:"updated_at"`
}

func (b *Book) BeforeCreate(tx *gorm.DB) error {
	t := time.Unix(1000000, 0)
	entropy := ulid.Monotonic(rand.New(rand.NewSource(t.UnixNano())), 0)
	tx.Statement.SetColumn("id", ulid.MustNew(ulid.Timestamp(t), entropy))
	tx.Statement.SetColumn("CreatedAt", time.Now().Unix())
	tx.Statement.SetColumn("UpdatedAt", time.Now().Unix())

	return nil
}

func (b *Book) BeforeUpdate(tx *gorm.DB) error {
	tx.Statement.SetColumn("UpdatedAt", time.Now().Unix())

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
