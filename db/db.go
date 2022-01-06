package db

import (
	"database/sql"
	"log"
	"os"

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

func ConnectDB() *sql.DB {
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
