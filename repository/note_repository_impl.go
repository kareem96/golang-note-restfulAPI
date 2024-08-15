package repository

import (
	"context"
	"errors"

	"golang-restful-api-crud/exception"
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
	var user domain.User
	if err := tx.WithContext(ctx).Where("id = ?", note.UserId).First(&user).Error; err != nil{
		helper.PanicIfError(exception.NewNotFoundError("User not found"))
	}
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
	err := tx.WithContext(ctx).Where("id = ?", noteId).First(&note).Error
	// helper.PanicIfError(err)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// return note, exception.HandlerGormError(err, note.Id)

			return note, exception.NewNotFoundError("Note not found")

		}
		return note, err
	}
	
	return note, nil
}

func (repository NoteRepositoryImpl) FindAll(ctx context.Context, tx *gorm.DB) []domain.Note {
	var notes []domain.Note
	err := tx.WithContext(ctx).Find(&notes).Error
	// err := tx.WithContext(ctx).Unscoped().Where("deleted_at IS NOT NULL").Find(&notes).Error
	helper.PanicIfError(err)
	return notes
}
func (repository NoteRepositoryImpl) FindByUserId(ctx context.Context, tx *gorm.DB, userId int) []domain.Note {
	var notes []domain.Note
	err := tx.WithContext(ctx).Where("user_id = ?", userId).Find(&notes).Error
	// err := tx.WithContext(ctx).Unscoped().Where("deleted_at IS NOT NULL").Find(&notes).Error
	helper.PanicIfError(err)
	return notes
}

