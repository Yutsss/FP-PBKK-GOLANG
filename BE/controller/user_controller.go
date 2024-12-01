package controller

import (
	"fmt"
	"github.com/Yutsss/FP-PBKK-GOLANG/BE/dto"
	"github.com/Yutsss/FP-PBKK-GOLANG/BE/service"
	"github.com/Yutsss/FP-PBKK-GOLANG/BE/utility"
	errorUtils "github.com/Yutsss/FP-PBKK-GOLANG/BE/utility/error"
	successUtils "github.com/Yutsss/FP-PBKK-GOLANG/BE/utility/success"
	"github.com/gin-gonic/gin"
	"net/http"
)

type (
	UserController interface {
		Register(ctx *gin.Context)
		Login(ctx *gin.Context)
	}

	userController struct {
		userService service.UserService
	}
)

func NewUserController(us service.UserService) UserController {
	return &userController{
		userService: us,
	}
}

func (c *userController) Register(ctx *gin.Context) {
	var req dto.UserRegisterRequest
	if err := ctx.ShouldBind(&req); err != nil {
		res := utility.ResponseError(errorUtils.MESSAGE_FAILED_TO_GET_DATA_FROM_BODY, errorUtils.ErrBadRequest.Error(), errorUtils.ErrBadRequest.Code())
		ctx.AbortWithStatusJSON(res.Code, res)
		return
	}

	resData, err := c.userService.Register(ctx, req)

	if err != nil {
		res := utility.ResponseError(errorUtils.MESSAGE_FAILED_REGISTER_USER, err.Error(), err.Code())
		ctx.AbortWithStatusJSON(res.Code, res)
		return
	}

	res := utility.ResponseSuccess(successUtils.MESSAGE_SUCCESS_REGISTER_USER, resData, http.StatusCreated)

	ctx.JSON(res.Code, res)
}

func (c *userController) Login(ctx *gin.Context) {
	var req dto.UserLoginRequest
	if err := ctx.ShouldBind(&req); err != nil {
		res := utility.ResponseError(errorUtils.MESSAGE_FAILED_TO_GET_DATA_FROM_BODY, errorUtils.ErrBadRequest.Error(), errorUtils.ErrBadRequest.Code())
		ctx.AbortWithStatusJSON(res.Code, res)
		return
	}

	resData, err := c.userService.Login(ctx, req)

	if err != nil {
		res := utility.ResponseError(errorUtils.MESSAGE_FAILED_LOGIN_USER, err.Error(), err.Code())
		ctx.AbortWithStatusJSON(res.Code, res)
		return
	}

	cookie := fmt.Sprintf(
		"access_token=%s; Path=/; Max-Age=3600; HttpOnly; Secure; SameSite=None",
		resData.AccessToken,
	)
	ctx.Header("Set-Cookie", cookie)

	res := utility.ResponseSuccess(successUtils.MESSAGE_SUCCESS_LOGIN_USER, nil, http.StatusOK)

	ctx.JSON(res.Code, res)
}
