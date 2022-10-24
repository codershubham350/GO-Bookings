package render

import (
	"bytes"
	"errors"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"time"

	"github.com/codershubham350/bookings/internal/config"
	"github.com/codershubham350/bookings/internal/models"
	"github.com/justinas/nosurf"
)

var functions = template.FuncMap{
	"humanDate":  HumanDate,
	"formatDate": FormatDate,
	"iterate":    Iterate,
	"add":        Add,
}

var app *config.AppConfig
var pathToTemplates = "./templates"

func Add(a, b int) int {
	return a + b
}

// Iterate returns a slice of ints, starting at 1, going to count
func Iterate(count int) []int {
	var i int
	var items []int
	for i = 0; i < count; i++ {
		items = append(items, i)
	}
	return items
}

// NewRenderer sets the config for the template package
func NewRenderer(a *config.AppConfig) {
	app = a
}

// HumanDate returns time in YYYY-MM-DD format
func HumanDate(t time.Time) string {
	return t.Format("2006-01-02")
}

// FormatDate
func FormatDate(t time.Time, f string) string {
	return t.Format(f)
}

func AddDefaultData(td *models.TemplateData, r *http.Request) *models.TemplateData {
	td.Flash = app.Session.PopString(r.Context(), "flash")
	td.Error = app.Session.PopString(r.Context(), "eror")
	td.Warning = app.Session.PopString(r.Context(), "warning")
	td.CSRFToken = nosurf.Token(r)
	if app.Session.Exists(r.Context(), "user_id") {
		td.IsAuthenticated = 1
	}
	return td
}

// Template renders template using tmpl/template
func Template(w http.ResponseWriter, r *http.Request, tmpl string, td *models.TemplateData) error {
	var tc map[string]*template.Template
	if app.UseCache {
		// get the template cache from the app config
		tc = app.TemplateCache

	} else {
		tc, _ = CreateTemplateCache()
	}

	// create a template cache
	// tc, err := CreateTemplateCache()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// get requested template cache
	t, ok := tc[tmpl]
	if !ok {
		// log.Fatal("Could not get template from template")
		return errors.New("can't get template from cache")
	}

	buf := new(bytes.Buffer)

	td = AddDefaultData(td, r)
	err := t.Execute(buf, td)
	if err != nil {
		log.Println(err)
	}
	// render the template
	_, err = buf.WriteTo(w)
	if err != nil {
		log.Println("Error writing template to browser", err)
		return err
	}
	return nil
}

// CreateTemplateCache creates a template cache as map
func CreateTemplateCache() (map[string]*template.Template, error) {
	//myCache := make(map[string]*template.Template)
	myCache := map[string]*template.Template{}

	// get all of the files named *.page.tmpl from ./templates

	pages, err := filepath.Glob(fmt.Sprintf("%s/*.page.tmpl", pathToTemplates))
	if err != nil {
		return myCache, err
	}

	// range through all of the files ending with *.page.tmpl
	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob(fmt.Sprintf("%s/*.layout.tmpl", pathToTemplates))
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob(fmt.Sprintf("%s/*.layout.tmpl", pathToTemplates))
			if err != nil {
				return myCache, err
			}

		}

		myCache[name] = ts
	}

	return myCache, nil
}
