package routes

import (
	"github.com/ashwinkg/go-library-api/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterLibraryApiRoutes = func(router *mux.Router) {
	router.HandleFunc("/books/", controllers.Getallbooks).Methods("GET")
	router.HandleFunc("/books/{id}", controllers.GetbookbyId).Methods("GET")
	router.HandleFunc("/books/", controllers.AddBook).Methods("POST")
	router.HandleFunc("/books/{id}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/books/{id}", controllers.DeleteBook).Methods("DELETE")
}
