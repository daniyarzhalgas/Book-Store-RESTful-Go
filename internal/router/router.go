package router

import (
	"BookStoreRest/internal/book"
	"net/http"
)

func NewRouter(handler *book.Handler) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/books", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handler.GetBooks(w, r)
		case http.MethodPost:
			handler.CreateBook(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/books/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodDelete {
			handler.DeleteBook(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	return mux
}
