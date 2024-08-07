package app

import (
	"golang-restful-api-crud/exception"
	"github.com/julienschmidt/httprouter"
)

func NewRouter()  *httprouter.Router{
	router := httprouter.New()

	//implement router
	router.GET("", nil)
	router.GET("", nil)
	router.POST("", nil)
	router.PUT("", nil)
	router.DELETE("", nil)

	router.PanicHandler = exception.ErrorHandler
	return router
}