package repository

import (
	"context"
	"errors"
	"github.com/Yutsss/FP-PBKK-GOLANG/BE/dto"
	"github.com/Yutsss/FP-PBKK-GOLANG/BE/entity"
	errorUtils "github.com/Yutsss/FP-PBKK-GOLANG/BE/utility/error"
	"gorm.io/gorm"
)

type (
	UserRepository interface {
		Create(ctx context.Context, tx *gorm.DB, data dto.UserRegisterRequest) (entity.User, errorUtils.CustomError)
		FindByEmail(ctx context.Context, tx *gorm.DB, email string) (entity.User, errorUtils.CustomError)
		FindById(ctx context.Context, tx *gorm.DB, id int64) (entity.User, errorUtils.CustomError)
		FindAll(ctx context.Context, tx *gorm.DB) ([]entity.User, errorUtils.CustomError)
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

func (r *userRepository) Create(ctx context.Context, tx *gorm.DB, data dto.UserRegisterRequest) (entity.User, errorUtils.CustomError) {
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
		return entity.User{}, errorUtils.ErrInternalServer
	}

	return user, nil
}

func (r *userRepository) FindByEmail(ctx context.Context, tx *gorm.DB, email string) (entity.User, errorUtils.CustomError) {
	if tx == nil {
		tx = r.db
	}

	var user entity.User

	err := tx.WithContext(ctx).Where("email = ?", email).First(&user).Error

	if err != nil {
		if errors.As(err, &gorm.ErrRecordNotFound) {
			return entity.User{}, nil
		}
		return entity.User{}, errorUtils.ErrInternalServer
	}

	return user, nil
}

func (r *userRepository) FindById(ctx context.Context, tx *gorm.DB, id int64) (entity.User, errorUtils.CustomError) {
	if tx == nil {
		tx = r.db
	}

	var user entity.User

	err := tx.WithContext(ctx).Where("id = ?", id).First(&user).Error

	if err != nil {
		if errors.As(err, &gorm.ErrRecordNotFound) {
			return entity.User{}, nil
		}
		return entity.User{}, errorUtils.ErrInternalServer
	}

	return user, nil
}

func (r *userRepository) FindAll(ctx context.Context, tx *gorm.DB) ([]entity.User, errorUtils.CustomError) {
	if tx == nil {
		tx = r.db
	}

	var users []entity.User

	err := tx.WithContext(ctx).Where("role = ?", "user").Find(&users).Error

	if err != nil {
		if errors.As(err, &gorm.ErrRecordNotFound) {
			return []entity.User{}, nil
		}
		return []entity.User{}, errorUtils.ErrInternalServer
	}

	return users, nil
}
