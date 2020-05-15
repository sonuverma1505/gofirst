package main

import("fmt" 
	"net/http")

func index_handler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "who, are you")
}

func about_handler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w,"i am sonu")
}

func main() {
	http.HandleFunc("/", index_handler)
	http.HandleFunc("/about", about_handler)
	http.ListenAndServe(":8080", nil)


}