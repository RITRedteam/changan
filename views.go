package main

import (
	"bytes"
	"encoding/json"
	"html/template"
	"net/http"
	"path/filepath"
	"time"

	"github.com/justinas/nosurf"
	"github.com/koalatea/changan/pkg/models"
)

type HTMLData struct {
	CSRFToken string
	Form      interface{}
	Path      string
	LoggedIn  bool
	Devices   []models.Device
	Device    *models.Device
	Subnets   []models.Subnet
	Subnet    *models.Subnet
	Reports   []models.Report
	Report    *models.Report
	Users     []models.User
	User      *models.User
}

type APIData struct {
	Subnets []models.Subnet `json:"subnets"`
	Devices []models.Device `json:"devices"`
	Reports []models.Report `json:"reports"`
	Result  string          `json:"result"`
	Error   string          `json:"error"`
}

func humanDate(t time.Time) string {
	return t.Format("02 Jan 2006 at 15:04")
}

func (app *App) ReturnAPI(w http.ResponseWriter, r *http.Request, data *APIData) {
	js, err := json.Marshal(data)
	if err != nil {
		app.ServerError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func (app *App) RenderHTML(w http.ResponseWriter, r *http.Request, page string, data *HTMLData) {
	if data == nil {
		data = &HTMLData{}
	}

	data.Path = r.URL.Path

	data.CSRFToken = nosurf.Token(r)

	var err error
	data.LoggedIn, err = app.LoggedIn(r)
	if err != nil {
		app.ServerError(w, err)
		return
	}

	files := []string{
		filepath.Join(app.HTMLDir, "base.html"),
		filepath.Join(app.HTMLDir, page),
	}

	funcs := template.FuncMap{
		"humanDate": humanDate,
	}

	ts, err := template.New("").Funcs(funcs).ParseFiles(files...)
	if err != nil {
		app.ServerError(w, err)
		return
	}

	buf := new(bytes.Buffer)

	err = ts.ExecuteTemplate(buf, "base", data)
	if err != nil {
		app.ServerError(w, err)
		return
	}

	buf.WriteTo(w)
}
