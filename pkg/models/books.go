package models

import (
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
	db.AutoMigrate(&Book{})
}

func (book *Book) AddBook() *Book {
	db.NewRecord(book)
	db.Create(&book)
	return book
}

func GetbookbyId(bookID int64) (*Book, *gorm.DB) {
	var book Book
	db.Where("ID=?", bookID).Find(&book)
	return &book, db
}

func Getallbooks() ([]Book, error) {
	var books []Book
	// Get all records
	result := db.Find(&books)
	if result.Error != nil {
		// Handle the error, return it, or log it
		return nil, result.Error
	}

	return books, nil
}

func DeleteBook(bookID int64) (*Book, error) {
	var book Book

	//First, find the book
	result := db.First(&book, bookID)
	if result.Error != nil {
		return nil, result.Error
	}

	//Then delete it
	result = db.Where("ID=?", bookID).Delete(&book)
	if result.Error != nil {
		return nil, result.Error
	}
	return &book, nil
}

func UpdateBook(bookID int64, updatedbook Book) (*Book, error) {

	var existingBook Book
	result := db.First(&existingBook, bookID)
	if result.Error != nil {
		return nil, result.Error // Handle errors, e.g., record not found
	}

	result = db.Model(&existingBook).Updates(updatedbook)
	if result.Error != nil {
		return nil, result.Error
	}
	return &existingBook, nil
}
