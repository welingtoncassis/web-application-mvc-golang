package main

import (
	"net/http"
	"text/template"
)

type Product struct {
	Name        string
	Description string
	Price       float64
	Amount      int
}

var temps = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	products := []Product{
		{"T-Shirt", "Black", 29.9, 10},
		{"Short", "Blue", 19.9, 5},
		{"Shoes", "Red", 89.9, 3},
	}
	// Parms: w(quem responde), index(html), valor passado para o html
	temps.ExecuteTemplate(w, "Index", products)
}
