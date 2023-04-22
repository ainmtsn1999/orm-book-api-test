package service

import "github.com/ainmtsn1999/orm-book-api-test/repository"

type Service struct {
	repo repository.RepoInterface
}

type ServiceInterface interface {
	BookService
}

func NewService(repo repository.RepoInterface) *Service {
	return &Service{repo: repo}
}
