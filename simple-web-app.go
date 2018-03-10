package main

import (
	"fmt"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `<h1>Whoa, Go is Cello's toy!</h1>
<p>Go is fast!</p>
<p>...and simple!</p>`)
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This site is created by Alfie Cello")
}
func contactHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `<p>Call me!
My Numb is <strong>+62812809319xxxx</strong></p>`)
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/about", aboutHandler)
	http.HandleFunc("/contactme", contactHandler)
	http.ListenAndServe(":8000", nil)
}
