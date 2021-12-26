package main

import (
	"net/http"
)

type testHandler struct {
	Message string
}

func (f *testHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(f.Message))
}

func main() {
	http.Handle("/", &testHandler{Message: "homepage"})
	http.ListenAndServe(":8080", nil)
}
