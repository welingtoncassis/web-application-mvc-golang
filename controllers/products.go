package controllers

import (
	"net/http"
	"store/models"
	"text/template"
)

var temps = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	allProducts := models.GetAll()
	temps.ExecuteTemplate(w, "Index", allProducts)
}
