package main

import (
	"github.com/koalatea/changan/pkg/models"

	"github.com/alexedwards/scs"
	"github.com/juju/loggo"
)

// App is a struct to hold application-wide dependencies and configuration settings for our web app
type App struct {
	Addr      string
	Database  *models.SQLDatabase
	Mongo     *models.Database
	HTMLDir   string
	StaticDir string
	Sessions  *scs.Manager
	TLSCert   string
	TLSKey    string
	Logger    loggo.Logger
}
