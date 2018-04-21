package main

import (
	"log"
	"net"
	"net/http"
	"time"

	"github.com/justinas/nosurf"
)

func (app *App) RequireLogin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// wait why do I need to call this? is it not a bool? (HTMLDATA)
		loggedIn, err := app.LoggedIn(r)
		if err != nil {
			app.ServerError(w, err)
			return
		}

		if !loggedIn {
			http.Redirect(w, r, "/user/login", 302) //correct not autherized code
			return                                  // return so rest of the chain is not execured
		}

		next.ServeHTTP(w, r)
	})
}

// unused function ( will probably merge both loggers ) logging needs to hit bad pages too
// for instance accessing a get /user/logout (we have no page for this but it does not 404)
// nor do I see the loging look at interactions between router.go and this might have to do with
// the chain or with how mux is set up
func LogRequest2(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		pattern := `%s - "%s %s %s"`
		log.Printf(pattern, r.RemoteAddr, r.Proto, r.Method, r.URL.RequestURI())

		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}

func (app *App) LogRequest(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		inner.ServeHTTP(w, r)
		ip, _, _ := net.SplitHostPort(r.RemoteAddr)
		log.Printf("%s\t%s\t%s\t%s\t%s", r.Method, r.RequestURI, name, time.Since(start), ip)
	})
}

func SecureHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.Header().Set("X-Frame-Options", "deny")
		w.Header()["X-XXS-Protection"] = []string{"1; mode=block"}

		next.ServeHTTP(w, r)
	})
}

func NoSurf(next http.HandlerFunc) http.Handler {
	csrfHandler := nosurf.New(next)
	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   true,
	})

	return csrfHandler
}
