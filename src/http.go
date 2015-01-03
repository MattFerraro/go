package main

import (
	"fmt"
	"net/http"
	"text/template"
	"github.com/gorilla/mux"
)

func hello(w http.ResponseWriter, r *http.Request) {
	template_path := "index.html"
	template, _ := template.ParseFiles(template_path)
	template.Execute(w, nil)
	fmt.Println("Hit index.html")
}

func main() {
    r := mux.NewRouter()
    r.HandleFunc("/", hello)
    r.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))
    http.Handle("/", r)
    http.ListenAndServe(":8000", nil)
}
