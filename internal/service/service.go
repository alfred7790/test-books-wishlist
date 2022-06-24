package service

import "test-books-wishlist/internal/repository"

type Service struct {
	Repo repository.Repository
}

func NewService() *Service {
	return &Service{}
}
