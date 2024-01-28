package main

import (
	"fmt"
	"net/http"
	"urlshortner"
)

func main() {
	mux := defaultMux()

	pathsToUrl := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}

	mapHandler := urlshortner.MapHandler(pathsToUrl, mux)
	http.ListenAndServe(":8080", mapHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", check)
	return mux
}

func check(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello")
}
