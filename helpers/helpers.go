package helpers

import (
	"fmt"
	"net/http"
	"runtime/debug"

	"github.com/zviedris/bookings/internal/config"
)

var app *config.AppConfig

//NewHelpers initialize app config
func NewHelpers(a *config.AppConfig) {
	app = a
}

//Error that just needs to be sent to client
func ClientError(w http.ResponseWriter, status int) {
	app.InfoLog.Println("Client error with status of", status)
	http.Error(w, http.StatusText(status), status)
}

//Error when something went wrong in server
func ServerError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.ErrorLog.Println(trace)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}
