package service

import (
	"context"
	"golang-restful-api-crud/helper"
	"golang-restful-api-crud/model/domain"
	"golang-restful-api-crud/model/web"
	"golang-restful-api-crud/repository"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)
type NoteServiceImpl struct {
	NoteRepository repository.NoteRepository
	DB *gorm.DB
	Validate *validator.Validate
}

func NewNoteService(noteRepository repository.NoteRepository, DB *gorm.DB, validate * validator.Validate) *NoteServiceImpl {
	return &NoteServiceImpl{
		NoteRepository: noteRepository,
		DB: DB,
		Validate: validate,
	}	
}

func (service NoteServiceImpl) Create(ctx context.Context, request web.NoteCreateRequest) web.NoteResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx := service.DB.Begin()
	defer func ()  {
		if r := recover(); r != nil {
			tx.Rollback()
			panic(r)
		}
	}()

	note := domain.Note{
		Title: request.Title,
		Description: request.Description,
	}
	
	note = service.NoteRepository.Save(ctx, tx, note)
	tx.Commit()
	return web.NoteResponse{
		ID: note.Id,
		Title: note.Title,
		Description: note.Description,
		CreatedAt: note.CreatedAt.String(),
		UpdatedAt: note.UpdatedAt.String(),
	}
}

func (service NoteServiceImpl) Update(ctx context.Context, request web.NoteCreateRequest) web.NoteResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx := service.DB.Begin()
	defer func ()  {
		if r := recover(); r != nil {
			tx.Rollback()
			panic(r)
		}
	}()

	note, err := service.NoteRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		tx.Rollback()
		panic(err)
	}
	note.Title = request.Title
	note.Description = request.Description

	note = service.NoteRepository.Update(ctx, tx, note)
	tx.Commit()

	return web.NoteResponse{
		ID: note.Id,
		Title: note.Title,
		Description: note.Description,
	}
}

func (service NoteServiceImpl) Delete(ctx context.Context, noteId int) {
	tx := service.DB.Begin()
	defer func ()  {
		if r := recover(); r != nil {
			tx.Rollback()
			panic(r)
		}	
	}()

	note, err := service.NoteRepository.FindById(ctx, tx, noteId)
	if err != nil {
		tx.Rollback()
		panic(err)
	}

	service.NoteRepository.Delete(ctx, tx, note)
	tx.Commit()
}

func (service NoteServiceImpl) FindById(ctx context.Context, noteId int) web.NoteResponse {
	tx := service.DB.Begin()
	defer func () {
		tx.Rollback()
	}()

	note, err := service.NoteRepository.FindById(ctx, tx, noteId)
	if err != nil {
		panic(err)
	}
	tx.Commit()

	return web.NoteResponse{
		ID: note.Id,
		Title: note.Title,
		Description: note.Description,
	}
}

func (service NoteServiceImpl) FindAll(ctx context.Context) []web.NoteResponse {
	tx := service.DB.Begin()
	defer func ()  {
		tx.Rollback()
	}()

	notes := service.NoteRepository.FindAll(ctx, tx)
	tx.Commit()

	var noteResponses []web.NoteResponse
	for _, note := range notes {
		noteResponses = append(noteResponses, web.NoteResponse{
			ID: note.Id,
			Title: note.Title,
			Description: note.Description,
		})
	}
	
	return noteResponses
}

