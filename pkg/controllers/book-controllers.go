package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/ashwinkg/go-library-api/pkg/models"
	"github.com/ashwinkg/go-library-api/pkg/utils"
)

func Getallbooks(w http.ResponseWriter, r *http.Request) {

}

func GetbookbyId(w http.ResponseWriter, r *http.Request) {

}

func AddBook(w http.ResponseWriter, r *http.Request) {
	Book := &models.Book{}
	utils.ParseBody(r, Book)
	book := Book.AddBook()
	response, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)

}

func UpdateBook(w http.ResponseWriter, r *http.Request) {

}

func DeleteBook(w http.ResponseWriter, r *http.Request) {

}
