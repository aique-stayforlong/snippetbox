package main

import "snippetbox.asuarez.net/internal/repository"

type templateData struct {
	Snippet  repository.Snippet
	Snippets []repository.Snippet
}
