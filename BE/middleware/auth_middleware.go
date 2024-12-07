package middleware

import (
	"github.com/Yutsss/FP-PBKK-GOLANG/BE/dto"
	"github.com/Yutsss/FP-PBKK-GOLANG/BE/utility"
	errorUtils "github.com/Yutsss/FP-PBKK-GOLANG/BE/utility/error"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func AuthMiddleware(jwtUtils utility.JWTUtils) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString, err := ctx.Cookie("access_token")
		if err != nil {
			res := utility.ResponseError(errorUtils.MESSAGE_FAILED_TO_AUTHORIZE_USER, errorUtils.ErrUnauthorized.Error(), errorUtils.ErrUnauthorized.Code())
			ctx.AbortWithStatusJSON(res.Code, res)
			return
		}

		token, err := jwtUtils.ValidateToken(tokenString)
		if err != nil {
			res := utility.ResponseError(errorUtils.MESSAGE_FAILED_TO_AUTHORIZE_USER, errorUtils.ErrUnauthorized.Error(), errorUtils.ErrUnauthorized.Code())
			ctx.AbortWithStatusJSON(res.Code, res)
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			user := dto.AuthPayload{
				UserID: uint(claims["user_id"].(float64)),
				Role:   claims["role"].(string),
			}

			ctx.Set("user", user)
		}

		ctx.Next()
	}
}
