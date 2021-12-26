package main

import (
	"net/http"
	"text/template"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	homepage := template.Must(template.ParseFiles("views/index.html"))

	switch r.Method {
	case http.MethodGet:
		homepage.Execute(w, nil)
	default:
		w.Write([]byte("Method not handled"))
	}
}

func gameHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		// Show game
	case http.MethodPost:
		// Get letter
	}
}

func main() {
	fileserver := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static", fileserver))

	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/game", gameHandler)
	http.ListenAndServe(":8080", nil)
}
