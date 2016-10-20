package main

import (
	"fmt"
	"net/http"
)

type Hello struct{}

func (h Hello) ServeHTTP(
w http.ResponseWriter,
r *http.Request) {
	fmt.Fprintln(w, "Hello!")
	fmt.Fprintln(w, w)

}

func main() {
	var h Hello
	http.ListenAndServe("localhost:4000", h)
}