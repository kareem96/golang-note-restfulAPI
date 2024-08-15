//go:build wireinject
// +build wireinject

package main

import (
	"golang-restful-api-crud/app"
	"golang-restful-api-crud/controller"
	"golang-restful-api-crud/middleware"
	"golang-restful-api-crud/repository"
	"golang-restful-api-crud/service"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
	"github.com/julienschmidt/httprouter"
)

var noteSet = wire.NewSet(
	repository.NewNoteRepository,
	wire.Bind(new(repository.NoteRepository), new(*repository.NoteRepositoryImpl)),
	
	service.NewNoteService,
	wire.Bind(new(service.NoteService), new(*service.NoteServiceImpl)),

	controller.NewNoteController,
	wire.Bind(new(controller.NoteController), new(*controller.NoteControllerImpl)),
)

var userSet = wire.NewSet(
	repository.NewUserRepository,
	wire.Bind(new(repository.UserRepository), new(*repository.UserRepositoryImpl)),
	
	service.NewUserService,
	wire.Bind(new(service.UserService), new(*service.UserServiceImpl)),

	controller.NewUserController,
	wire.Bind(new(controller.UserController), new(*controller.UserControllerImpl)),
)

func ProvideValidator() *validator.Validate {
	return validator.New()
}

func ProvideConfig() app.Config {
	return app.NewViper()
}

func ProvideExcludedRoutes() []string {
	return []string{
		"/api/users",
	}
}

func InitInject() *http.Server {
	wire.Build(
		ProvideConfig,
		app.OpenConnectionDB,
		ProvideValidator,
		ProvideExcludedRoutes,
		noteSet,
		userSet,
		wire.Bind(new(http.Handler), new(*httprouter.Router)),
		app.NewRouter,
		middleware.NewAuthMiddleware,
		NewServer,
	)
	return nil
}
