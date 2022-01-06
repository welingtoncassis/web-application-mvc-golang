package routes

import (
	"net/http"
	"store/controllers"
)

func GetRoutes() {
	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/new", controllers.New)
	http.HandleFunc("/create", controllers.Create)
}
