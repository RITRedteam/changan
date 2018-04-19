// Coder: koalatea
// Email: koalateac@gmail.com

package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (app *App) NewRouter() *mux.Router {
	routes := app.GenerateRoutes()

	router := mux.NewRouter().StrictSlash(true)
	fileServer := http.FileServer(http.Dir(app.StaticDir))
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static", fileServer))
	for _, route := range routes {
		var handler http.Handler

		if route.API {
			handler = route.HandlerFunc
		} else {
			handler = NoSurf(route.HandlerFunc)
		}
		if route.Auth != nil {
			// TODO RBAC with configurable roles
			handler = app.RequireLogin(handler) // TODO add RBAC
		}
		// TODO reenable
		handler = app.LogRequest(SecureHeaders(handler), route.Name)

		router.Methods(route.Method).Path(route.Pattern).Name(route.Name).Handler(handler)
	}

	return router

}
