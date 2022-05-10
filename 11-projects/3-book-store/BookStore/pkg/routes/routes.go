package routes

import (
	"github.com/gorilla/mux"

	"BookStore/pkg/controllers"
)

var RegisterBookStoreRoutes = func(router *mux.Router) {
	//router.HandleFunc("/books", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/books", controllers.GetBooks).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.GetBookById).Methods("GET")
	//router.HandleFunc("/book/{bookId}", controllers.UpdateBook).Methods("PUT")
	//router.HandleFunc("/book/{bookId}", controllers.DeleteBook).Methods("DELETE")
}
