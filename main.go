package main

import (
	"bufio"
	"math/rand"
	"net/http"
	"os"
	"text/template"
	"time"
)

var state State

type Letter struct {
	Value string
	Used  bool
}

type State struct {
	Level        string
	CompleteWord string
	Letters      []Letter
	Errors       int
	CurrentWord  []string
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	page := template.Must(template.ParseFiles("views/index.html"))

	page.Execute(w, state)
}

func selectLevelHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {

	case http.MethodGet:
		http.ServeFile(w, r, "views/select-level.html")

	case http.MethodPost:
		r.ParseForm()
		state.Level = r.FormValue("level")
		http.Redirect(w, r, "/", http.StatusMovedPermanently)

	default:
		http.Redirect(w, r, "/", http.StatusMovedPermanently)
	}
}

func gameHandler(w http.ResponseWriter, r *http.Request) {
	page := template.Must(template.ParseFiles("views/game.html"))

	switch r.Method {
	case http.MethodGet:
		word := getNewWord(state.Level)
		state.CompleteWord = word
		state.Letters = initializeLetters()
		state.Errors = 0
		state.CurrentWord = initializeCurrentWord(len(word))
		page.Execute(w, state)

	case http.MethodPost:
		page.Execute(w, state)
	}
}

func getNewWord(level string) string {
	file, _ := os.Open("files/" + level + ".txt")
	scanner := bufio.NewScanner(file)
	var words []string

	for scanner.Scan() {
		words = append(words, scanner.Text())
	}

	rand.Seed(time.Now().UnixNano())
	num := rand.Intn(len(words))

	return words[num]
}

func initializeLetters() []Letter {
	var letters []Letter
	a := 97
	total := 26

	for i := a; i < a+total; i++ {
		letters = append(letters, Letter{Value: string(i), Used: false})
	}

	return letters
}

func initializeCurrentWord(n int) []string {
	var s []string

	for i := 0; i < n; i++ {
		s = append(s, "_")
	}

	return s
}

func main() {
	state.Level = "jrpgs"

	fileserver := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static", fileserver))

	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/select-level", selectLevelHandler)
	http.HandleFunc("/play", gameHandler)
	http.ListenAndServe(":8080", nil)
}
