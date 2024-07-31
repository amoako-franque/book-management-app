package routes

import (
	"github.com/amoako-franque/go-api-book-management-app/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterBookRoutes = func(router *mux.Router) {
	router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/books/", controllers.GetBooks).Methods("GET")
	router.HandleFunc("/books/{bookId}", controllers.GetBookById).Methods("GET")
	router.HandleFunc("/books/{bookId}", controllers.UpdateBookById).Methods("PUT")
	router.HandleFunc("/books/{bookId}", controllers.DeleteBookById).Methods("DELETE")
}
