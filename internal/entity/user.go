package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	Id        string `gorm:"type:uuid;primaryKey" json:"id"`
	Name      string `gorm:"type:varchar" json:"name"`
	Email     string `gorm:"type:varchar" json:"email"`
	Password  string `gorm:"type:varchar" json:"-"`
	CreatedAt string `gorm:"type:bigint" json:"created_at"`
	UpdatedAt string `gorm:"type:bigint" json:"updated_at"`
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	uuid := uuid.New()
	time := time.Now().UnixMilli()
	tx.Statement.SetColumn("id", uuid)
	tx.Statement.SetColumn("CreatedAt", time)
	tx.Statement.SetColumn("UpdatedAt", time)

	return nil
}

func (u *User) BeforeUpdate(tx *gorm.DB) error {
	time := time.Now().UnixMilli()
	tx.Statement.SetColumn("UpdatedAt", time)

	return nil
}
