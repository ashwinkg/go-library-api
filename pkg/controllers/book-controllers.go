package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/ashwinkg/go-library-api/pkg/models"
	"github.com/ashwinkg/go-library-api/pkg/utils"
	"github.com/gorilla/mux"
)

func Getallbooks(w http.ResponseWriter, r *http.Request) {
	books, err := models.Getallbooks()
	if err != nil {
		log.Fatal("Error fetaching all the books")
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response, _ := json.Marshal(books)
	w.Write(response)

}

func GetbookbyId(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.ParseInt(params["id"], 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	book, _ := models.GetbookbyId(id)
	if book.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	response, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)

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
	params := mux.Vars(r)
	id, err := strconv.ParseInt(params["id"], 0, 0)
	if err != nil {
		http.Error(w, "Invalid book id", http.StatusBadRequest)
		return
	}

	updatedbook := &models.Book{}
	utils.ParseBody(r, updatedbook)

	book, err := models.UpdateBook(id, *updatedbook)
	if err != nil {
		http.Error(w, "Error updating book", http.StatusInternalServerError)
		return
	}
	response, err := json.Marshal(book)
	if err != nil {
		http.Error(w, "Error marshalling response", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Updated the below book"))
	w.Write(response)

}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.ParseInt(params["id"], 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}
	book, err := models.DeleteBook(id)
	if err != nil {
		http.Error(w, "Error deleting book", http.StatusInternalServerError)
		return
	}
	response, err := json.Marshal(book)
	if err != nil {
		http.Error(w, "Error marshalling response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Deleted the below book"))
	w.Write(response)
}
