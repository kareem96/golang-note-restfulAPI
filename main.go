package main

import (
	"fmt"
	"golang-restful-api-crud/app"
	"golang-restful-api-crud/helper"
	"golang-restful-api-crud/middleware"
	"net/http"
)

func NewServer(authMiddleware *middleware.AuthMiddleware, config app.Config) *http.Server  {
	return &http.Server{
		
		Addr: fmt.Sprintf(":%d", config.AppPort),
		// Addr: "localhost:3000",
		Handler: authMiddleware,
	}
}

func main() {
	server := InitInject()
    fmt.Println("Server setup")
	err := server.ListenAndServe()
    fmt.Println("Server is running")
	helper.PanicIfError(err)
}