package main

import (
	"net/http"
	"text/template"
)

var temps = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	// Parms: w(quem responde), index(html), valor passado para o html
	temps.ExecuteTemplate(w, "Index", nil)
}
