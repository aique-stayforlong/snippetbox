package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"
)

func main() {
	// define command line flags
	addr := flag.String("addr", ":4000", "HTTP network address")
	// parse command line flags
	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

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

	logger.Info("starting server", slog.String("addr", *addr))

	err := http.ListenAndServe(*addr, mux)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
}
