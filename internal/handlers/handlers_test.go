package handlers

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

type postDate struct {
	key   string
	value string
}

var theTests = []struct {
	name               string
	url                string
	method             string
	params             []postDate
	expectedStatusCode int
}{
	{"home", "/", "GET", []postDate{}, http.StatusOK},
	{"about", "/about", "GET", []postDate{}, http.StatusOK},
	{"forestroom", "/forestroom", "GET", []postDate{}, http.StatusOK},
	{"searoom", "/searoom", "GET", []postDate{}, http.StatusOK},
	{"contact", "/contact", "GET", []postDate{}, http.StatusOK},
	{"searchroom", "/searchroom", "GET", []postDate{}, http.StatusOK},
	{"make-reservation", "/make-reservation", "GET", []postDate{}, http.StatusOK},
	{"reservation-summary", "/reservation-summary", "GET", []postDate{}, http.StatusOK},
	{"searchroom", "/searchroom", "POST", []postDate{
		{key: "start", value: "2020-01-01"},
		{key: "end", value: "2020-01-01"},
	}, http.StatusOK},
	{"searchroom-availability", "/searchroom-availability", "POST", []postDate{
		{key: "start", value: "2020-01-01"},
		{key: "end", value: "2020-01-01"},
	}, http.StatusOK},
	{"make-reservation", "/make-reservation", "POST", []postDate{
		{key: "fist_name", value: "John"},
		{key: "last_name", value: "Smith"},
		{key: "email", value: "me@example.com"},
		{key: "phone", value: "252000000"},
	}, http.StatusOK},
}

func TestHandlers(t *testing.T) {
	routes := getRoutes()
	ts := httptest.NewTLSServer(routes)
	defer ts.Close()

	for _, e := range theTests {
		if e.method == "GET" {
			resp, err := ts.Client().Get(ts.URL + e.url)
			if err != nil {
				t.Log(err)
				t.Fatal(err)
			}
			if resp.StatusCode != e.expectedStatusCode {
				t.Errorf(fmt.Sprintf("For %s, expected status %d, got status %d", e.name, e.expectedStatusCode, resp.StatusCode))
			}
		} else {
			values := url.Values{}
			for _, x := range e.params {
				values.Add(x.key, x.value)
			}
			resp, err := ts.Client().PostForm(ts.URL+e.url, values)
			if err != nil {
				t.Log(err)
				t.Fatal(err)
			}
			if resp.StatusCode != e.expectedStatusCode {
				t.Errorf(fmt.Sprintf("For %s, expected status %d, got status %d", e.name, e.expectedStatusCode, resp.StatusCode))
			}
		}
	}
}
