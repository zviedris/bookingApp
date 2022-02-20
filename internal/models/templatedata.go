package models

import "github.com/zviedris/bookings/internal/forms"

//Data that is passed to template data
type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]int
	Float     map[string]float32
	Data      map[string]interface{}
	CSRFToken string
	Flash     string
	Warning   string
	Error     string
	Form      *forms.Form
}
