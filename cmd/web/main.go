package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"
)

// dependency injection container
type application struct {
	logger *slog.Logger
}

func main() {
	// define command line flags
	addr := flag.String("addr", ":4000", "HTTP network address")
	// parse command line flags
	flag.Parse()

	// init dependency injection container
	app := application{
		logger: slog.New(slog.NewTextHandler(os.Stdout, nil)),
	}

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

	app.logger.Info("starting server", slog.String("addr", *addr))

	err := http.ListenAndServe(*addr, mux)
	if err != nil {
		app.logger.Error(err.Error())
		os.Exit(1)
	}
}
