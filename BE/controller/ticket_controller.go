package controller

import (
	"github.com/Yutsss/FP-PBKK-GOLANG/BE/dto"
	"github.com/Yutsss/FP-PBKK-GOLANG/BE/service"
	"github.com/Yutsss/FP-PBKK-GOLANG/BE/utility"
	errorUtils "github.com/Yutsss/FP-PBKK-GOLANG/BE/utility/error"
	successUtils "github.com/Yutsss/FP-PBKK-GOLANG/BE/utility/success"
	"github.com/gin-gonic/gin"
	"net/http"
)

type (
	TicketController interface {
		Create(ctx *gin.Context)
		GetAll(ctx *gin.Context)
		GetById(ctx *gin.Context)
		GetByUserId(ctx *gin.Context)
	}

	ticketController struct {
		ticketService service.TicketService
	}
)

func NewTicketController(ts service.TicketService) TicketController {
	return &ticketController{
		ticketService: ts,
	}
}

func (c *ticketController) Create(ctx *gin.Context) {
	var req dto.CreateTicketRequest

	user := ctx.MustGet("user")

	req.UserID = user.(dto.AuthPayload).UserID

	if err := ctx.ShouldBind(&req); err != nil {
		res := utility.ResponseError(errorUtils.MESSAGE_FAILED_TO_GET_DATA_FROM_BODY, errorUtils.ErrBadRequest.Error(), errorUtils.ErrBadRequest.Code())
		ctx.AbortWithStatusJSON(res.Code, res)
		return
	}

	resData, err := c.ticketService.Create(ctx.Request.Context(), req)

	if err != nil {
		res := utility.ResponseError(errorUtils.MESSAGE_FAILED_CREATE_TICKET, err.Error(), err.Code())
		ctx.AbortWithStatusJSON(res.Code, res)
		return
	}

	res := utility.ResponseSuccess(successUtils.MESSAGE_SUCCESS_CREATE_TICKET, resData, http.StatusCreated)
	ctx.JSON(res.Code, res)
}

func (c *ticketController) GetAll(ctx *gin.Context) {
	resData, err := c.ticketService.GetAll(ctx.Request.Context())

	if err != nil {
		res := utility.ResponseError(errorUtils.MESSAGE_FAILED_GET_ALL_TICKET, err.Error(), err.Code())
		ctx.AbortWithStatusJSON(res.Code, res)
		return
	}

	if len(resData.Tickets) == 0 {
		res := utility.ResponseSuccess(successUtils.MESSAGE_SUCCESS_GET_ALL_TICKET_EMPTY, resData, http.StatusOK)
		ctx.JSON(res.Code, res)
		return
	}

	res := utility.ResponseSuccess(successUtils.MESSAGE_SUCCESS_GET_ALL_TICKET, resData, http.StatusOK)
	ctx.JSON(res.Code, res)
}

func (c *ticketController) GetById(ctx *gin.Context) {
	var req dto.GetTicketByIDRequest
	var err errorUtils.CustomError

	req.ID, err = utility.StringToUUID(ctx.Param("ticket_id"))

	if err != nil {
		res := utility.ResponseError(errorUtils.MESSAGE_FAILED_TO_GET_DATA_FROM_BODY, errorUtils.ErrBadRequest.Error(), errorUtils.ErrBadRequest.Code())
		ctx.AbortWithStatusJSON(res.Code, res)
		return
	}

	resData, err := c.ticketService.GetById(ctx.Request.Context(), req)

	if err != nil {
		res := utility.ResponseError(errorUtils.MESSAGE_FAILED_GET_TICKET, err.Error(), err.Code())
		ctx.AbortWithStatusJSON(res.Code, res)
		return
	}

	res := utility.ResponseSuccess(successUtils.MESSAGE_SUCCESS_GET_TICKET, resData, http.StatusOK)

	ctx.JSON(res.Code, res)
}

func (c *ticketController) GetByUserId(ctx *gin.Context) {
	var req dto.GetTicketByUserIDRequest
	var err errorUtils.CustomError

	req.UserID = ctx.MustGet("user").(dto.AuthPayload).UserID

	resData, err := c.ticketService.GetByUserId(ctx.Request.Context(), req)

	if err != nil {
		res := utility.ResponseError(errorUtils.MESSAGE_FAILED_GET_TICKET, err.Error(), err.Code())
		ctx.AbortWithStatusJSON(res.Code, res)
		return
	}

	if len(resData.Tickets) == 0 {
		res := utility.ResponseSuccess(successUtils.MESSAGE_SUCCESS_GET_ALL_TICKET_EMPTY, resData, http.StatusOK)
		ctx.JSON(res.Code, res)
		return
	}

	res := utility.ResponseSuccess(successUtils.MESSAGE_SUCCESS_GET_TICKET, resData, http.StatusOK)

	ctx.JSON(res.Code, res)
}
