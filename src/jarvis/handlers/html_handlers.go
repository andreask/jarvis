package handlers

import (
	"net/http"
	"os"
	"html/template"
)

func loadWeather() ()

func Index(w http.ResponseWriter, r *http.Request) {
	files := []string{os.Getenv("GOPATH") + "/templates/base.tmpl", os.Getenv("GOPATH") + "/templates/index.tpml"}

	templates := template.Must(template.ParseFiles(files...))
	threads, err := data.Threads(); if err == nil {
		templates.ExecuteTemplate(w, "base", threads)
	}
}