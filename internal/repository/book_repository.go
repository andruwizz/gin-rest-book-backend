package repository

import (
	"errors"
	"math"

	"github.com/andruwizz/gin-collective-library-backend/internal/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type BookRepository interface {
	Create(data *model.Book) (*model.Book, error)
	Find(id string) (*model.Book, error)
	List(limit int, page int) (*[]model.Book, *model.Pagination, error)
	Update(id string, data model.Book) (*model.Book, error)
	Delete(id string) error
}

type bookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) BookRepository {
	return &bookRepository{db}
}

func (b *bookRepository) Create(data *model.Book) (*model.Book, error) {
	err := b.db.Create(&data).Error
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (b *bookRepository) Find(id string) (*model.Book, error) {
	var book model.Book

	err := b.db.
		Model(&model.Book{}).
		Where("id = ?", id).
		First(&book).
		Error

	if err != nil {
		return nil, err
	}

	return &book, nil
}

func (b *bookRepository) List(limit int, page int) (*[]model.Book, *model.Pagination, error) {
	var books []model.Book
	var pagination model.Pagination

	query := b.db

	pagination.Limit = limit
	pagination.Page = page

	err := query.Scopes(b.Paginate(books, &pagination, query)).
		Find(&books).
		Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil, nil
		}
		return nil, nil, err
	}

	return &books, &pagination, nil
}

func (b *bookRepository) Update(id string, data model.Book) (*model.Book, error) {
	var book model.Book

	curr, err := b.Find(id)
	if err != nil {
		return nil, err
	}

	err = b.db.
		Model(&book).
		Where("id = ?", id).
		Clauses(clause.Returning{}).
		Updates(data).
		Error

	if err != nil {
		return nil, err
	}
	book.Id = curr.Id
	book.CreatedAt = curr.CreatedAt

	return &book, nil
}

func (b *bookRepository) Delete(id string) error {
	_, err := b.Find(id)
	if err != nil {
		return nil
	}

	err = b.db.
		Where("id = ?", id).
		Delete(&model.Book{}).
		Error

	if err != nil {
		return err
	}

	return nil
}

func (b *bookRepository) Paginate(value []model.Book, pagination *model.Pagination, db *gorm.DB) func(db *gorm.DB) *gorm.DB {
	var totalRecords int64
	var currRecord int64

	currRecord = int64(len(value))

	b.db.Model(value).Count(&totalRecords)
	pagination.TotalRecords = totalRecords
	pagination.TotalPage = int(math.Ceil(float64(totalRecords) / float64(pagination.GetPageLimit())))
	pagination.Records = int64(pagination.Limit*(pagination.Page-1)) + int64(currRecord)

	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(pagination.GetOffset()).Limit(pagination.Limit)
	}
}
