package repository

import (
	"context"
	"database/sql"
	"golang-restful-api-crud/model/domain"
)

type NoteRepositoryImpl struct {
	
}

func NewNoteRepository() *NoteRepositoryImpl {
	return &NoteRepositoryImpl{}
}

func (repository NoteRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, note domain.Note) domain.Note {
	panic("not implemented") // TODO: Implement
}

func (repository NoteRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, note domain.Note) domain.Note {
	panic("not implemented") // TODO: Implement
}

func (repository NoteRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, note domain.Note) {
	panic("not implemented") // TODO: Implement
}

func (repository NoteRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, noteId int) (domain.Note, error) {
	panic("not implemented") // TODO: Implement
}

func (repository NoteRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Note {
	panic("not implemented") // TODO: Implement
}

