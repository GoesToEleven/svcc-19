package main

import (
	"fmt"
	"net/http"
)

type hotdog int

func (hd hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "here is my response from a value of type hotdog")
}

func main() {
	var x hotdog
	http.HandleFunc("/", foo)
	http.HandleFunc("/about", about)
	http.Handle("/ballgame", x)
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
