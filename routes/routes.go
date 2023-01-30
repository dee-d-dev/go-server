package routes

import (
	"net/http"

	"github.com/dee-d-dev/go-server/controllers"
	"github.com/dee-d-dev/go-server/middlewares"
	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/", middlewares.AuthRequired(controllers.IndexHandler)).Methods("GET")
	r.HandleFunc("/", middlewares.AuthRequired(controllers.IndexPostHandler)).Methods("POST")
	r.HandleFunc("/login", controllers.LoginHandler).Methods("GET")
	r.HandleFunc("/login", controllers.LoginPostHandler).Methods("POST")
	r.HandleFunc("/register", controllers.RegisterHandler).Methods("GET")
	r.HandleFunc("/register", controllers.RegisterPostHandler).Methods("POST")

	fs := http.FileServer(http.Dir("./static/"))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))

	return r
}
