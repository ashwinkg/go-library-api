package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ashwinkg/go-library-api/pkg/routes"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	routes.RegisterLibraryApiRoutes(router)
	http.Handle("/", router)
	fmt.Println("Library Server is running...")
	log.Fatal(http.ListenAndServe(":8080", router))

}
