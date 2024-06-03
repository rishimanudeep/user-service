package auth

import (
	"github.com/golang-jwt/jwt/v4"
	"os"
	"time"
)

var jwtKey = []byte(os.Getenv("JWT_SECRET_KEY"))

// Claims defines the structure of the JWT claims
type MyCustomClaims struct {
	UserID    int       `json:"userID"`
	ExpiresAt time.Time `json:"expires_at"`
	jwt.RegisteredClaims
}

func GenerateJWTToken(userID int, latitude, longitude float64) (string, error) {
	claims := jwt.MapClaims{
		"userID":     userID,
		"latitude":   latitude,
		"longitude":  longitude,
		"expires_at": time.Now().Add(time.Hour * 72).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
