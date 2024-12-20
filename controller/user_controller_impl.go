package controller

import (
	"fmt"
	"golang-restful-api-crud/exception"
	"golang-restful-api-crud/helper"
	"golang-restful-api-crud/model/web"
	"golang-restful-api-crud/service"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)


type UserControllerImpl struct {
	UserService service.UserService
}

func NewUserController(userService service.UserService) *UserControllerImpl  {
	return &UserControllerImpl{
		UserService: userService,
	}
}

func (controller UserControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userCreateRequest := web.RegisterUserRequest{}
    helper.ReadFromRequestBody(request, &userCreateRequest)

    fmt.Printf("Create request received: %+v\n", userCreateRequest) // Add this line

    userResponse := controller.UserService.Create(request.Context(), userCreateRequest)
    webResponse := web.WebResponse{
        Code: 200,
        Status: "OK",
        Data: userResponse,
    }

    helper.WriteToResponseBody(writer, webResponse)
}

func (controller UserControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userUpdateRequest := web.UpdateUserRequest{}
	helper.ReadFromRequestBody(request, &userUpdateRequest)

	userId := params.ByName("userId")
	id, err := strconv.Atoi(userId)
	helper.PanicIfError(err)

	userUpdateRequest.ID = id
	userResponse := controller.UserService.Update(request.Context(), userUpdateRequest)
	webResponse := web.WebResponse{
		Code: 200,
		Status: "OK",
		Data: userResponse,
	}
	helper.WriteToResponseBody(writer, webResponse)
}
	
func (controller UserControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userId := params.ByName("userId")
	id, err := strconv.Atoi(userId)
	helper.PanicIfError(err)

	controller.UserService.Delete(request.Context(), id)
	webResponse := web.WebResponse{
		Code: 200,
		Status: "OK",
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller UserControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userId := params.ByName("userId")
	id, err := strconv.Atoi(userId)
	helper.PanicIfError(err)

	userResponse := controller.UserService.FindById(request.Context(), id)
	webResponse := web.WebResponse{
		Code: 200,
		Status: "OK",
		Data: userResponse,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller UserControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userResponses := controller.UserService.FindAll(request.Context())
	webResponse := web.WebResponse{
		Code: 200,
		Status: "OK",
		Data: userResponses,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller UserControllerImpl) Login(writer http.ResponseWriter, request *http.Request, params httprouter.Params)  {
	var loginRequest web.LoginRequest
	helper.ReadFromRequestBody(request, &loginRequest)

	// helper.PanicIfError(err)

	response, err := controller.UserService.Login(request.Context(), loginRequest)
	if err != nil{
		exception.ErrorHandler(writer, request, err)
		return
	}
	webResponse := web.WebResponse{
		Code: http.StatusOK,
		Status: "OK",
		Data: response,
	}
	helper.WriteToResponseBody(writer, webResponse)
}