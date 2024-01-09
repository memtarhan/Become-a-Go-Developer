package controllers

import (
	"BookStore/pkg/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"

	"BookStore/pkg/models"
)

var book models.Book

func GetBooks(writer http.ResponseWriter, request *http.Request) {
	books := models.GetBooks()
	response, _ := json.Marshal(books)
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	_, err := writer.Write(response)
	if err != nil {
		return
	}
}

func GetBookById(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	bookId := vars["bookId"]
	id, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("Error on GetBookById(:_) ", err)
		return
	}
	book, _ := models.GetBookById(id)
	response, _ := json.Marshal(book)
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	_, err = writer.Write(response)
	if err != nil {
		return
	}
}

func CreateBook(writer http.ResponseWriter, request *http.Request) {
	book := &models.Book{}
	utils.ParseBody(request, book)
	newBook := book.CreateBook()
	response, _ := json.Marshal(newBook)
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	_, err := writer.Write(response)
	if err != nil {
		return
	}
}

func DeleteBook(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	bookId := vars["bookId"]
	id, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("Error on DeleteBook(:_) ", err)
		return
	}
	book := models.DeleteBook(id)
	response, _ := json.Marshal(book)
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	_, err = writer.Write(response)
	if err != nil {
		return
	}
}

func SaveDummyBooks() {
	fmt.Println("Saving a dummy book")
	book := models.Book{
		ID:          -10,
		Name:        "Dummy for dummies",
		Author:      "Dumb dummy",
		Publication: "Dummy books",
	}

	book.CreateBook()

	fmt.Println("Saved a dummy book ", &book.Name)

}
