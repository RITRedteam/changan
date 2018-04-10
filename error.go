package main

import (
	"encoding/json"
	"net/http"
	"runtime/debug"
)

// ServerError helper writes an error message and stack trace to the log, then returns a generic
// 500 Internal Server Error to the user
func (app *App) ServerError(w http.ResponseWriter, err error) {
	app.Logger.Errorf("%s\n%s", err.Error(), debug.Stack())
	http.Error(w, "Internal Server Error", http.StatusInternalServerError)
}

// ClientError helper sends a specific status code and corresponding description to the user.
func (app *App) ClientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

// NotFound helper is a convience wrapper around ClientError we will probably not use this with
// mux in the program.
func (app *App) NotFound(w http.ResponseWriter) {
	app.ClientError(w, http.StatusNotFound)
}

func (app *App) APIServerError(w http.ResponseWriter, err error) {
	app.Logger.Errorf("%s\n%s", err.Error(), debug.Stack())
	data := APIData{Error: err.Error()}
	js, err := json.Marshal(data)
	if err != nil {
		app.ServerError(w, err)
		return
	}
	w.WriteHeader(http.StatusInternalServerError)
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
