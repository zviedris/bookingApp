package render

import (
	"bytes"
	"errors"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/justinas/nosurf"
	"github.com/zviedris/bookings/internal/config"
	"github.com/zviedris/bookings/internal/models"
)

var functions = template.FuncMap{}

var app *config.AppConfig
var pathToTemplates = "./templates/"

func AddDefaultData(td *models.TemplateData, r *http.Request) *models.TemplateData {
	td.Flash = app.Session.PopString(r.Context(), "flash")
	td.Error = app.Session.PopString(r.Context(), "error")
	td.Warning = app.Session.PopString(r.Context(), "warning")
	td.CSRFToken = nosurf.Token(r)
	return td
}

//NewTemplates sets the config for the template package
func NewTemplates(a *config.AppConfig) {
	app = a
}

func RenderTemplate(w http.ResponseWriter, r *http.Request, tmpl string, td *models.TemplateData) error {
	var tc map[string]*template.Template

	//get template cache from the app config
	if app.UseCache {
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}

	templatePage, ok := tc[tmpl]
	if !ok {
		//log.Fatal("Page not found")
		return errors.New("Can't get template from cache")
	}

	//add default data to template data
	td = AddDefaultData(td, r)

	buf := new(bytes.Buffer)
	_ = templatePage.Execute(buf, td)

	_, err := buf.WriteTo(w)
	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil

}

//CreateTemplateCache creates a template cache as a map
func CreateTemplateCache() (map[string]*template.Template, error) {
	{

		myCache := map[string]*template.Template{}

		pages, err := filepath.Glob(fmt.Sprintf("%s*page.tmpl", pathToTemplates))
		if err != nil {
			return myCache, err
		}

		for _, page := range pages {
			name := filepath.Base(page)
			ts, err := template.New(name).Funcs(functions).ParseFiles(page)
			if err != nil {
				return myCache, err
			}

			matches, err := filepath.Glob(fmt.Sprintf("%s*layout.tmpl", pathToTemplates))
			if err != nil {
				return myCache, err
			}

			if len(matches) > 0 {
				ts, err = ts.ParseGlob(fmt.Sprintf("%s*layout.tmpl", pathToTemplates))
				if err != nil {
					return myCache, err
				}
			}

			myCache[name] = ts
		}

		return myCache, nil
	}
}
