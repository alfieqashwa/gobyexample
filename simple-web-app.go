package main

import (
	"fmt"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Whoa, Go is Cello's toy!")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This site is created by Alfie Cello")
}
func contactHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `Call me right away.
					My Numb is +62812809319xxxx`)
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/about", aboutHandler)
	http.HandleFunc("/contactme", contactHandler)
	http.ListenAndServe(":8000", nil)
}
