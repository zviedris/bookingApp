package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/zviedris/bookings/internal/config"
	"github.com/zviedris/bookings/internal/forms"
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

//Post Reservation handles posting of the reservation form
func (m *Repository) Reservation(w http.ResponseWriter, r *http.Request) {
	var emptyReservation models.Reservation
	data := make(map[string]interface{})
	data["reservation"] = emptyReservation

	render.RenderTemplate(w, r, "makereservation.page.tmpl", &models.TemplateData{
		Form: forms.New(nil),
		Data: data,
	})
}

//Post Reservation handles posting of the reservation form
func (m *Repository) PostReservation(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
		return
	}

	reservation := models.Reservation{
		FirstName: r.Form.Get("first_name"),
		LastName:  r.Form.Get("last_name"),
		Phone:     r.Form.Get("phone"),
		Email:     r.Form.Get("email"),
	}

	form := forms.New(r.PostForm)

	form.Required("first_name", "last_name", "phone", "email")
	form.MinLength("first_name", 3, r)
	form.IsEmail("email")

	if !form.Valid() {
		data := make(map[string]interface{})
		data["reservation"] = reservation

		render.RenderTemplate(w, r, "makereservation.page.tmpl", &models.TemplateData{
			Form: form,
			Data: data,
		})
		return
	}

	m.App.Session.Put(r.Context(), "reservation", reservation)

	http.Redirect(w, r, "/reservation-summary", http.StatusSeeOther)
}

//Home is the home page handler
func (m *Repository) ReservationSummary(w http.ResponseWriter, r *http.Request) {
	reservation, ok := m.App.Session.Get(r.Context(), "reservation").(models.Reservation)
	if !ok {
		log.Println("Cannot get data from session")
		m.App.Session.Put(r.Context(), "error", "Cannot get data from session")
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	data := make(map[string]interface{})
	data["reservation"] = reservation

	render.RenderTemplate(w, r, "reservationsummary.page.tmpl", &models.TemplateData{
		Data: data,
	})
}
