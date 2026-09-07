package repository

import (
	"errors"
	"math"

	"github.com/andruwizz/gin-rest-book-backend/internal/entity"
	"github.com/andruwizz/gin-rest-book-backend/internal/model"
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
	m := model.FromBookEntity(data)
	if err := b.db.Create(m).Error; err != nil {
		return err
	}
	*data = *m.ToEntity()
	return nil
}

func (b *bookRepository) List(limit int, page int) ([]entity.Book, *entity.Pagination, error) {
	var records []model.Book
	var pagination entity.Pagination
	var totalRecords int64

	query := b.db.Model(&model.Book{})

	pagination.Limit = limit
	pagination.Page = page
	query.Count(&totalRecords)
	pagination.TotalRecords = totalRecords
	pagination.TotalPage = int(math.Ceil(float64(totalRecords) / float64(pagination.GetPageLimit())))

	err := query.Clauses(clause.Limit{Offset: pagination.GetOffset(), Limit: &pagination.Limit}).
		Find(&records).
		Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil, nil
		}
		return nil, nil, err
	}

	books := make([]entity.Book, len(records))
	for i, r := range records {
		books[i] = *r.ToEntity()
	}
	pagination.Records = int64(pagination.Limit*(pagination.Page-1)) + int64(len(books))

	return books, &pagination, nil
}

func (b *bookRepository) Find(id string, data *entity.Book) error {
	m := model.FromBookEntity(data)
	if err := b.db.Where("id = ?", id).Take(m).Error; err != nil {
		return err
	}
	*data = *m.ToEntity()
	return nil
}

func (b *bookRepository) Update(data *entity.Book) error {
	m := model.FromBookEntity(data)
	if err := b.db.Save(m).Error; err != nil {
		return err
	}
	*data = *m.ToEntity()
	return nil
}

func (b *bookRepository) Delete(data *entity.Book) error {
	m := model.FromBookEntity(data)
	if err := b.db.Delete(m).Error; err != nil {
		return err
	}
	return nil
}
