package main

import (
	"BookStoreRest/internal/book"
	"BookStoreRest/internal/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	repo := book.NewMemoryRepository()
	service := book.NewService(repo)
	handler := book.NewHandler(service)

	r := router.NewRouter(handler)

	fmt.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
