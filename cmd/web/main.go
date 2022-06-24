package main

import (
	"fmt"
	"go-server-basic/pkg/config"
	"go-server-basic/pkg/handlers"
	"go-server-basic/pkg/render"
	"log"
	"net/http"
)

const portNumber = ":8080"

// main is the main application fuction
func main() {

	var app config.Appconfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create template cache ")
	}

	app.TemplateCache = tc
	app.UseCache = false 

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(*repo)

	render.NewTemplates(&app)

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)

	fmt.Printf("Starting application on port %s", portNumber)
	_ = http.ListenAndServe(portNumber, nil)

}
