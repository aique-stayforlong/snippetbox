package repository

import (
	"database/sql"
	"time"
)

type Snippet struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}

type SnippetRepository struct {
	DB *sql.DB
}

func (m *SnippetRepository) Insert(title string, content string, expires int) (int, error) {
	return 0, nil
}

func (m *SnippetRepository) Get(id int) (Snippet, error) {
	return Snippet{}, nil
}

func (m *SnippetRepository) Latest() ([]Snippet, error) {
	return nil, nil
}
