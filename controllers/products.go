package controllers

import (
	"log"
	"net/http"
	"store/models"
	"strconv"
	"text/template"
)

var temps = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	allProducts := models.GetAll()
	temps.ExecuteTemplate(w, "Index", allProducts)
}

func New(w http.ResponseWriter, r *http.Request) {
	temps.ExecuteTemplate(w, "New", nil)
}

func Create(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		name := r.FormValue("nome")
		description := r.FormValue("descricao")
		price := r.FormValue("preco")
		amount := r.FormValue("quantidade")

		convertedPrice, err := strconv.ParseFloat(price, 64)
		if err != nil {
			log.Println("Erro na conversão do preço: ", err)
		}

		convertedAmount, err := strconv.Atoi(amount)
		if err != nil {
			log.Println("Erro na conversão da quantidade: ", err)
		}

		models.Create(name, description, convertedPrice, convertedAmount)
	}
	http.Redirect(w, r, "/", 301)
}
