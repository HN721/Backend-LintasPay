package jwt

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

func GenerateToken(userID uint, email string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"email":   email,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
		"iat":     time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(os.Getenv("SECRET_KEY")))
}
func VerifyToken(tokens string) (*jwt.Token, error) {
	godotenv.Load()
	return jwt.Parse(tokens, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(os.Getenv("SECRET_KEY")), nil
	})
}
