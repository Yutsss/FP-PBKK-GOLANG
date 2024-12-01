package dto

import "github.com/golang-jwt/jwt/v4"

type (
	AuthRequest struct {
		AccessToken string
	}

	JwtCustomClaims struct {
		UserID uint   `json:"user_id"`
		Role   string `json:"role"`
		jwt.RegisteredClaims
	}

	AuthPayload struct {
		UserID uint   `json:"user_id"`
		Role   string `json:"role"`
	}
)
