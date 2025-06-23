package main

import (
	"database/sql"
	"flag"
	"fmt"
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
	port := flag.String("port", "3308", "Database port")

	flag.Parse() // get command line flags

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	// database connection
	dsn := flag.String("dsn", fmt.Sprintf("web:pass@tcp(127.0.0.1:%s)/snippetbox?parseTime=true", *port), "MySQL data source name")

	db, err := openDB(*dsn)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	defer db.Close() // closes database right before exit

	// init dependency injection container
	app := application{
		logger:   logger,
		snippets: &repository.SnippetRepository{DB: db},
	}

	// init logger dependency
	app.logger.Info("starting server", slog.String("addr", *addr))

	routes := app.routes() // set server routes

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
