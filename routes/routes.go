package routes

import (
	"net/http"
	"store/controllers"
)

func GetRoutes() {
	http.HandleFunc("/", controllers.Index)
}
