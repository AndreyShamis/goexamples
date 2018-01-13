package main

import (
	"flag"
	"os/exec"
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
)

// Book struct (Model)
type Book struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

// Author struct
type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

// Init books var as a slice Book struct
var books []Book

// Get all books
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

// Get single book
func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Gets params
	// Loop through books and find one with the id from the params
	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Book{})
}

// Add new book
func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(100000000)) // Mock ID - not safe
	books = append(books, book)
	json.NewEncoder(w).Encode(book)
}

// Update book
func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			var book Book
			_ = json.NewDecoder(r.Body).Decode(&book)
			book.ID = params["id"]
			books = append(books, book)
			json.NewEncoder(w).Encode(book)
			return
		}
	}
}

// Delete book
func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			break
		}
		json.NewEncoder(w).Encode(books)
	}
}

func printStart() (string) {
	log.Printf("Starting application.")
	return "Starting application."
}

func printEnd()(string){
	log.Printf("Finishing application.")
	return "Finishing application."
}
//
//func execute(cmd string, wg *sync.WaitGroup) {
//	fmt.Println("command is ",cmd)
//	// splitting head => g++ parts => rest of the command
//	parts := strings.Fields(cmd)
//	head := parts[0]
//	parts = parts[1:len(parts)]
//
//	out, err := exec.Command(head,parts...).Output()
//	if err != nil {
//		fmt.Printf("%s", err)
//	}
//	fmt.Printf("%s", out)
//	wg.Done() // Need to signal to waitgroup that this goroutine is done
//}
func getIfconfig() (string,  error){
	cmd := "ifconfig"
	out, err := exec.Command("sh","-c",cmd).Output()
	return string(out), err
}
// Main function
func main() {
	printStart()
	out, err := getIfconfig()
	log.Printf("out: %s , err : %s", out, err)
	// Init router
	bind := flag.String("l", "0.0.0.0:8000", "listen on ip:port")
	r := mux.NewRouter()

	// Hardcoded data - @todo: add database
	books = append(books, Book{ID: "1", Isbn: "978-5-389-04926-0", Title: "Crime and Punishment", Author: &Author{Firstname: "Fyodor", Lastname: "Dostoyevsky"}})
	books = append(books, Book{ID: "2", Isbn: "454555", Title: "Another book example", Author: &Author{Firstname: "Hello", Lastname: "World"}})

	log.Printf("Route handles & endpoints.")
	// Route handles & endpoints
	r.HandleFunc("/", getBooks).Methods("GET")
	r.HandleFunc("/{id}", getBook).Methods("GET")
	r.HandleFunc("/", createBook).Methods("POST")
	r.HandleFunc("/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/{id}", deleteBook).Methods("DELETE")
	log.Printf("Listening on http://%s.", *bind)
	// Start server
	log.Fatal(http.ListenAndServe(*bind, r))
	printEnd()
}

// Request sample
// {
// 	"isbn":"4545454",
// 	"title":"Book Three",
// 	"author":{"firstname":"Harry","lastname":"White"}
// }
