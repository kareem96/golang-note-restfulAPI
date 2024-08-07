package controller

import (
	"golang-restful-api-crud/helper"
	"golang-restful-api-crud/model/web"
	"golang-restful-api-crud/service"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)


type NoteControllerImpl struct {
	NoteService service.NoteService
}

func NewNoteController(noteService service.NoteService) *NoteControllerImpl{
	return &NoteControllerImpl{
		NoteService: noteService,
	}
}

func (controller NoteControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	noteCreateRequest := web.NoteCreateRequest{}
	helper.ReadFromRequestBody(request, &noteCreateRequest)

	noteResponse := controller.NoteService.Create(request.Context(), noteCreateRequest)
	webResponse := web.WebResponse{
		Code: 200,
		Status: "OK",
		Data: noteResponse,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller NoteControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	noteUpdateRequest := web.NoteUpdateRequest{}
	helper.ReadFromRequestBody(request, &noteUpdateRequest)

	noteId := params.ByName("noteId")
	id, err := strconv.Atoi(noteId)
	helper.PanicIfError(err)
	
	noteUpdateRequest.Id = id
	noteResponse := controller.NoteService.Update(request.Context(), noteUpdateRequest)
	webResponse := web.WebResponse{
		Code: 200,
		Status: "OK",
		Data: noteResponse,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller NoteControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	noteId := params.ByName("noteId")
	id, err := strconv.Atoi(noteId)
	helper.PanicIfError(err)

	controller.NoteService.Delete(request.Context(), id)
	webResponse := web.WebResponse{
		Code: 200,
		Status: "OK",
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller NoteControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	noteId := params.ByName("noteId")
	id, err := strconv.Atoi(noteId)
	helper.PanicIfError(err)

	noteResponse := controller.NoteService.FindById(request.Context(), id)
	webResponse := web.WebResponse{
		Code: 200,
		Status: "OK",
		Data: noteResponse,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller NoteControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, param httprouter.Params) {
	noteResponses := controller.NoteService.FindAll(request.Context())
	webResponse := web.WebResponse{
		Code: 200,
		Status: "OK",
		Data: noteResponses,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

