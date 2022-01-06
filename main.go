package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"text/template"

	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

func goDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

func connectDB() *sql.DB {
	user := goDotEnvVariable("DB_USER")
	password := goDotEnvVariable("DB_PASSWORD")
	dbName := goDotEnvVariable("DB_NAME")
	host := goDotEnvVariable("DB_HOST")
	sslmode := goDotEnvVariable("DB_SSLMODE")

	connectionString := "user=" + user + " dbname=" + dbName + " password=" + password + " host=" + host + " sslmode=" + sslmode

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		panic(err.Error())
	}
	return db
}

type Product struct {
	Id          int
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
	db := connectDB()
	resultQuery, err := db.Query("select * from products")
	if err != nil {
		panic(err.Error())
	}

	p := Product{}
	products := []Product{}

	for resultQuery.Next() {
		var id, amount int
		var name, description string
		var price float64

		err = resultQuery.Scan(&id, &name, &description, &price, &amount)
		if err != nil {
			panic(err.Error())
		}

		p.Id = id
		p.Name = name
		p.Description = description
		p.Price = price
		p.Amount = amount

		products = append(products, p)
	}

	// Parms: w(quem responde), index(html), valor passado para o html
	temps.ExecuteTemplate(w, "Index", products)
	defer db.Close()
}
