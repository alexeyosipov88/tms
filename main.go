package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// book struct

type Book struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

type Author struct {
	firstName string `json: "firsname"`
	lastName  string `json: "lastname`
}

var books []Book

func main() {
	books = append(books, Book{ID: "1", Isbn: "448743", Title: "Book One", Author: &Author{firstName: "Jon", lastName: "Smith"}})
	books = append(books, Book{ID: "2", Isbn: "448243", Title: "Book Two", Author: &Author{firstName: "David", lastName: "Bryan"}})

	r := mux.NewRouter()

	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books/	", getBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", getBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", getBook).Methods("DELETE")

	// r.HandleFunc("api/books/{id}", getBook).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", r))
	fmt.Println("")
}

// mock data

// Init books as slice

func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)

}

func getBook(w http.ResponseWriter, r *http.Request) {

}

// func getBooks(w http.ResponseWriter, *r http.Request) {

// }

// //
// func createBook(w http.ResponseWriter, *r http.Request) {

// }

// func updateBook() {

// }

// func deleteBook() {

// }
