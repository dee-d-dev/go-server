package controllers

import (
	"net/http"

	"github.com/dee-d-dev/go-server/models"
	"github.com/dee-d-dev/go-server/sessions"
	"github.com/dee-d-dev/go-server/utils"
	"github.com/gorilla/mux"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {

	updates, err := models.GetAllUpdates()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
		return
	}

	utils.ExecuteTemplate(w, "index.html", struct {
		Title string
		Updates []*models.Update
	}{
		Title: "All updates",
		Updates: updates,
	})

}

func IndexPostHandler(w http.ResponseWriter, r *http.Request) {

	session, _ := sessions.Store.Get(r, "session-name")
	untypedUserId := session.Values["user_id"]
	userId, ok := untypedUserId.(int64)

	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
		return
	}

	r.ParseForm()

	body := r.PostForm.Get("update")
	err := models.PostUpdate(userId, body)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
		return
	}
	http.Redirect(w, r, "/", 302)
}

func UserGetHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	username := vars["username"]
	user, err := models.GetUserByUsername(username)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
		return
	}
	userId, err := user.GetId()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
		return
	}
	updates, err := models.GetUpdates(userId)
	
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
		return
	}

	utils.ExecuteTemplate(w, "index.html",struct {
		Title string
		Updates []*models.Update
	}{
		Title: username,
		Updates: updates,
	})

}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	utils.ExecuteTemplate(w, "login.html", nil)
}

func LoginPostHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	username := r.PostForm.Get("username")
	password := r.PostForm.Get("password")
	user, err := models.AuthenticateUser(username, password)

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

	userId, err := user.GetId()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
		return
	}

	session, _ := sessions.Store.Get(r, "session-name")
	session.Values["user_id"] = userId
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
