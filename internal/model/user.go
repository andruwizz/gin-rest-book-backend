package model

import (
	"time"

	"github.com/andruwizz/gin-rest-book-backend/internal/entity"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	Id        string `gorm:"type:uuid;primaryKey"`
	Name      string `gorm:"type:varchar"`
	Email     string `gorm:"type:varchar"`
	Password  string `gorm:"type:varchar"`
	CreatedAt int64  `gorm:"type:bigint"`
	UpdatedAt int64  `gorm:"type:bigint"`
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	newUUID := uuid.New()
	currentTime := time.Now().UnixMilli()
	tx.Statement.SetColumn("id", newUUID)
	tx.Statement.SetColumn("CreatedAt", currentTime)
	tx.Statement.SetColumn("UpdatedAt", currentTime)

	return nil
}

func (u *User) BeforeUpdate(tx *gorm.DB) error {
	currentTime := time.Now().UnixMilli()
	tx.Statement.SetColumn("UpdatedAt", currentTime)

	return nil
}

func FromUserEntity(e *entity.User) *User {
	return &User{
		Id:        e.Id,
		Name:      e.Name,
		Email:     e.Email,
		Password:  e.Password,
		CreatedAt: e.CreatedAt,
		UpdatedAt: e.UpdatedAt,
	}
}

func (m *User) ToEntity() *entity.User {
	return &entity.User{
		Id:        m.Id,
		Name:      m.Name,
		Email:     m.Email,
		Password:  m.Password,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}
}
