package jwt

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/kjasuquo/usdngn-exchange/internal/services/web"
	"net/http"
	"os"
	"time"
)

const AccessTokenValidity = time.Hour * 24

// verifyAccessToken verifies a token
func verifyToken(tokenString string, secret string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})
}

func ValidateToken(token string, secret string) (*jwt.Token, error) {
	tk, err := verifyToken(token, secret)
	if err != nil {
		return nil, fmt.Errorf("invalid token: %v", err)
	}
	if !tk.Valid {
		return nil, web.New("invalid token", http.StatusUnauthorized)
	}
	return tk, nil
}

func getClaims(token *jwt.Token) (jwt.MapClaims, error) {
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("could not get claims")
	}
	return claims, claims.Valid()
}

func ValidateAndGetClaims(tokenString string, secret string) (jwt.MapClaims, error) {
	if tokenString == "" {
		return nil, web.New("invalid token (token is empty)", http.StatusUnauthorized)
	}
	token, err := ValidateToken(tokenString, secret)
	if err != nil {
		return nil, fmt.Errorf("failed to validate token: %v", err)
	}
	claims, err := getClaims(token)
	if err != nil {
		return nil, fmt.Errorf("failed to get claims: %v", err)
	}
	return claims, nil
}

// GenerateToken generates only an access token
func GenerateToken(email string, expiryDuration time.Duration) (string, error) {
	secret := os.Getenv("JWT_SECRET")
	// Generate claims
	claims := generateClaims(email, expiryDuration)

	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func generateClaims(email string, expiryDuration time.Duration) jwt.MapClaims {
	accessClaims := jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add(expiryDuration).Unix(),
	}
	return accessClaims
}
