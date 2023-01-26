package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"html/template"
	"github.com/go-redis/redis"
	"github.com/gorilla/sessions"
)

var templates *template.Template
var store = sessions.NewCookieStore([]byte("t0p-s3cr3t"))

var client *redis.Client

func main(){
	client = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	templates = template.Must(template.ParseGlob("templates/*.html"))
	r := mux.NewRouter()


	r.HandleFunc("/", indexhandler).Methods("GET")
	r.HandleFunc("/", indexPostHandler).Methods("POST")
	r.HandleFunc("/login", loginHandler).Methods("GET")
	r.HandleFunc("/login", loginPostHandler).Methods("POST")

	fs := http.FileServer(http.Dir("./static/"))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))
	http.Handle("/", r)
	http.ListenAndServe(":2500", nil)
}

func indexhandler(w http.ResponseWriter, r *http.Request){
	comments, err := client.LRange("comments", 0, 10).Result()
	if err != nil {
		panic(err)
	}


	templates.ExecuteTemplate(w, "index.html", comments)
}


func indexPostHandler(w http.ResponseWriter, r *http.Request){
	r.ParseForm()
	comment := r.PostForm.Get("comment")
	client.LPush("comments", comment)
	http.Redirect(w, r, "/", 302)
}

func loginHandler(w http.ResponseWriter, r *http.Request){
	templates.ExecuteTemplate(w, "login.html", nil)
}

func loginPostHandler(w http.ResponseWriter, r *http.Request){
	r.ParseForm()
	username := r.PostForm.Get("username")
	session, _ := store.Get(r, "session")
	session.Values["username"] = username
	session.Save(r, w)
}

func testSession(w http.ResponseWriter, r *http.Request){
	session, _ := store.Get(r, "session")

	untyped, ok := session.Values["username"]

	if !ok{
		return
	}

	username, ok := untyped.(string)
	w.Write([]byte(username))
}