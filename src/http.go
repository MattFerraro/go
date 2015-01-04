package main

import (
	_ "bufio"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/twinj/uuid"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"text/template"
)

type Move struct {
	// A Move is just a coordinate but if it has the special value (-1, -1)
	// it represents a skipped move
	X int
	Y int
}
type Game struct {
	Uuid    string
	Player1 string
	Player2 string
	Size    int
	Moves   []Move
}

/*  Custom marshal and unmarshal methods are needed because without them, the
	list of moves looks like this:
		[{"X":8,"Y":4}, ...]
	With the custom marshalling, the list of moves looks like this:
		[[8,4], ...]
	Which uses 5 bytes per move instead of 13
*/
func (m Move) MarshalJSON() ([]byte, error) {
	return json.Marshal([]int{m.X, m.Y})
}
func (m *Move) UnmarshalJSON(data []byte) error {
	var coordinate [2]int
	json.Unmarshal(data, &coordinate)
	m.X = coordinate[0]
	m.Y = coordinate[1]
	return nil
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

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

func gameData(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	gameId := vars["gameId"]
	fmt.Println(gameId)

	dat, err := ioutil.ReadFile("./" + gameId + ".json")
	check(err)

	// To ensure that all strings served from this method are in fact valid
	// json, it is necessary to try unmarshalling the data before returning it
	var game Game
	check(json.Unmarshal(dat, &game))

	io.WriteString(w, string(dat))
}

func newGame(w http.ResponseWriter, r *http.Request) {
	// The UUID for this game which will be used everywhere
	// as this game's identifier
	u := uuid.NewV4()
	io.WriteString(w, u.String())

	var g Game
	g.Uuid = u.String()
	g.Player1 = "foo"
	g.Player2 = "bar"
	g.Moves = append(g.Moves, Move{0, 0})
	g.Moves = append(g.Moves, Move{1, 1})
	g.Size = 19

	b, _ := json.Marshal(g)

	// Create a file with this name and initialize it
	f, err := os.Create("./" + u.String() + ".json")
	check(err)
	defer f.Close()
	f.Write(b)
	f.WriteString("\n")
}

func main() {
	// The default UUID format has gross curly brackets in it!
	uuid.SwitchFormat(uuid.CleanHyphen)

	// Setup all routes and serve
	r := mux.NewRouter()
	r.HandleFunc("/", hello)
	r.HandleFunc("/newgame", newGame)
	r.HandleFunc("/game/{gameId}", game)
	r.HandleFunc("/gamedata/{gameId}", gameData)
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

	game and gameData http endpoints should be merged and the returned type
	(either html or json) should reflect the mimetype requested
*/
