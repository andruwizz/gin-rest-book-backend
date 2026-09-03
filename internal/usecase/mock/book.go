package mock

import (
	"time"

	"github.com/andruwizz/gin-rest-book-backend/internal/delivery/request"
	"github.com/andruwizz/gin-rest-book-backend/internal/entity"
	"github.com/andruwizz/gin-rest-book-backend/internal/repository"
	"github.com/andruwizz/gin-rest-book-backend/internal/repository/mock"
	"github.com/google/uuid"
)

type BookUsecase interface {
	Create(req *request.BookCreate) (*entity.Book, error)
	List(req *request.BookList) ([]entity.Book, *entity.Pagination, error)
	Get(req *request.BookGet) (*entity.Book, error)
	Update(req *request.BookUpdate) (*entity.Book, error)
	Delete(req *request.BookDelete) error
}

type mockBookUsecase struct {
	repository mock.BookRepository
}

func NewMockBookUsecase(repository repository.BookRepository) BookUsecase {
	return &mockBookUsecase{repository}
}

func (c *mockBookUsecase) Create(req *request.BookCreate) (*entity.Book, error) {
	time := time.Now().UnixMilli()
	book := &entity.Book{
		Id:        uuid.NewString(),
		Title:     req.Title,
		Author:    req.Author,
		CreatedAt: time,
		UpdatedAt: time,
	}

	return book, nil
}

func (c *mockBookUsecase) List(req *request.BookList) ([]entity.Book, *entity.Pagination, error) {
	time := time.Now().UnixMilli()
	res := []entity.Book{
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

	pagination := &entity.Pagination{
		Page:         1,
		TotalRecords: 100,
		Records:      10,
		Limit:        10,
		TotalPage:    10,
	}

	return res, pagination, nil
}

func (c *mockBookUsecase) Get(req *request.BookGet) (*entity.Book, error) {
	time := time.Now().UnixMilli()
	book := &entity.Book{
		Id:        uuid.NewString(),
		Title:     "Book Title",
		Author:    "Book AUthor",
		CreatedAt: time,
		UpdatedAt: time,
	}

	return book, nil
}

func (c *mockBookUsecase) Update(req *request.BookUpdate) (*entity.Book, error) {
	timePass := time.Now().Add(time.Duration(-10) * time.Hour).UnixMilli()
	timeNow := time.Now().UnixMilli()
	book := &entity.Book{
		Id:        uuid.NewString(),
		Title:     req.Title,
		Author:    req.Author,
		CreatedAt: timePass,
		UpdatedAt: timeNow,
	}

	return book, nil
}

func (c *mockBookUsecase) Delete(req *request.BookDelete) error {
	return nil
}
