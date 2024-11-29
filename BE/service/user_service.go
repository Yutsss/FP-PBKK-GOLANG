package service

import (
	"context"
	"github.com/Yutsss/FP-PBKK-GOLANG/BE/dto"
	"github.com/Yutsss/FP-PBKK-GOLANG/BE/repository"
	errorUtils "github.com/Yutsss/FP-PBKK-GOLANG/BE/utility/error"
	"github.com/Yutsss/FP-PBKK-GOLANG/BE/validation"
)

type (
	UserService interface {
		Register(ctx context.Context, data dto.UserRegisterRequest) (dto.UserRegisterResponse, error)
	}

	userService struct {
		userRepo repository.UserRepository
	}
)

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{
		userRepo: userRepo,
	}
}

func (s *userService) Register(ctx context.Context, data dto.UserRegisterRequest) (dto.UserRegisterResponse, error) {
	if err := validation.Validate(data); err != nil {
		return dto.UserRegisterResponse{}, err
	}

	userExist, err := s.userRepo.FindByEmail(ctx, data.Email)
	if err != nil {
		return dto.UserRegisterResponse{}, err
	}

	if userExist.ID != 0 {
		return dto.UserRegisterResponse{}, errorUtils.ErrEmailAlreadyExist
	}

	user, err := s.userRepo.Create(ctx, nil, data)
	if err != nil {
		return dto.UserRegisterResponse{}, err
	}

	res := dto.UserRegisterResponse{
		ID:           user.ID,
		Name:         user.Name,
		CompleteName: user.CompleteName,
		Email:        user.Email,
		PhoneNumber:  user.PhoneNumber,
		Address:      user.Address,
		Role:         user.Role,
	}

	return res, nil
}
