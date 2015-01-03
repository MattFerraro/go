package main

import (
	"fmt"
	"net/http"
	"text/template"
	"github.com/gorilla/mux"
)

func hello(w http.ResponseWriter, r *http.Request) {
	templatePath := "index.html"
	template, _ := template.ParseFiles(templatePath)
	template.Execute(w, nil)
	fmt.Println("Hit index.html")
}

func game(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	gameId := vars["gameId"]
	fmt.Println(gameId)

	gameMap := make(map[string]string)
	gameMap["GameId"] = gameId

	templatePath := "game.html"
	template, _ := template.ParseFiles(templatePath)
	template.Execute(w, gameMap)
	fmt.Println("Hit game.html")
}

func main() {
    r := mux.NewRouter()
    r.HandleFunc("/", hello)
    r.HandleFunc("/game/{gameId}", game)
    r.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))
    http.Handle("/", r)
    http.ListenAndServe(":8000", nil)
}
