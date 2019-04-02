package main

import (
	"encoding/json"
	// 	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Book struct {
	ID     int    `json:id`
	Title  string `json:title`
	Author string `json:author`
	Year   string `json:year`
}

var books []Book

func main() {
	router := mux.NewRouter()

	books = append(books,
		Book{ID: 1, Title: "Golang Pointers", Author: "Mr. Golang", Year: "2019"},
		Book{ID: 2, Title: "Restful API With Golang", Author: "Mr. RestAPI", Year: "2014"},
		Book{ID: 3, Title: "Golang Mux", Author: "Mr. Gorilla", Year: "2017"},
		Book{ID: 4, Title: "Golang Routines", Author: "Mr. Routine", Year: "2017"},
		Book{ID: 5, Title: "Golang Routers", Author: "Mr. Routers", Year: "2018"})

	router.HandleFunc("/books", getBooks).Methods("GET")
	router.HandleFunc("/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/books", addBook).Methods("POST")
	router.HandleFunc("/books", updateBook).Methods("PUT")
	router.HandleFunc("/books{id}", removeBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
}

func getBooks(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(books)
}

func getBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	i, _ := strconv.Atoi(params["id"])

	for _, book := range books {
		if book.ID == i {
			json.NewEncoder(w).Encode(&book)
		}
	}
}

func addBook(w http.ResponseWriter, r *http.Request) {
	log.Println("Add one book")
}

func updateBook(w http.ResponseWriter, r *http.Request) {
	log.Println("Update one book")
}

func removeBook(w http.ResponseWriter, r *http.Request) {
	log.Println("Remove one book")
}
