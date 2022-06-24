package handlers

import (
	"go-server-basic/pkg/config"
	"go-server-basic/pkg/render"
	"net/http"
)

// Repo is the repository used by the handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.Appconfig
}

//NewRepo creates a new repository
func NewRepo(a *config.Appconfig) *Repository{

	return &Repository{
		App : a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers (r Repository){
	Repo = &r
}

// Home is the handler for the home page
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl")
}

// About is the handler for the about page
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.tmpl")
}
