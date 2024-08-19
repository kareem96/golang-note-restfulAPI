package service

import (
	"context"
	"golang-restful-api-crud/helper"
	"golang-restful-api-crud/model/domain"
	"golang-restful-api-crud/model/web"
	"golang-restful-api-crud/repository"
	"log"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)
type UserServiceImpl struct {
	UserRepository repository.UserRepository
	DB *gorm.DB
	Validate *validator.Validate
}

func NewUserService(userRepository repository.UserRepository, DB *gorm.DB, validate *validator.Validate) *UserServiceImpl  {
	return &UserServiceImpl{
		UserRepository: userRepository,
		DB: DB,
		Validate: validate,
	}	
}

func (service UserServiceImpl) Create(ctx context.Context, request web.RegisterUserRequest) web.UserResponse {
    err := service.Validate.Struct(request)
    helper.PanicIfError(err)
	tx := service.DB.Begin()
    defer func() {
        if r := recover(); r != nil {
            tx.Rollback()
        }
    }()

	user := domain.User{
        Password: helper.HashPassword(request.Password),
        Name: request.Name,
		
    }
	user = service.UserRepository.Save(ctx, tx, user)

    token, err := helper.GenerateJWT(strconv.Itoa(user.ID), 2 * time.Minute)
    if err != nil {
        log.Printf("JWT generation error: %v", err)
        return web.UserResponse{}
    }
	tx.Commit()
    return web.UserResponse{
        ID: user.ID,
        Name: user.Name,
        Token: token,
        CreatedAt: user.CreatedAt,
        UpdatedAt: user.UpdatedAt,
    }
}

func (service UserServiceImpl) Update(ctx context.Context, request web.UpdateUserRequest) web.UserResponse {
	err := service.Validate.Struct(request)
	if err != nil{
		panic(err)
	}

	tx := service.DB.Begin()
	defer tx.Commit()

	user, err := service.UserRepository.FindById(ctx, tx, request.ID)
	if err != nil{
		panic(err)
	}
	if request.Name != ""{
		user.Name = request.Name
	}
	if request.Password != ""{
		user.Name = helper.HashPassword(request.Password)
	}

	user = service.UserRepository.Update(ctx, tx, user)
	return web.UserResponse{
		ID: user.ID,
		Name: user.Name,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func (service UserServiceImpl) Delete(ctx context.Context, userId int) {
	tx := service.DB.Begin()
	defer tx.Commit()
	
	user, err :=service.UserRepository.FindById(ctx, tx, userId)
	if err != nil{
		panic(err)
	}
	service.UserRepository.Delete(ctx, tx, user)
}

func (service UserServiceImpl) FindById(ctx context.Context, userId int) web.UserResponse {
	tx := service.DB.Begin()
	defer tx.Commit()

	user, err := service.UserRepository.FindById(ctx, tx, userId)
	if err != nil{
		panic(err)
	}
	return web.UserResponse{
		ID: user.ID,
		Name: user.Name,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func (service UserServiceImpl) FindAll(ctx context.Context) []web.UserResponse {
	tx := service.DB.Begin()
	defer tx.Commit()

	users := service.UserRepository.FindAll(ctx, tx)

	var userResponses []web.UserResponse
	for _, user := range users {
		userResponses = append(userResponses, web.UserResponse{
			ID: user.ID,
			Name: user.Name,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		})
	}
	
	return userResponses
}


func (service UserServiceImpl) Login(ctx context.Context, request web.LoginRequest) web.UserResponse {
    tx := service.DB.Begin()
    defer func() {
        if r := recover(); r != nil {
            tx.Rollback()
            log.Printf("Recovered from panic: %v", r)
        } else {
            tx.Commit()
        }
    }()

    user, err := service.UserRepository.Login(ctx, tx, request.Name)
    if err != nil {
        log.Printf("Error finding user: %v", err)
        return web.UserResponse{}
    }

    if !helper.CheckPasswordHash(request.Password, user.Password) {
        log.Printf("Invalid credentials for user: %s", request.Name)
        return web.UserResponse{}
    }

    token, err := helper.GenerateJWT(strconv.Itoa(user.ID), 2*time.Minute)
    if err != nil {
        log.Printf("JWT generation error: %v", err)
        return web.UserResponse{}
    }

    return web.UserResponse{
        ID:        user.ID,
        Name:      user.Name,
        Token:     token,
        CreatedAt: user.CreatedAt,
        UpdatedAt: user.UpdatedAt,
    }
}