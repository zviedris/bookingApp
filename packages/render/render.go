package render

import (
	"bytes"
	"log"
	"net/http"
	"path/filepath"
	"text/template"

	"github.com/zviedris/bookings/packages/config"
	"github.com/zviedris/bookings/packages/models"
)

var functions = template.FuncMap{}

var app *config.AppConfig

func AddDefaultData(td *models.TemplateData) *models.TemplateData {
	return td
}

//NewTemplates sets the config for the template package
func NewTemplates(a *config.AppConfig) {
	app = a
}

func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {
	var tc map[string]*template.Template

	//get template cache from the app config
	if app.UseCache {
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}

	templatePage, ok := tc[tmpl]
	if !ok {
		log.Fatal("Page not found")
	}

	//add default data to template data
	td = AddDefaultData(td)

	buf := new(bytes.Buffer)
	_ = templatePage.Execute(buf, td)

	_, err := buf.WriteTo(w)
	if err != nil {
		log.Fatal(err)
	}

	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)
	err = parsedTemplate.Execute(w, nil)
	if err != nil {
		log.Println("error parsing templat:", err)
	}
}

//CreateTemplateCache creates a template cache as a map
func CreateTemplateCache() (map[string]*template.Template, error) {
	{

		myCache := map[string]*template.Template{}

		pages, err := filepath.Glob(("./templates/*page.tmpl"))
		if err != nil {
			return myCache, err
		}

		for _, page := range pages {
			name := filepath.Base(page)
			ts, err := template.New(name).Funcs(functions).ParseFiles(page)
			if err != nil {
				return myCache, err
			}

			matches, err := filepath.Glob("./templates/*.layout.tmpl")
			if err != nil {
				return myCache, err
			}

			if len(matches) > 0 {
				ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
				if err != nil {
					return myCache, err
				}
			}

			myCache[name] = ts
		}

		return myCache, nil
	}
}
