package repository

import (
	"context"
	"errors"
	"imp-backend/application/infrastructure"
	"imp-backend/application/misc"
	"imp-backend/domain"

	"gorm.io/gorm"
)

type AuthRepository struct {
	DB *gorm.DB
}

func NewAuthRepository(db *gorm.DB) infrastructure.AuthRepository {
	return &AuthRepository{
		DB: db,
	}
}

func (r *AuthRepository) GetUser(ctx context.Context, params *domain.Users) (*domain.Users, error) {
	db := r.DB
	if params.ID > 0 {
		db = db.Where("id = ?", params.ID)
	}
	if params.Username != "" {
		db = db.Where("username = ?", params.Username)
	}

	if err := db.First(&params).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		misc.LogEf(`AuthRepository-GetUser Error %s`, err.Error())
		return nil, err
	}
	return params, nil
}

func (r *AuthRepository) CreateUser(ctx context.Context, params *domain.Users) (userID int64, err error) {
	err = r.DB.Create(&params).Error
	if err != nil {
		misc.LogEf(`AuthRepository-CreateUser Error %s`, err.Error())
		return userID, err
	}
	userID = params.ID
	return userID, nil
}

func (r *AuthRepository) UpdateUser(ctx context.Context, params *domain.Users) error {
	err := r.DB.Save(&params).Error
	if err != nil {
		misc.LogEf(`AuthRepository-UpdateUser Error %s`, err.Error())
		return err
	}
	return nil
}

func (r *AuthRepository) SelectUser(ctx context.Context, params *domain.UserParams) ([]domain.Users, error) {
	res := []domain.Users{}

	db := r.DB

	if params.Page > 0 && params.Limit > 0 {
		offset := (params.Page - 1) * params.Limit
		db = db.Offset(offset).Limit(params.Limit)
	}

	if err := db.Order("created_at DESC").Find(&res).Error; err != nil {
		misc.LogEf(`AuthRepository-SelectUser Error %s`, err.Error())
		return nil, err
	}
	return res, nil
}
