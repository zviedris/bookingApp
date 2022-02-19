package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/zviedris/bookings/internal/config"
	"github.com/zviedris/bookings/internal/handlers"
	"github.com/zviedris/bookings/internal/render"
)

const portNumber = ":8080"

var app config.AppConfig
var s1 *scs.SessionManager

func main() {
	var app config.AppConfig

	//change to true when production
	app.InProduction = false

	s1 = scs.New()
	s1.Lifetime = 24 * time.Hour
	s1.Cookie.Persist = true
	s1.Cookie.SameSite = http.SameSiteLaxMode
	s1.Cookie.Secure = app.InProduction

	app.Session = s1

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	//create a new repo for handlers
	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	//http.HandleFunc("/", handlers.Repo.Home)
	//http.HandleFunc("/about", handlers.Repo.About)

	fmt.Println(fmt.Sprintf("Port number is %s", portNumber))
	//http.ListenAndServe(portNumber, nil)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}
