package handlers

import (
	"encoding/gob"
	"fmt"

	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/justinas/nosurf"
	"github.com/zviedris/bookings/internal/config"
	"github.com/zviedris/bookings/internal/models"
	"github.com/zviedris/bookings/internal/render"
)

var app config.AppConfig
var s1 *scs.SessionManager
var pathToTemplates = "./../../templates/"
var functions = template.FuncMap{}

func getRoutes() http.Handler {

	//change to true when production
	app.InProduction = false

	//add value to store in session
	gob.Register(models.Reservation{})

	s1 = scs.New()
	s1.Lifetime = 24 * time.Hour
	s1.Cookie.Persist = true
	s1.Cookie.SameSite = http.SameSiteLaxMode
	s1.Cookie.Secure = app.InProduction

	app.Session = s1

	tc, err := CreateTestTemplateCache()
	if err != nil {
		log.Fatal("cannont create template cache")
		//return err
	}

	app.TemplateCache = tc
	//set to true otherwise will call wrong template dir
	app.UseCache = true

	//create a new repo for handlers
	repo := NewRepo(&app)
	NewHandlers(repo)

	render.NewTemplates(&app)

	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	//mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", Repo.Home)
	mux.Get("/about", Repo.About)
	mux.Get("/forestroom", Repo.Forest)
	mux.Get("/searoom", Repo.Sea)
	mux.Get("/contact", Repo.Contact)
	mux.Get("/searchroom", Repo.SearchRoom)
	mux.Post("/searchroom", Repo.PostSearchRoom)

	mux.Post("/searchroom-availability", Repo.AvailabilityJson)

	mux.Get("/make-reservation", Repo.Reservation)
	mux.Post("/make-reservation", Repo.PostReservation)

	mux.Get("/reservation-summary", Repo.ReservationSummary)

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))
	return mux
}

//NoSurf adds CSRF cookies
func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)

	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Secure:   app.InProduction,
		Path:     "/",
		SameSite: http.SameSiteLaxMode,
	})

	return csrfHandler
}

// SessionLoad loads and saves session data for current request
func SessionLoad(next http.Handler) http.Handler {
	return s1.LoadAndSave(next)
}

//CreateTemplateCache creates a template cache as a map
func CreateTestTemplateCache() (map[string]*template.Template, error) {
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
