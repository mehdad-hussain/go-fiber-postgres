package middleware

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte("your_jwt_secret")

// GenerateJWT generates a JWT token for the user
func GenerateJWT(userID uint) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 72).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

// JWTMiddleware is middleware that checks the JWT token
func JWTMiddleware(c *fiber.Ctx) error {
	tokenString := c.Get("Authorization")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil || !token.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Invalid token"})
	}

	return c.Next()
}
