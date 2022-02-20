package render

import (
	"net/http"
	"testing"

	"github.com/zviedris/bookings/internal/models"
)

func TestAddDefaultData(t *testing.T) {
	var td models.TemplateData

	r, err := getSession()
	if err != nil {
		t.Error(err)
		return
	}

	s1.Put(r.Context(), "flash", "123")

	result := AddDefaultData(&td, r)
	if result.Flash != "123" {
		t.Error("Cannot get flash from session")
	}

}

func TestRenderTemplate(t *testing.T) {
	pathToTemplates = "./../../templates/"
	tc, err := CreateTemplateCache()
	if err != nil {
		t.Error(err)
	}

	app.TemplateCache = tc
	r, err := getSession()
	if err != nil {
		t.Error(err)
	}

	var ww myWriter

	err = RenderTemplate(&ww, r, "home.page.tmpl", &models.TemplateData{})
	if err != nil {
		t.Error(err)
	}

	err = RenderTemplate(&ww, r, "non-existing-template.page.tmpl", &models.TemplateData{})
	if err == nil {
		t.Error("Found non existing template")
	}

}

func TestNewTemplates(t *testing.T) {
	NewTemplates(app)
}

func TestCreateTemplateCache(t *testing.T) {
	pathToTemplates = "./../../templates/"
	_, err := CreateTemplateCache()
	if err != nil {
		t.Error(err)
	}
	/*if len (){
		t.Error("Found non existing template")
	}*/
}

func getSession() (*http.Request, error) {
	r, err := http.NewRequest("Get", "/test-url", nil)
	if err != nil {
		return nil, err
	}

	ctx := r.Context()
	ctx, _ = s1.Load(ctx, r.Header.Get("X-Session"))
	r = r.WithContext(ctx)
	return r, nil
}
