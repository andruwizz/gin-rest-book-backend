package repository

import (
	"github.com/andruwizz/gin-book-sharing-backend/internal/entity"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(data *entity.User) error
	Find(email string, data *entity.User) error
	Update(data *entity.User) error
	Delete(data *entity.User) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

func (u *userRepository) Create(data *entity.User) error {
	return u.db.Create(&data).Error
}

func (u *userRepository) Find(email string, data *entity.User) error {
	return u.db.Where("email = ?", email).Take(data).Error
}

func (u *userRepository) Update(data *entity.User) error {
	return u.db.Save(data).Error
}

func (u *userRepository) Delete(data *entity.User) error {
	return u.db.Delete(data).Error
}
