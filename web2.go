package main

import ("fmt"
"net/http")


func index_handler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, `<h1>hey baby</h1>
		<h1>hey baby</h1>
		<h1>hey baby</h1>`)
	fmt.Fprintf(w,"<h1>hey baby</h1>")
	fmt.Fprintf(w,"<p>you %s even add %s</p>","can","<strong>variables</strong>")

}

func main(){
	http.HandleFunc("/", index_handler)
	http.ListenAndServe(":8000", nil)
}