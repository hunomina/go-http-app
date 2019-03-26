package main

import (
	"httpApplication/controller"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", controller.HomeController{}.Hello)
	http.HandleFunc("/css/", controller.ResourcesController{}.GetCssFile)
	log.Fatal(http.ListenAndServe(":8080", nil))
}