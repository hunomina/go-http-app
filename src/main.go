package main

import (
	"controller"
	"log"
	"net/http"
)

func main() {

	// catch every request not caught
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		log.Println(request.Method + " : " + request.RequestURI)
	})

	http.HandleFunc("/hello", controller.HomeController{}.Hello)
	http.HandleFunc("/css/", controller.ResourcesController{}.GetCssFile)
	http.HandleFunc("/users", controller.UserController{}.GetAllUsers)

	log.Fatal(http.ListenAndServe(":8080", nil))
}