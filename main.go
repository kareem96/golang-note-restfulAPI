package main

import (
	"golang-restful-api-crud/helper"
	"golang-restful-api-crud/middleware"
	"net/http"
)

func NewServer(authMiddleware *middleware.AuthMiddleware) *http.Server  {
	return &http.Server{
		Addr: "localhost:3000",
		Handler: authMiddleware,
	}
}

func main() {
	server := Initialiaze()
	err := server.ListenAndServe()
	helper.PanicIfError(err)
}