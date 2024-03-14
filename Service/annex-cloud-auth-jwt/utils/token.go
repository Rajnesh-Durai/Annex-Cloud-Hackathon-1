package utils

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

type CustomClaims struct {
	Sub      string `json:"sub"`
	Role     string `json:"role"`
	Username string `json:"username"`
	jwt.StandardClaims
}

func GenerateToken(ttl time.Duration, payload CustomClaims, secretJWTKey string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	tokenString, err := token.SignedString([]byte(secretJWTKey))

	if err != nil {
		return "", fmt.Errorf("generating JWT Token failed: %w", err)
	}

	return tokenString, nil
}

func ValidateToken(token string, signedJWTKey string) (CustomClaims, error) {
	var claims CustomClaims
	tok, err := jwt.ParseWithClaims(token, &claims, func(jwtToken *jwt.Token) (interface{}, error) {
		if _, ok := jwtToken.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected method: %s", jwtToken.Header["alg"])
		}

		return []byte(signedJWTKey), nil
	})
	if err != nil {
		return CustomClaims{}, fmt.Errorf("invalidate token: %w", err)
	}

	if !tok.Valid {
		return CustomClaims{}, fmt.Errorf("invalid token")
	}

	return claims, nil
}
