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

func ProvideValidator() *validator.Validate {
	return validator.New()
}

func Initialiaze() *http.Server {
	wire.Build(
		app.OpenConnectionDB,
		ProvideValidator,
		noteSet,
		app.NewRouter,
		wire.Bind(new(http.Handler), new(*httprouter.Router)),
		middleware.NewAuthMiddleware,
		NewServer,
	)
	return nil
}