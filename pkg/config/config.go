package config

import "html/template"

// Appconfig holds the application config
type Appconfig struct {
	UseCache      bool   
	TemplateCache map[string]*template.Template
}
