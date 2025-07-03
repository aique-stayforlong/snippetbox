package main

import "net/http"

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()

	// static files server
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	// removes /static prefix to get the file path
	filePath := http.StripPrefix("/static", fileServer)
	// returns the static file
	mux.Handle("GET /static/", filePath)

	// regular http routes
	mux.HandleFunc("GET /{$}", app.home)
	mux.HandleFunc("GET /snippet/view/{id}", app.snippetView)
	mux.HandleFunc("GET /snippet/create", app.snippetCreate)
	mux.HandleFunc("POST /snippet/create", app.snippetCreatePost)

	return commonHeaders(mux)
}
