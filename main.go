package main

import (
	
	"net/http"
	
	"github.com/dee-d-dev/go-server/models"
	"github.com/dee-d-dev/go-server/utils"
	"github.com/dee-d-dev/go-server/routes"
)





func main() {
	//redis connection
	models.InitRedis()
	utils.LoadTemplates("templates/*.html")
	r := routes.NewRouter()
	http.Handle("/", r)
	http.ListenAndServe(":2500", nil)
}

