package handlers

import (
	"net/http"

	"github.com/zviedris/bookings/packages/config"
	"github.com/zviedris/bookings/packages/models"
	"github.com/zviedris/bookings/packages/render"
)

//Repo the repository used by the handlers
var Repo *Repository

//Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

//NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

//NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

//Home is the home page handler
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIp := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIp)
	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

//About is the about page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, agian."
	remoteIp := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIp

	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

//Home is the home page handler
func (m *Repository) Forest(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "forestroom.page.tmpl", &models.TemplateData{})
}

//Home is the home page handler
func (m *Repository) Sea(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "searoom.page.tmpl", &models.TemplateData{})
}

//Home is the home page handler
func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "contact.page.tmpl", &models.TemplateData{})
}

//Home is the home page handler
func (m *Repository) Booknow(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "booknow.page.tmpl", &models.TemplateData{})
}
