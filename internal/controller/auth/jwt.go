package auth

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type Claims struct {
	ID      string `json:"id"`
	Email   string `json:"email"`
	IsAdmin bool   `json:"is_admin"`
	jwt.RegisteredClaims
}

type JWTManager struct {
	key []byte
}

func NewTokenGenerator() (*JWTManager, error) {
	key := os.Getenv("JWT_KEY")
	if key == "" {
		return nil, fmt.Errorf("JWT_KEY is not set")
	}

	return &JWTManager{key: []byte(key)}, nil
}

func (t *JWTManager) GenerateToken(email, userID string, isAdmin bool, expiresAt time.Time) (string, *Claims, error) {
	tokenID, err := uuid.NewUUID()
	if err != nil {
		return "", nil, err
	}

	claims := Claims{
		ID:      userID,
		Email:   email,
		IsAdmin: isAdmin,
		RegisteredClaims: jwt.RegisteredClaims{
			ID:        tokenID.String(),
			Issuer:    "ecomm",
			Subject:   email,
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(expiresAt),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(t.key)
	if err != nil {
		return "", nil, err
	}

	return tokenString, &claims, nil
}

func (t *JWTManager) ValidateToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return t.key, nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("invalid token")
}
