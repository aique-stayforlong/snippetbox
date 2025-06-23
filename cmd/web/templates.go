package main

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
	"snippetbox.asuarez.net/internal/repository"
	"time"
)

type templateData struct {
	CurrentYear int
	Snippet     repository.Snippet
	Snippets    []repository.Snippet
}

func (app *application) newTemplateData(r *http.Request) templateData {
	return templateData{
		CurrentYear: time.Now().Year(),
	}
}

func humanDate(t time.Time) string {
	return t.Format("02 Jan 2006 at 15:04")
}

var functions = template.FuncMap{
	"humanDate": humanDate,
}

func newTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	pages, err := filepath.Glob("./ui/html/pages/*.tmpl.html")
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		// parse the base file
		ts, err := template.New(name).Funcs(functions).ParseFiles("./ui/html/base.tmpl.html")
		if err != nil {
			return nil, err
		}

		// parse any partials file
		ts, err = ts.ParseGlob("./ui/html/partials/*.tmpl.html")
		if err != nil {
			return nil, err
		}

		// parse current page template
		ts, err = ts.ParseFiles(page)
		if err != nil {
			return nil, err
		}

		cache[name] = ts // add the template set to the map
	}

	return cache, nil
}

func (app *application) render(w http.ResponseWriter, r *http.Request, status int, page string, data templateData) {
	ts, ok := app.templateCache[page]
	if !ok {
		err := fmt.Errorf("the template %s does not exist", page)
		app.serverError(w, r, err)
		return
	}

	// compile templates and check for errors
	buf := new(bytes.Buffer)
	err := ts.ExecuteTemplate(buf, "base", data)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	// write request status code
	w.WriteHeader(status)

	// write buffer template content to the standard output
	_, err = buf.WriteTo(w)
	if err != nil {
		app.serverError(w, r, err)
	}
}
