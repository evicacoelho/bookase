package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

type Book struct {
	ID     int
	title  string
	author string
	owner  string
}

var books []Book

// function to list all books
func GetBooks(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(books)
}

// function to add new book to the list of books
func AddBooks(w http.ResponseWriter, r *http.Request) {
	var newBook Book
	newBook.ID = len(books)
	// opening a new reader to get a new book
	reader := bufio.NewReader((os.Stdin))
	//user inputs into a new book
	fmt.Println(" please, enter the book name: ")
	newBook.title, _ = reader.ReadString('\n')
	fmt.Println(" please, enter the author of the book: ")
	newBook.author, _ = reader.ReadString('\n')
	fmt.Println(" please enter your name for register: ")
	newBook.owner, _ = reader.ReadString('\n')

	// appending the new book
	books = append(books, newBook)

	// response
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode()
}

func main() {
	// opening the router
	router := mux.NewRouter()

	// starting the bookcase
	books = append(books, Book{ID: 1, title: "the library is open", author: "Ru Paul Jones", owner: "the owner"})

	// calling endpoints
	router.HandleFunc("/books", GetBooks).Methods("GET")
	router.HandleFunc("/books", AddBooks).Methods("POST")

	// err handling
	log.Fatal(http.ListenAndServe(":8000", router))

}
