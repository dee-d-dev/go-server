package routes

import (
	"net/http"

	"github.com/dee-d-dev/go-server/controllers"
	"github.com/dee-d-dev/go-server/middlewares"
	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/", middlewares.AuthRequired(controllers.Indexhandler)).Methods("GET")
	r.HandleFunc("/", middlewares.AuthRequired(controllers.IndexPostHandler)).Methods("POST")
	r.HandleFunc("/login", controllers.LoginHandler).Methods("GET")
	r.HandleFunc("/login", controllers.LoginPostHandler).Methods("POST")
	r.HandleFunc("/register", controllers.RegisterHandler).Methods("GET")
	r.HandleFunc("/register", controllers.RegisterPostHandler).Methods("POST")

	fs := http.FileServer(http.Dir("./static/"))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))

	return r
}

// func indexhandler(w http.ResponseWriter, r *http.Request) {

// 	comments, err := models.GetComments()
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		w.Write([]byte("Internal Server Error"))
// 		return
// 	}

// 	utils.ExecuteTemplate(w, "index.html", comments)

// }

// func indexPostHandler(w http.ResponseWriter, r *http.Request) {
// 	r.ParseForm()
// 	comment := r.PostForm.Get("comment")
// 	err := models.PostComment(comment)

// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		w.Write([]byte("Internal Server Error"))
// 		return
// 	}
// 	http.Redirect(w, r, "/", 302)
// }

// func loginHandler(w http.ResponseWriter, r *http.Request) {
// 	utils.ExecuteTemplate(w, "login.html", nil)
// }

// func loginPostHandler(w http.ResponseWriter, r *http.Request) {
// 	r.ParseForm()
// 	username := r.PostForm.Get("username")
// 	password := r.PostForm.Get("password")
// 	err := models.AuthenticateUser(username, password)

// 	if err != nil {
// 		switch err {
// 		case models.ErrNotFound:
// 			utils.ExecuteTemplate(w, "login.html", "Unknown User")
// 		case models.ErrInvalidLogin:
// 			utils.ExecuteTemplate(w, "login.html", "Invalid Login")
// 		default:
// 			w.WriteHeader(http.StatusInternalServerError)
// 			w.Write([]byte("Internal Server Error"))
// 		}
// 		return
// 	}

// 	session, _ := sessions.Store.Get(r, "session-name")
// 	session.Values["username"] = username
// 	session.Save(r, w)
// 	http.Redirect(w, r, "/", 302)
// }

// func registerHandler(w http.ResponseWriter, r *http.Request) {
// 	utils.ExecuteTemplate(w, "register.html", nil)
// }

// func registerPostHandler(w http.ResponseWriter, r *http.Request) {
// 	r.ParseForm()

// 	username := r.PostForm.Get("username")
// 	password := r.PostForm.Get("password")

// 	err := models.RegisterUser(username, password)

// 	//saved username to redis
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		w.Write([]byte("Internal Server error"))
// 		return
// 	}
// 	http.Redirect(w, r, "/login", 302)
// }
