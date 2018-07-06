package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func bookHandle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	title := vars["title"]
	page := vars["page"]

	fmt.Fprintf(w, "You've requested the book: %s on page %s\n", title, page)
}

func indexHandle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

func startServer(port string) {
	r := mux.NewRouter()
	r.HandleFunc("/books/{title}/page/{page}", bookHandle)
	r.HandleFunc("/", indexHandle)

	fmt.Println(fmt.Sprintf("Starting web server on port %s", port))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), r))
}

func main() {
	startServer("8080")
}
