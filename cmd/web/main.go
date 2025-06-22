package main

import (
	"database/sql"
	"flag"
	_ "github.com/go-sql-driver/mysql"
	"log/slog"
	"net/http"
	"os"
	"snippetbox.asuarez.net/internal/repository"
)

// dependency injection container
type application struct {
	logger   *slog.Logger
	snippets *repository.SnippetRepository
}

func main() {
	// define command line flags
	addr := flag.String("addr", ":4000", "HTTP network address")
	// parse command line flags
	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	// database connection
	dsn := flag.String("dsn", "web:pass@/snippetbox?parseTime=true", "MySQL data source name")

	db, err := openDB(*dsn)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	// closes database right before exit
	defer db.Close()

	// init dependency injection container
	app := application{
		logger:   logger,
		snippets: &repository.SnippetRepository{DB: db},
	}

	// init logger dependency
	app.logger.Info("starting server", slog.String("addr", *addr))

	// set server routes
	routes := app.routes()

	// start server
	err = http.ListenAndServe(*addr, routes)
	if err != nil {
		app.logger.Error(err.Error())
		os.Exit(1)
	}
}

func openDB(dsn string) (*sql.DB, error) {
	// creates a new connection in the connection pool
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	// check if database can be reached
	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}
