package forms

import (
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestForm_Valid(t *testing.T) {
	r := httptest.NewRequest("POST", "/testw", nil)
	form := New(r.PostForm)

	isValid := form.Valid()
	if !isValid {
		t.Error("got invalid when should have been valid")
	}

}

func TestForm_Valid_False(t *testing.T) {
	postdata := url.Values{}
	postdata.Add("name", "Smith")

	r := httptest.NewRequest("POST", "/testw", nil)
	r.PostForm = postdata

	form := New(r.PostForm)

	form.Required("name")

	isValid := form.Valid()
	if !isValid {
		t.Error("should found required field, but got error")
	}

	form.Required("a")

	isValid = form.Valid()
	if isValid {
		t.Error("should found error as missing required field")
	}
}

func TestForm_Has(t *testing.T) {
	postdata := url.Values{}
	postdata.Add("name", "Smith")

	r := httptest.NewRequest("POST", "/testw", nil)
	r.PostForm = postdata

	form := New(r.PostForm)
	testVal := form.Has("name", r)
	if !testVal {
		t.Error("got has no field should have got field")
	}

	testVal = form.Has("surname", r)
	if testVal {
		t.Error("got has field shouldn 't have field")
	}

}

func TestForm_MinLength(t *testing.T) {
	postdata := url.Values{}
	postdata.Add("name", "Smith")

	r := httptest.NewRequest("POST", "/testw", nil)
	r.PostForm = postdata

	form := New(r.PostForm)
	testVal := form.MinLength("name", 6, r)
	if testVal {
		t.Error("Min length function not finding too short string")
	}

	testVal = form.MinLength("name", 4, r)
	if !testVal {
		t.Error("Min length function error on correct string")
	}

}

func TestForm_Email(t *testing.T) {
	postdata := url.Values{}
	postdata.Add("email", "Smith@example.com")
	postdata.Add("name", "Jonh@smith")

	r := httptest.NewRequest("POST", "/testw", nil)
	r.PostForm = postdata

	form := New(r.PostForm)
	form.IsEmail("email")
	if len(form.Errors) > 0 {
		t.Error("Should verify email, got error")
	}

	form.IsEmail("name")
	if len(form.Errors) == 0 {
		t.Error("Should found email error got correct")
	}

}
