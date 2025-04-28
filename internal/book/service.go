package book

import "BookStoreRest/pkg/models"

type Service interface {
	GetBooks() []models.Book
	CreateBook(book models.Book) models.Book
	DeleteBook(id int) bool
}

type service struct {
	repo Repository
}

func NewService(r Repository) Service {
	return &service{repo: r}
}

func (s *service) GetBooks() []models.Book {
	return s.repo.GetAll()
}

func (s *service) CreateBook(book models.Book) models.Book {
	return s.repo.Create(book)
}

func (s *service) DeleteBook(id int) bool {
	return s.repo.Delete(id)
}
