package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"html/template"
	"github.com/go-redis/redis"
)

var templates *template.Template

var client *redis.Client

func main(){
	client = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	templates = template.Must(template.ParseGlob("templates/*.html"))
	r := mux.NewRouter()


	r.HandleFunc("/", indexhandler).Methods("GET")
	
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
