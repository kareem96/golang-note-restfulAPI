// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
	"golang-restful-api-crud/app"
	"golang-restful-api-crud/controller"
	"golang-restful-api-crud/middleware"
	"golang-restful-api-crud/repository"
	"golang-restful-api-crud/service"
	"net/http"
)

// Injectors from injector.go:

func InitInject() *http.Server {
	noteRepositoryImpl := repository.NewNoteRepository()
	config := ProvideConfig()
	db := app.OpenConnectionDB(config)
	validate := ProvideValidator()
	noteServiceImpl := service.NewNoteService(noteRepositoryImpl, db, validate)
	noteControllerImpl := controller.NewNoteController(noteServiceImpl)
	userRepositoryImpl := repository.NewUserRepository()
	userServiceImpl := service.NewUserService(userRepositoryImpl, db, validate)
	userControllerImpl := controller.NewUserController(userServiceImpl)
	router := app.NewRouter(noteControllerImpl, userControllerImpl)
	v := ProvideExcludedRoutes()
	authMiddleware := middleware.NewAuthMiddleware(router, v)
	server := NewServer(authMiddleware, config)
	return server
}

// injector.go:

var noteSet = wire.NewSet(repository.NewNoteRepository, wire.Bind(new(repository.NoteRepository), new(*repository.NoteRepositoryImpl)), service.NewNoteService, wire.Bind(new(service.NoteService), new(*service.NoteServiceImpl)), controller.NewNoteController, wire.Bind(new(controller.NoteController), new(*controller.NoteControllerImpl)))

var userSet = wire.NewSet(repository.NewUserRepository, wire.Bind(new(repository.UserRepository), new(*repository.UserRepositoryImpl)), service.NewUserService, wire.Bind(new(service.UserService), new(*service.UserServiceImpl)), controller.NewUserController, wire.Bind(new(controller.UserController), new(*controller.UserControllerImpl)))

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