package middlewares

import (
	"net/http"
	"github.com/dee-d-dev/go-server/sessions"
)
func AuthRequired(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, _ := sessions.Store.Get(r, "session-name")
		_, ok := session.Values["username"]

		if !ok {
			http.Redirect(w, r, "/login", 302)
			return
		}
		handler.ServeHTTP(w, r)
	}
}