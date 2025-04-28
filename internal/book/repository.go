package book

import (
	"BookStoreRest/pkg/models"
	"sync"
)

type Repository interface {
	GetAll() []models.Book
	Create(book models.Book) models.Book
	Delete(id int) bool
}

type memoryRepo struct {
	mu     sync.Mutex
	books  []models.Book
	nextID int
}

func NewMemoryRepository() Repository {
	return &memoryRepo{
		books: []models.Book{
			{ID: 1, Title: "Head First Java: A Brain-Friendly Guide", Author: "Kathy Sierra", Price: 10.99},
			{ID: 2, Title: "Learning Go: An Idiomatic Approach to Real-World Go Programming", Author: "Jon Bodner", Price: 12.49},
			{ID: 3, Title: "Python Crash Course, 3rd Edition: A Hands-On, Project-Based Introduction to Programming", Author: "Eric Matthes", Price: 15.49},
		},
		nextID: 4,
	}
}

func (r *memoryRepo) GetAll() []models.Book {
	r.mu.Lock()
	defer r.mu.Unlock()
	return r.books
}

func (r *memoryRepo) Create(book models.Book) models.Book {
	r.mu.Lock()
	defer r.mu.Unlock()
	book.ID = r.nextID
	r.nextID++
	r.books = append(r.books, book)
	return book
}

func (r *memoryRepo) Delete(id int) bool {
	r.mu.Lock()
	defer r.mu.Unlock()
	for i, book := range r.books {
		if book.ID == id {
			r.books = append(r.books[:i], r.books[i+1:]...)
			return true
		}
	}
	return false
}
