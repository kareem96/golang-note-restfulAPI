package app

import (
	"fmt"
	"golang-restful-api-crud/controller"
	"golang-restful-api-crud/exception"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func NewRouter(noteController controller.NoteController, userController controller.UserController) *httprouter.Router {
    router := httprouter.New()

    // Debug log
    fmt.Println("Setting up routes")

    // Public routes
    router.POST("/api/users", userController.Create)
    
    //logn public router
    router.POST("/api/users/login", userController.Login)

    // Serve Swagger UI static files
    router.ServeFiles("/swagger-ui/*filepath", http.Dir("swagger-ui"))

    // Serve Swagger UI index file
    router.GET("/api/apispec", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
        http.ServeFile(w, r, "./swagger-ui/index.html")
    })

    // Serve the OpenAPI spec
    router.GET("/api/apispec.json", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
        http.ServeFile(w, r, ".apispec.json")
    })

    router.GET("/api/apispecs", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
        http.ServeFile(w, r, ".test.html")
    })


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