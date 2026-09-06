package repository

import (
	"github.com/andruwizz/gin-rest-book-backend/internal/entity"
	"github.com/andruwizz/gin-rest-book-backend/internal/model"
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
	m := model.FromUserEntity(data)
	return u.db.Create(m).Error
}

func (u *userRepository) Find(email string, data *entity.User) error {
	m := model.FromUserEntity(data)
	return u.db.Where("email = ?", email).Take(m).Error
}

func (u *userRepository) Update(data *entity.User) error {
	m := model.FromUserEntity(data)
	return u.db.Save(m).Error
}

func (u *userRepository) Delete(data *entity.User) error {
	m := model.FromUserEntity(data)
	return u.db.Delete(m).Error
}
