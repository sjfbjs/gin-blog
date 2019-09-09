package main

import (
	"gin-blog/controller"
	"gin-blog/model"
	"log"
	"net/http"
)

func init() {
	model.Load()
}

func main() {
	router := controller.MapRoutes()
	server := &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: router,
	}
	if err := server.ListenAndServe(); nil != err {
		log.Fatal("listen and serve failed: " + err.Error())
	}
}
