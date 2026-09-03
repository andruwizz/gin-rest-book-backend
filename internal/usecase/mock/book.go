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
	res := []entity.Book{
		{
			Id:        "305c8059-28d7-49f7-a15c-acba512f2b0a",
			Title:     "Book One",
			Author:    "Book Author",
			CreatedAt: 1788410288537,
			UpdatedAt: 1788410288537,
		}, {
			Id:        "348c6370-801d-4fab-80a3-5b3ecbc88760",
			Title:     "Book Two",
			Author:    "Book Author",
			CreatedAt: 1788410288537,
			UpdatedAt: 1788410288537,
		}, {
			Id:        "5a5a9021-7b19-47f4-bda8-05d2551f8ed8",
			Title:     "Book Three",
			Author:    "Book Author",
			CreatedAt: 1788410288537,
			UpdatedAt: 1788410288537,
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
	book := &entity.Book{
		Id:        "348c6370-801d-4fab-80a3-5b3ecbc88760",
		Title:     "Book Two",
		Author:    "Book Author",
		CreatedAt: 1788410288537,
		UpdatedAt: 1788410288537,
	}

	return book, nil
}

func (c *mockBookUsecase) Update(req *request.BookUpdate) (*entity.Book, error) {
	timeNow := time.Now().UnixMilli()
	book := &entity.Book{
		Id:        "348c6370-801d-4fab-80a3-5b3ecbc88760",
		Title:     req.Title,
		Author:    req.Author,
		CreatedAt: 1788410288537,
		UpdatedAt: timeNow,
	}

	return book, nil
}

func (c *mockBookUsecase) Delete(req *request.BookDelete) error {
	return nil
}
