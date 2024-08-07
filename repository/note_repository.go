package repository

import (
	"context"
	"golang-restful-api-crud/model/domain"
	"gorm.io/gorm"
)

type NoteRepository interface {
	Save(ctx context.Context, tx *gorm.DB, note domain.Note) domain.Note
	Update(ctx context.Context, tx *gorm.DB, note domain.Note) domain.Note
	Delete(ctx context.Context, tx *gorm.DB, note domain.Note)
	FindById(ctx context.Context, tx *gorm.DB, noteId int) (domain.Note, error)
	FindAll(ctx context.Context, tx *gorm.DB) []domain.Note
}