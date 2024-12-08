package dto

import "github.com/golang-jwt/jwt/v4"

type (
	AuthRequest struct {
		AccessToken string
	}

	JwtCustomClaims struct {
		UserID int64  `json:"user_id"`
		Role   string `json:"role"`
		jwt.RegisteredClaims
	}

	AuthPayload struct {
		UserID int64  `json:"user_id"`
		Role   string `json:"role"`
	}
)
