package app

import (
	"fmt"
	"golang-restful-api-crud/controller"
	"golang-restful-api-crud/exception"
	"github.com/julienschmidt/httprouter"
)

func NewRouter(noteController controller.NoteController, userController controller.UserController) *httprouter.Router {
    router := httprouter.New()

    // Debug log
    fmt.Println("Setting up routes")

    // Public routes
    router.POST("/api/users", userController.Create)

    // Authenticated routes
    router.GET("/api/notes/:noteId", noteController.FindById)
    router.GET("/api/notes", noteController.FindAll)
    router.POST("/api/notes", noteController.Create)
    router.PUT("/api/notes/:noteId", noteController.Update)
    router.DELETE("/api/notes/:noteId", noteController.Delete)
    
    router.GET("/api/users/:userId/notes", noteController.FindByUserId)
    router.GET("/api/usersDetail", userController.FindAll)
    router.GET("/api/users/:userId", userController.FindById)
    router.PUT("/api/users/:userId", userController.Update)
    router.DELETE("/api/users/:userId", userController.Delete)

    router.PanicHandler = exception.ErrorHandler
    return router
}