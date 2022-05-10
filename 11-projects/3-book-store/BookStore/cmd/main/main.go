package main

import (
	"BookStore/pkg/routes"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	routes.RegisterBookStoreRoutes(router)

	//controllers.SaveDummyBooks()

	http.Handle("/", router)

	fmt.Println("Starting server at localhost:8000")
	log.Fatal(http.ListenAndServe("localhost:8000", router))
}
