package main

import (
	"net/http"
	"store/models"
	"text/template"
)

var temps = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	allProducts := models.GetAll()
	temps.ExecuteTemplate(w, "Index", allProducts)
}
