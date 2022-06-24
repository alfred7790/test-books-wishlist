package service

import (
	"test-books-wishlist/internal/client"
	"test-books-wishlist/internal/repository"
)

type Service struct {
	Repo           repository.Repository
	GoogleBooksAPI *client.Handle
}

func NewService() *Service {
	return &Service{}
}
