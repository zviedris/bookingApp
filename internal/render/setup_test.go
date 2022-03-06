package render

import (
	"encoding/gob"
	"log"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/zviedris/bookings/internal/config"
	"github.com/zviedris/bookings/internal/models"
)

var s1 *scs.SessionManager
var testApp config.AppConfig

func TestMain(m *testing.M) {

	//change to true when production
	testApp.InProduction = false

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	testApp.InfoLog = infoLog

	errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	testApp.ErrorLog = errorLog

	//add value to store in session
	gob.Register(models.Reservation{})

	s1 = scs.New()
	s1.Lifetime = 24 * time.Hour
	s1.Cookie.Persist = true
	s1.Cookie.SameSite = http.SameSiteLaxMode
	s1.Cookie.Secure = false

	testApp.Session = s1

	app = &testApp

	os.Exit(m.Run())
}

type myWriter struct {
}

func (mw *myWriter) Header() http.Header {
	var h http.Header
	return h
}

func (mx *myWriter) WriteHeader(i int) {

}

func (mx *myWriter) Write(b []byte) (int, error) {
	l := len(b)
	return l, nil
}
