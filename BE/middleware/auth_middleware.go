package middleware

import (
	"github.com/Yutsss/FP-PBKK-GOLANG/BE/utility"
	errorUtils "github.com/Yutsss/FP-PBKK-GOLANG/BE/utility/error"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware(jwtUtils utility.JWTUtils) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString, err := ctx.Cookie("access_token")
		if err != nil {
			res := utility.ResponseError(errorUtils.MESSAGE_FAILED_TO_AUTHORIZE_USER, errorUtils.ErrUnauthorized.Error(), errorUtils.ErrUnauthorized.Code())
			ctx.AbortWithStatusJSON(res.Code, res)
			return
		}

		payload, err := jwtUtils.GetPayload(tokenString)

		if err != nil {
			res := utility.ResponseError(errorUtils.MESSAGE_FAILED_TO_AUTHORIZE_USER, errorUtils.ErrUnauthorized.Error(), errorUtils.ErrUnauthorized.Code())
			ctx.AbortWithStatusJSON(res.Code, res)
			return
		}

		ctx.Set("user", payload)

		ctx.Next()
	}
}
