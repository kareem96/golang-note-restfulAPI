package service

import (
	"context"
	"golang-restful-api-crud/model/web"
)


type UserService interface {
	Create(ctx context.Context, request web.RegisterUserRequest) web.UserResponse
	Update(ctx context.Context, request web.UpdateUserRequest) web.UserResponse
	Delete(ctx context.Context, userId int)
	FindById(ctx context.Context, userId int) web.UserResponse
	FindAll(ctx context.Context) []web.UserResponse
	
	Login(ctx context.Context, request web.LoginRequest) web.UserResponse
}