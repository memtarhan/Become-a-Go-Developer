package models

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	d, err := gorm.Open(sqlite.Open("store.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db = d
	err = d.AutoMigrate(&Book{})
	if err != nil {
		fmt.Println("failed to connect database, ", err)
		panic("failed to connect database")
	}
}

func (book *Book) CreateBook() *Book {
	db.Create(&book)
	return book
}

func GetBooks() []Book {
	var books []Book
	db.Find(&books)
	return books
}

func GetBookById(id int64) (*Book, *gorm.DB) {
	var book Book
	db := db.Where("ID=?", id).Find(&book)
	return &book, db
}

func DeleteBook(id int64) Book {
	var book Book
	db.Where("ID=?", id).Delete(book)
	return book
}
