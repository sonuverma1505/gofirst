package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func index_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "who, are you")
}

func about_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "i am God")
}
func menu_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "No menu go back to home page")
}

func main() {
	r := mux.NewRouter()
	http.HandleFunc("/", index_handler)
	http.HandleFunc("/about", about_handler)
	http.HandleFunc("about/menu", menu_handler)
	http.ListenAndServe(":8080", nil)
}
