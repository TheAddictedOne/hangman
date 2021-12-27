package main

import (
	"bufio"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"text/template"
	"time"
)

// ┌────────────────────────────────────────────────────────────┐
// │ Globals				             						│
// └────────────────────────────────────────────────────────────┘

var state State // This is the state of ze gaaaaaame

// ┌────────────────────────────────────────────────────────────┐
// │ Structs				             						│
// └────────────────────────────────────────────────────────────┘

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
	GameOver     string
}

// ┌────────────────────────────────────────────────────────────┐
// │ Route handlers			             						│
// └────────────────────────────────────────────────────────────┘

func homeHandler(w http.ResponseWriter, r *http.Request) {
	page := template.Must(template.ParseFiles("views/index.html"))
	page.Execute(w, state)
}

func selectLevelHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {

	// ┌────────────────────────────────┐
	// │ Show available levels			│
	// └────────────────────────────────┘
	case http.MethodGet:
		http.ServeFile(w, r, "views/select-level.html")

	// ┌────────────────────────────────┐
	// │ Set the level in global state	│
	// └────────────────────────────────┘
	case http.MethodPost:
		r.ParseForm()
		state.Level = r.FormValue("level")
		http.Redirect(w, r, "/", http.StatusMovedPermanently)

	default:
		http.Redirect(w, r, "/", http.StatusMovedPermanently)
	}
}

func playHandler(w http.ResponseWriter, r *http.Request) {
	page := template.Must(template.ParseFiles("views/game.html"))

	// ┌────────────────────────────────┐
	// │ Initiliaze the game			│
	// └────────────────────────────────┘
	switch r.Method {
	case http.MethodGet:
		word := getNewWord(state.Level)
		state.CompleteWord = word
		state.Letters = initializeLetters()
		state.Errors = 0
		state.CurrentWord = initializeCurrentWord(word)
		state.GameOver = ""
		page.Execute(w, state)

	// ┌────────────────────────────────┐
	// │ Read the letter sent by player │
	// └────────────────────────────────┘
	case http.MethodPost:
		r.ParseForm()
		letter := r.FormValue("letter")
		isError := true

		// Replace "_" with the letter from the player, if found
		for i, v := range state.CompleteWord {
			if string(v) == letter {
				isError = false
				state.CurrentWord[i] = letter
			}
			// If all letters from the word have been checked, the letter has not been found isError stays "true"
		}

		for i, v := range state.Letters {
			if v.Value == letter {
				state.Letters[i] = Letter{Value: v.Value, Used: true}
				break
			}
		}

		if isError {
			state.Errors++
		}

		switch isGameOver(state.CurrentWord, state.Errors) {
		case 2:
			for i, v := range state.Letters {
				state.Letters[i] = Letter{Value: v.Value, Used: true}
			}
			state.GameOver = "You lose! Game over"
			state.CurrentWord = getCompleteWord(state.CompleteWord)

		case 1:
			for i, v := range state.Letters {
				state.Letters[i] = Letter{Value: v.Value, Used: true}
			}
			state.GameOver = "You win! Game over"
		}

		page.Execute(w, state)
	}
}

// ┌────────────────────────────────────────────────────────────┐
// │ Utilities				             						│
// └────────────────────────────────────────────────────────────┘

func getNewWord(level string) string {
	file, _ := os.Open("files/" + level + ".txt")
	scanner := bufio.NewScanner(file)
	var words []string

	for scanner.Scan() {
		words = append(words, scanner.Text())
	}

	rand.Seed(time.Now().UnixNano())
	num := rand.Intn(len(words))

	return strings.ToLower(words[num])
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

func initializeCurrentWord(w string) []string {
	var s []string

	for i := 0; i < len(w); i++ {
		if w[i] == ' ' {
			s = append(s, " ")
		} else {
			s = append(s, "_")
		}
	}

	return s
}

func getCompleteWord(w string) []string {
	var s []string
	for _, letter := range w {
		s = append(s, string(letter))
	}

	return s
}

func isGameOver(word []string, errors int) int {
	if errors == 6 {
		return 2
	}

	for _, letter := range word {
		if letter == "_" {
			return 0
		}
	}
	return 1
}

// ┌────────────────────────────────────────────────────────────┐
// │ Main					             						│
// └────────────────────────────────────────────────────────────┘

func main() {
	// ┌────────────────────────────────┐
	// │ Initiliaze						│
	// └────────────────────────────────┘
	state.Level = "jrpgs"

	// ┌────────────────────────────────┐
	// │ Serve static files				│
	// └────────────────────────────────┘
	fileserver := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static", fileserver))

	// ┌────────────────────────────────┐
	// │ Routes							│
	// └────────────────────────────────┘
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/select-level", selectLevelHandler)
	http.HandleFunc("/play", playHandler)

	// ┌────────────────────────────────┐
	// │ Start the server				│
	// └────────────────────────────────┘
	http.ListenAndServe(":8080", nil)
}
