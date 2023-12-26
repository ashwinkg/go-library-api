package models

import (
	"log"

	"github.com/ashwinkg/go-library-api/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `json:"name"`
	Author      string `json:"author"`
	Isbn        string `json:"isbn"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()

	//Peform database migration
	err := db.AutoMigrate(&Book{})
	if err != nil {
		log.Fatal(err)
	}
}

func (book *Book) AddBook() *Book {
	db.NewRecord(book)
	db.Create(&book)
	return book
}

func GetbookbyId(bookID int) (*Book, *gorm.DB) {
	var book Book
	db.Where("ID=?", bookID).Find(&book)
	return &book, db
}
