package app

import (
	"golang-restful-api-crud/controller"
	"golang-restful-api-crud/exception"
	"github.com/julienschmidt/httprouter"
)

func NewRouter(noteController controller.NoteController)  *httprouter.Router{
	router := httprouter.New()

	//implement router
	router.GET("/api/notes", noteController.FindAll)
	router.GET("/api/notes/:noteId", noteController.FindById)
	router.POST("/api/notes", noteController.Create)
	router.PUT("/api/notes/:noteId", noteController.Update)
	router.DELETE("/api/notes/:noteId", noteController.Delete)

	router.PanicHandler = exception.ErrorHandler
	return router
}