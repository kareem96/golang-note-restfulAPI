package repository

import (
	"context"
	"golang-restful-api-crud/helper"
	"golang-restful-api-crud/model/domain"
	"gorm.io/gorm"
)

type NoteRepositoryImpl struct {
	
}
func NewNoteRepository() *NoteRepositoryImpl {
	return &NoteRepositoryImpl{}
}

func (repository NoteRepositoryImpl) Save(ctx context.Context, tx *gorm.DB, note domain.Note) domain.Note {
	err := tx.WithContext(ctx).Create(&note).Error
	helper.PanicIfError(err)
	return note
}

func (repository NoteRepositoryImpl) Update(ctx context.Context, tx *gorm.DB, note domain.Note) domain.Note {
	err := tx.WithContext(ctx).Save(&note).Error
	helper.PanicIfError(err)
	return note
}

func (repository NoteRepositoryImpl) Delete(ctx context.Context, tx *gorm.DB, note domain.Note) {
	err := tx.WithContext(ctx).Delete(&note).Error
	helper.PanicIfError(err)
	
}

func (repository NoteRepositoryImpl) FindById(ctx context.Context, tx *gorm.DB, noteId int) (domain.Note, error) {
	var note domain.Note
	err := tx.WithContext(ctx).First(&note, noteId).Error
	helper.PanicIfError(err)
	return note, nil
}

func (repository NoteRepositoryImpl) FindAll(ctx context.Context, tx *gorm.DB) []domain.Note {
	var notes []domain.Note
	err := tx.WithContext(ctx).Find(&notes).Error
	helper.PanicIfError(err)
	return notes
}

