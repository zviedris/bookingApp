package forms

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/asaskevich/govalidator"
)

//Form creates a custom form struct
type Form struct {
	url.Values
	Errors errors
}

//New initilaizes a form structure
func New(data url.Values) *Form {
	return &Form{
		data,
		errors(map[string][]string{}),
	}
}

//Has checks if a form is in post and not an empty field
func (f *Form) Has(field string, r *http.Request) bool {
	x := r.PostForm.Get(field)
	if x == "" {
		f.Errors.Add(field, "This field cannot be blank")
		return false
	}
	return true
}

//Valid returns true if there is no errors otherwise false
func (f *Form) Valid() bool {
	return len(f.Errors) == 0
}

func (f *Form) Required(fields ...string) {
	for _, field := range fields {
		value := f.Get(field)
		if strings.TrimSpace(value) == "" {
			f.Errors.Add(field, "This field cannot be blank")
		}
	}
}

//MinLength check for minimal length
func (f *Form) MinLength(field string, length int, r *http.Request) bool {
	x := r.PostForm.Get(field)
	if len(x) < length {
		f.Errors.Add(field, fmt.Sprintf("Field should be at least %d in length", length))
		return false
	}
	return true
}

//IsEmail check for a valid email
func (f *Form) IsEmail(field string) {
	if !govalidator.IsEmail(f.Get(field)) {
		f.Errors.Add(field, "Invalid email account")
	}
}
