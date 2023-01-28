package controllers

import (
	"net/http"
	"github.com/dee-d-dev/go-server/models"
	"github.com/dee-d-dev/go-server/utils"
	"github.com/dee-d-dev/go-server/sessions"
)

func Indexhandler(w http.ResponseWriter, r *http.Request) {

	comments, err := models.GetComments()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
		return
	}

	utils.ExecuteTemplate(w, "index.html", comments)

}

func IndexPostHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	comment := r.PostForm.Get("comment")
	err := models.PostComment(comment)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
		return
	}
	http.Redirect(w, r, "/", 302)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	utils.ExecuteTemplate(w, "login.html", nil)
}

func LoginPostHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	username := r.PostForm.Get("username")
	password := r.PostForm.Get("password")
	err := models.AuthenticateUser(username, password)

	if err != nil {
		switch err {
		case models.ErrNotFound:
			utils.ExecuteTemplate(w, "login.html", "Unknown User")
		case models.ErrInvalidLogin:
			utils.ExecuteTemplate(w, "login.html", "Invalid Login")
		default:
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Internal Server Error"))
		}
		return
	}

	session, _ := sessions.Store.Get(r, "session-name")
	session.Values["username"] = username
	session.Save(r, w)
	http.Redirect(w, r, "/", 302)
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	utils.ExecuteTemplate(w, "register.html", nil)
}

func RegisterPostHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	username := r.PostForm.Get("username")
	password := r.PostForm.Get("password")

	err := models.RegisterUser(username, password)

	//saved username to redis
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server error"))
		return
	}
	http.Redirect(w, r, "/login", 302)
}
