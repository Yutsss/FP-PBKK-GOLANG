package repository

import (
	"context"
	"errors"
	"github.com/Yutsss/FP-PBKK-GOLANG/BE/dto"
	"github.com/Yutsss/FP-PBKK-GOLANG/BE/entity"
	"gorm.io/gorm"
)

type (
	UserRepository interface {
		Create(ctx context.Context, tx *gorm.DB, data dto.UserRegisterRequest) (entity.User, error)
		FindByEmail(ctx context.Context, email string) (entity.User, error)
		FIndById(ctx context.Context, id uint) (entity.User, error)
	}

	userRepository struct {
		db *gorm.DB
	}
)

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) Create(ctx context.Context, tx *gorm.DB, data dto.UserRegisterRequest) (entity.User, error) {
	if tx == nil {
		tx = r.db
	}

	user := entity.User{
		Name:         data.Name,
		CompleteName: data.CompleteName,
		Email:        data.Email,
		Password:     data.Password,
		PhoneNumber:  data.PhoneNumber,
		Address:      data.Address,
	}

	err := tx.WithContext(ctx).Create(&user).Error

	if err != nil {
		return entity.User{}, err
	}

	return user, nil
}

func (r *userRepository) FindByEmail(ctx context.Context, email string) (entity.User, error) {
	var user entity.User

	err := r.db.WithContext(ctx).Where("email = ?", email).First(&user).Error

	if err != nil {
		if errors.As(err, &gorm.ErrRecordNotFound) {
			return entity.User{}, nil
		}
		return entity.User{}, err
	}

	return user, nil
}

func (r *userRepository) FIndById(ctx context.Context, id uint) (entity.User, error) {
	var user entity.User

	err := r.db.WithContext(ctx).Where("id = ?", id).First(&user).Error

	if err != nil {
		if errors.As(err, &gorm.ErrRecordNotFound) {
			return entity.User{}, nil
		}
		return entity.User{}, err
	}

	return user, nil
}
