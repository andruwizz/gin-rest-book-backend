package repository

import (
	"errors"
	"math"

	"github.com/andruwizz/gin-rest-book-backend/internal/entity"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type BookRepository interface {
	Create(data *entity.Book) error
	List(limit int, page int) ([]entity.Book, *entity.Pagination, error)
	Find(id string, data *entity.Book) error
	Update(data *entity.Book) error
	Delete(data *entity.Book) error
}

type bookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) BookRepository {
	return &bookRepository{db}
}

func (b *bookRepository) Create(data *entity.Book) error {
	return b.db.Create(&data).Error
}

func (b *bookRepository) List(limit int, page int) ([]entity.Book, *entity.Pagination, error) {
	var books []entity.Book
	var pagination entity.Pagination
	var totalRecords int64

	query := b.db

	pagination.Limit = limit
	pagination.Page = page

	query.Model(books).Count(&totalRecords)

	pagination.TotalRecords = totalRecords
	pagination.TotalPage = int(math.Ceil(float64(totalRecords) / float64(pagination.GetPageLimit())))
	pagination.Records = int64(pagination.Limit*(pagination.Page-1)) + int64(len(books))

	err := query.Clauses(clause.Limit{Offset: pagination.GetOffset(), Limit: &pagination.Limit}).
		Find(&books).
		Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil, nil
		}
		return nil, nil, err
	}

	return books, &pagination, nil
}

func (b *bookRepository) Find(id string, data *entity.Book) error {
	return b.db.Where("id = ?", id).Take(data).Error
}

func (b *bookRepository) Update(data *entity.Book) error {
	return b.db.Save(data).Error
}

func (b *bookRepository) Delete(data *entity.Book) error {
	return b.db.Delete(data).Error
}
