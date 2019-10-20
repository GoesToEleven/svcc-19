package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", foo)
	http.ListenAndServe(":3000", nil)
}
func foo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "here's the ip %s", r.RemoteAddr)
}
