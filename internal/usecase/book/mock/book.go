package mock

import (
	"github.com/andruwizz/gin-rest-book-backend/internal/entity"
	"github.com/andruwizz/gin-rest-book-backend/internal/usecase/book"
)

type BookUsecase interface {
	Create(req *book.BookCreateParam) (*entity.Book, error)
	List(req *book.BookListParam) ([]entity.Book, *entity.Pagination, error)
	Get(req *book.BookGetParam) (*entity.Book, error)
	Update(req *book.BookUpdateParam) (*entity.Book, error)
	Delete(req *book.BookDeleteParam) error
}

type mockBookUsecase struct{}

func NewMockBookUsecase() BookUsecase {
	return &mockBookUsecase{}
}

func (c *mockBookUsecase) Create(req *book.BookCreateParam) (*entity.Book, error) {
	book := &entity.Book{
		Id:        "348c6370-801d-4fab-80a3-5b3ecbc88760",
		Title:     req.Title,
		Author:    req.Author,
		CreatedAt: 1788410288537,
		UpdatedAt: 1788410288537,
	}

	return book, nil
}

func (c *mockBookUsecase) List(req *book.BookListParam) ([]entity.Book, *entity.Pagination, error) {
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

func (c *mockBookUsecase) Get(req *book.BookGetParam) (*entity.Book, error) {
	book := &entity.Book{
		Id:        "348c6370-801d-4fab-80a3-5b3ecbc88760",
		Title:     "Book Two",
		Author:    "Book Author",
		CreatedAt: 1788410288537,
		UpdatedAt: 1788410288537,
	}

	return book, nil
}

func (c *mockBookUsecase) Update(req *book.BookUpdateParam) (*entity.Book, error) {
	book := &entity.Book{
		Id:        "348c6370-801d-4fab-80a3-5b3ecbc88760",
		Title:     req.Title,
		Author:    req.Author,
		CreatedAt: 1788410288537,
		UpdatedAt: 1788410288537,
	}

	return book, nil
}

func (c *mockBookUsecase) Delete(req *book.BookDeleteParam) error {
	return nil
}
