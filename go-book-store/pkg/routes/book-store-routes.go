package routes

import (
	"github.com/gorilla/mux"
	"github.com/richardHaggioGwati/go_projects/go-book-store/pkg/controllers"
)

var RegisterBookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/book", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book", controllers.GetBook).Methods("GET")
	router.HandleFunc("/book{bookId}", controllers.GetBookById).Methods("GET")
	router.HandleFunc("/book{bookId}", controllers.UpdateBook).Methods("PUT") // will consider switching to PATCH
	router.HandleFunc("/book{bookId}", controllers.DeleteBook).Methods("DELETE")
}
