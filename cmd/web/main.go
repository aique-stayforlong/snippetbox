package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	// static files server
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	// removes /static prefix to get the file path
	filePath := http.StripPrefix("/static", fileServer)
	// returns the static file
	mux.Handle("GET /static/", filePath)

	// regular http routes
	mux.HandleFunc("GET /{$}", home)
	mux.HandleFunc("GET /snippet/view/{id}", snippetView)
	mux.HandleFunc("GET /snippet/create", snippetCreate)
	mux.HandleFunc("POST /snippet/create", snippetCreatePost)

	log.Print("starting server on :4000")

	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
