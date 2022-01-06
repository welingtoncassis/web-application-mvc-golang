package main

import (
	"net/http"
	"store/routes"
)

func main() {
	routes.GetRoutes()
	http.ListenAndServe(":8000", nil)
}
