package main

import (
	"fmt"
	"net/http"
	"text/template"
	_ "io/ioutil"
	"io"
	"github.com/gorilla/mux"
	"github.com/twinj/uuid"
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

func newGame(w http.ResponseWriter, r *http.Request) {
	// The UUID for this game which will be used everywhere
	// as this game's identifier
	u := uuid.NewV4()
	io.WriteString(w, u.String())
}

func main() {
	// The default UUID format has gross curly brackets in it!
	uuid.SwitchFormat(uuid.CleanHyphen)

	// Setup all routes and serve
    r := mux.NewRouter()
    r.HandleFunc("/", hello)
    r.HandleFunc("/newgame", newGame)
    r.HandleFunc("/game/{gameId}", game)
    r.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))
    http.Handle("/", r)
    http.ListenAndServe(":8000", nil)
}


// Misc notes:
/*
	Games should probably be saved as individual json files, <UUID>.json
	Those files should have a history of all moves that were made in this game

	Maybe there should also be a copy of the board as it was at the very end?
	This SHOULD be figure-out-able given just the history of moves, but
	people making very primitive AIs shouldn't have to roll their own board
	generator.
*/
