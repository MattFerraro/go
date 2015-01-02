package main

import (
	"io"
	"net/http"
	"github.com/gorilla/mux"
)

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello world and stuff!")
}

func main() {
	http.HandleFunc("/", hello)
	http.ListenAndServe(":8000", nil)
}
