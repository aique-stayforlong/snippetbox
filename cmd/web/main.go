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

	// init logger dependency
	app.logger.Info("starting server", slog.String("addr", *addr))

	// set server routes
	routes := app.routes()

	// start server
	err := http.ListenAndServe(*addr, routes)
	if err != nil {
		app.logger.Error(err.Error())
		os.Exit(1)
	}
}
