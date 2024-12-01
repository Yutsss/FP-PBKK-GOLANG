package middleware

import (
	"github.com/Yutsss/FP-PBKK-GOLANG/BE/dto"
	"github.com/Yutsss/FP-PBKK-GOLANG/BE/utility"
	errorUtils "github.com/Yutsss/FP-PBKK-GOLANG/BE/utility/error"
	"github.com/gin-gonic/gin"
)

func RoleMiddleware(roles []string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user, ok := ctx.Get("user")
		if !ok {
			res := utility.ResponseError(errorUtils.MESSAGE_FAILED_TO_AUTHORIZE_USER, errorUtils.ErrUnauthorized.Error(), errorUtils.ErrUnauthorized.Code())
			ctx.AbortWithStatusJSON(res.Code, res)
			return
		}

		userRole := user.(dto.AuthPayload).Role

		for _, role := range roles {
			if role == userRole {
				ctx.Next()
				return
			}
		}

		res := utility.ResponseError(errorUtils.MESSAGE_FAILED_TO_AUTHORIZE_USER, errorUtils.ErrNotAllowed.Error(), errorUtils.ErrNotAllowed.Code())
		ctx.AbortWithStatusJSON(res.Code, res)
	}
}
