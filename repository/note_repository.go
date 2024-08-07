package repository

import (
	"context"
	"database/sql"
	"golang-restful-api-crud/model/domain"
)

type NoteRepository interface {
	Save(ctx context.Context, tx *sql.Tx, note domain.Note) domain.Note
	Update(ctx context.Context, tx *sql.Tx, note domain.Note) domain.Note
	Delete(ctx context.Context, tx *sql.Tx, note domain.Note)
	FindById(ctx context.Context, tx *sql.Tx, noteId int) (domain.Note, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Note
}