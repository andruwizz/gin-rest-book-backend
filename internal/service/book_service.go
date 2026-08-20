package service

import "github.com/andruwizz/gin-collective-library-backend/internal/repository"

type BookService struct {
	repository *repository.BookRepository
}

func NewBookService(r *repository.BookRepository) *BookService {
	return &BookService{repository: r}
}
