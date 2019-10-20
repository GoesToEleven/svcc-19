package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/about", about)
	http.ListenAndServe(":8080", nil)
}

// interface in go says,
// "Hey baby, if you've got these methods, then you're my type."

func foo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello world")
}

func about(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "we are at about")
}
