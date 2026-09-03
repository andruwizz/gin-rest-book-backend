package mock

import (
	"errors"
	"math"
	"time"

	"github.com/andruwizz/gin-rest-book-backend/internal/entity"
	"github.com/google/uuid"
)

type BookRepository interface {
	Create(data *entity.Book) error
	List(limit int, page int) ([]entity.Book, *entity.Pagination, error)
	Find(id string, data *entity.Book) error
	Update(data *entity.Book) error
	Delete(data *entity.Book) error
}

type mockBookRepository struct {
	book map[string]*entity.Book
}

func NewMockBookRepository() BookRepository {
	return &mockBookRepository{
		book: make(map[string]*entity.Book),
	}
}

func (b *mockBookRepository) Create(data *entity.Book) error {
	b.book[data.Id] = data

	return nil
}

func (b *mockBookRepository) List(limit int, page int) ([]entity.Book, *entity.Pagination, error) {
	var books []entity.Book
	var pagination entity.Pagination
	var totalRecords int64

	time := time.Now().UnixMilli()
	books = []entity.Book{
		{
			Id:        uuid.NewString(),
			Title:     "Book One",
			Author:    "Book AUthor",
			CreatedAt: time,
			UpdatedAt: time,
		}, {
			Id:        uuid.NewString(),
			Title:     "Book Two",
			Author:    "Book AUthor",
			CreatedAt: time,
			UpdatedAt: time,
		}, {
			Id:        uuid.NewString(),
			Title:     "Book Three",
			Author:    "Book AUthor",
			CreatedAt: time,
			UpdatedAt: time,
		},
	}

	pagination.Limit = limit
	pagination.Page = page
	totalRecords = 10

	pagination.TotalRecords = totalRecords
	pagination.TotalPage = int(math.Ceil(float64(totalRecords) / float64(pagination.GetPageLimit())))
	pagination.Records = int64(pagination.Limit*(pagination.Page-1)) + int64(len(books))

	return books, &pagination, nil
}

func (b *mockBookRepository) Find(id string, data *entity.Book) error {
	data = b.book[id]
	if data != nil {
		return nil
	}

	return errors.New("record not found")
}

func (b *mockBookRepository) Update(data *entity.Book) error {
	b.book[data.Id] = data
	return nil
}

func (b *mockBookRepository) Delete(data *entity.Book) error {
	b.book[data.Id] = nil
	return nil
}
