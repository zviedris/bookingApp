package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/zviedris/bookings/packages/config"
	"github.com/zviedris/bookings/packages/handlers"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)
	mux.Get("/forestroom", handlers.Repo.Forest)
	mux.Get("/searoom", handlers.Repo.Sea)
	mux.Get("/contact", handlers.Repo.Contact)
	mux.Get("/booknow", handlers.Repo.Booknow)

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))
	return mux
}
