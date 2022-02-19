package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/zviedris/bookings/internal/config"
	"github.com/zviedris/bookings/internal/models"
	"github.com/zviedris/bookings/internal/render"
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
	render.RenderTemplate(w, r, "home.page.tmpl", &models.TemplateData{})
}

//About is the about page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, agian."
	remoteIp := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIp

	render.RenderTemplate(w, r, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

//Home is the home page handler
func (m *Repository) Forest(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "forestroom.page.tmpl", &models.TemplateData{})
}

//Home is the home page handler
func (m *Repository) Sea(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "searoom.page.tmpl", &models.TemplateData{})
}

//Home is the home page handler
func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "contact.page.tmpl", &models.TemplateData{})
}

//Home is the home page handler
func (m *Repository) SearchRoom(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "booknow.page.tmpl", &models.TemplateData{})
}

//Post  is the home page handler
func (m *Repository) PostSearchRoom(w http.ResponseWriter, r *http.Request) {
	startDateString := r.Form.Get("start")
	endDateString := r.Form.Get("end")
	w.Write([]byte(fmt.Sprintf("Request to search from %s till day %s", startDateString, endDateString)))
}

type availabilityJsonResponse struct {
	Ok      bool   `json:"ok"`
	Message string `json:"message"`
}

//Post  is the home page handler
func (m *Repository) AvailabilityJson(w http.ResponseWriter, r *http.Request) {
	resp := availabilityJsonResponse{
		Ok:      true,
		Message: "Available",
	}

	out, err := json.MarshalIndent(resp, "", "   ")
	if err != nil {
		log.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}
