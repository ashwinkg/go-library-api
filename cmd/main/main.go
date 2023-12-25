package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	fmt.Println("Library Server is running...")
	log.Fatal(http.ListenAndServe(":8080", router))

}
