package service

import (
	"context"
	"golang-restful-api-crud/model/web"
)

type NoteService interface {
	Create(ctx context.Context, request web.NoteCreateRequest) web.NoteResponse
	Update(ctx context.Context, request web.NoteUpdateRequest) web.NoteResponse
	Delete(ctx context.Context, noteId int)
	FindById(ctx context.Context, noteId int) web.NoteResponse
	FindAll(ctx context.Context) []web.NoteResponse
	FindByUserId(ctx context.Context, userId int) []web.NoteResponse
}