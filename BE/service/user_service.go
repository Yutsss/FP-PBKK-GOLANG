package service

import (
	"context"
	"github.com/Yutsss/FP-PBKK-GOLANG/BE/dto"
	"github.com/Yutsss/FP-PBKK-GOLANG/BE/repository"
	"github.com/Yutsss/FP-PBKK-GOLANG/BE/utility"
	errorUtils "github.com/Yutsss/FP-PBKK-GOLANG/BE/utility/error"
	"github.com/Yutsss/FP-PBKK-GOLANG/BE/validation"
)

type (
	UserService interface {
		Register(ctx context.Context, data dto.UserRegisterRequest) (dto.UserRegisterResponse, errorUtils.CustomError)
		Login(ctx context.Context, data dto.UserLoginRequest) (dto.UserLoginResponse, errorUtils.CustomError)
		GetById(ctx context.Context, data dto.UserGetByIdRequest) (dto.UserGetByIdResponse, errorUtils.CustomError)
		GetAll(ctx context.Context) (dto.UserGetAllResponse, errorUtils.CustomError)
	}

	userService struct {
		userRepo repository.UserRepository
		jwtUtils utility.JWTUtils
	}
)

func NewUserService(userRepo repository.UserRepository, jwtUtils utility.JWTUtils) UserService {
	return &userService{
		userRepo: userRepo,
		jwtUtils: jwtUtils,
	}
}

func (s *userService) Register(ctx context.Context, data dto.UserRegisterRequest) (dto.UserRegisterResponse, errorUtils.CustomError) {
	if err := validation.Validate(data); err != nil {
		return dto.UserRegisterResponse{}, err
	}

	userExist, err := s.userRepo.FindByEmail(ctx, nil, data.Email)
	if err != nil {
		return dto.UserRegisterResponse{}, errorUtils.ErrInternalServer
	}

	if userExist.ID != 0 {
		return dto.UserRegisterResponse{}, errorUtils.ErrEmailAlreadyExist
	}

	user, err := s.userRepo.Create(ctx, nil, data)
	if err != nil {
		return dto.UserRegisterResponse{}, errorUtils.ErrInternalServer
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

func (s *userService) Login(ctx context.Context, data dto.UserLoginRequest) (dto.UserLoginResponse, errorUtils.CustomError) {
	if err := validation.Validate(data); err != nil {
		return dto.UserLoginResponse{}, err
	}

	user, err := s.userRepo.FindByEmail(ctx, nil, data.Email)
	if err != nil {
		return dto.UserLoginResponse{}, errorUtils.ErrInternalServer
	}

	if user.ID == 0 {
		return dto.UserLoginResponse{}, errorUtils.ErrWrongEmailOrPassword
	}

	if isPasswordMatch := utility.CheckPassword(user.Password, data.Password); !isPasswordMatch {
		return dto.UserLoginResponse{}, errorUtils.ErrWrongEmailOrPassword
	}

	AccessToken, err := s.jwtUtils.GenerateToken(user.ID, user.Role)

	if err != nil {
		return dto.UserLoginResponse{}, errorUtils.ErrInternalServer
	}

	return dto.UserLoginResponse{
		AccessToken: AccessToken,
		Role:        user.Role,
	}, nil
}

func (s *userService) GetById(ctx context.Context, data dto.UserGetByIdRequest) (dto.UserGetByIdResponse, errorUtils.CustomError) {
	if err := validation.Validate(data); err != nil {
		return dto.UserGetByIdResponse{}, err
	}

	user, err := s.userRepo.FindById(ctx, nil, data.UserID)
	if err != nil {
		return dto.UserGetByIdResponse{}, errorUtils.ErrInternalServer
	}

	if user.ID == 0 {
		return dto.UserGetByIdResponse{}, errorUtils.ErrUserNotFound
	}

	res := dto.UserGetByIdResponse{
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

func (s *userService) GetAll(ctx context.Context) (dto.UserGetAllResponse, errorUtils.CustomError) {
	users, err := s.userRepo.FindAll(ctx, nil)
	if err != nil {
		return dto.UserGetAllResponse{}, errorUtils.ErrInternalServer
	}

	res := dto.UserGetAllResponse{
		Users: make([]dto.UserGetByIdResponse, 0),
	}

	for _, user := range users {
		res.Users = append(res.Users, dto.UserGetByIdResponse{
			ID:           user.ID,
			Name:         user.Name,
			CompleteName: user.CompleteName,
			Email:        user.Email,
			PhoneNumber:  user.PhoneNumber,
			Address:      user.Address,
			Role:         user.Role,
		})
	}

	return res, nil
}
