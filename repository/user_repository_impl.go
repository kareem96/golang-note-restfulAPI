package repository

import (
	"context"
	"errors"
	"fmt"
	"golang-restful-api-crud/exception"
	"golang-restful-api-crud/helper"
	"golang-restful-api-crud/model/domain"

	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	
}

func NewUserRepository() *UserRepositoryImpl {
	return &UserRepositoryImpl{}
}

func (repository UserRepositoryImpl) Save(ctx context.Context, tx *gorm.DB, user domain.User) domain.User {
	err := tx.WithContext(ctx).Create(&user).Error
	helper.PanicIfError(err)
	return user
}

func (repository UserRepositoryImpl) Update(ctx context.Context, tx *gorm.DB, user domain.User) domain.User {
	err := tx.WithContext(ctx).Save(&user).Error
	helper.PanicIfError(err)
	return user
}

func (repository UserRepositoryImpl) Delete(ctx context.Context, tx *gorm.DB, user domain.User) {
	err := tx.WithContext(ctx).Delete(&user).Error
	helper.PanicIfError(err)
}

func (repository UserRepositoryImpl) FindById(ctx context.Context, tx *gorm.DB, userId int) (domain.User, error) {
	var user domain.User
	err := tx.WithContext(ctx).Where("id = ?", userId).First(&user).Error
	helper.PanicIfError(err)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// return note, exception.HandlerGormError(err, note.Id)

			return user, exception.NewNotFoundError("Note not found")

		}
		return user, err
	}
	return user, nil
}

func (repository UserRepositoryImpl) FindAll(ctx context.Context, tx *gorm.DB) []domain.User {
	var users []domain.User
	err := tx.WithContext(ctx).Find(&users).Error
	helper.PanicIfError(err)
	return users
}

func (repository UserRepositoryImpl) Login(ctx context.Context, tx *gorm.DB, username string) (domain.User, error) {
	var user domain.User
	err := tx.WithContext(ctx).Where("name = ?", username).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return user, fmt.Errorf("user not found: %w", err)
		}
		return user, fmt.Errorf("error retrieving user: %w", err)
	}
	return user, nil
}

