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
		res := utility.ResponseError(errorUtils.MESSAGE_FAILED_TO_GET_DATA_FROM_BODY, errorUtils.ErrBadRequest)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		fmt.Println(err)
		return
	}

	resData, err := c.userService.Register(ctx, req)

	if err != nil {
		res := utility.ResponseError(errorUtils.MESSAGE_FAILED_REGISTER_USER, err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	res := utility.ResponseSuccess(successUtils.MESSAGE_SUCCESS_REGISTER_USER, resData)

	ctx.JSON(http.StatusOK, res)
}
