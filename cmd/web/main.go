package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/zviedris/bookings/helpers"
	"github.com/zviedris/bookings/internal/config"
	"github.com/zviedris/bookings/internal/handlers"
	"github.com/zviedris/bookings/internal/models"
	"github.com/zviedris/bookings/internal/render"
)

const portNumber = ":8080"

var app config.AppConfig
var s1 *scs.SessionManager
var infoLog *log.Logger
var errorLog *log.Logger

func main() {
	err := run()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(fmt.Sprintf("Port number is %s", portNumber))
	//http.ListenAndServe(portNumber, nil)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}

func run() error {
	var app config.AppConfig

	//change to true when production
	app.InProduction = false

	infoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	app.InfoLog = infoLog

	errorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	app.ErrorLog = errorLog

	//add value to store in session
	gob.Register(models.Reservation{})

	s1 = scs.New()
	s1.Lifetime = 24 * time.Hour
	s1.Cookie.Persist = true
	s1.Cookie.SameSite = http.SameSiteLaxMode
	s1.Cookie.Secure = app.InProduction

	app.Session = s1

	tc, err := render.CreateTemplateCache()
	if err != nil {
		return err
	}

	app.TemplateCache = tc
	app.UseCache = false

	//create a new repo for handlers
	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)
	helpers.NewHelpers(&app)
	render.NewTemplates(&app)

	//http.HandleFunc("/", handlers.Repo.Home)
	//http.HandleFunc("/about", handlers.Repo.About)
	return nil
}
