package auth

import (
	"gohabits/internal/domain/user/dto"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTManager struct {
	SecretKey     string
	TokenDuration time.Duration
}

func NewJWTManager(secret string) *JWTManager {
	return &JWTManager{
		SecretKey:     secret,
		TokenDuration: 24 * time.Hour,
	}
}

func (j *JWTManager) GenerateToken(userID uint, email string) (dto.GenerateTokenResponse, error) {
	claims := jwt.MapClaims{
		"sub":   userID,
		"email": email,
		"exp":   time.Now().Add(j.TokenDuration).Unix(),
		"iat":   time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedString, _ := token.SignedString([]byte(j.SecretKey))

	return dto.GenerateTokenResponse{
		Token:     signedString,
		ExpiresAt: time.Now().Add(j.TokenDuration).Format(time.RFC3339),
	}, nil
}
