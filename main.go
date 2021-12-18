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

	connectionString := "user=" + user + " dbname=" + dbName + " password=" + password + " host=" + host + "sslmode" + sslmode

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		panic(err.Error())
	}
	return db
}

type Product struct {
	Name        string
	Description string
	Price       float64
	Amount      int
}

var temps = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	db := connectDB()
	defer db.Close()
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
