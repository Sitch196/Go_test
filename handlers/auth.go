package handlers

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte("supersecretkey")

// Simple user struct and in-memory user store
type User struct {
	Username string
	Password string
}

var users = []User{
	{Username: "admin", Password: "password"},
}

// Login handler
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Login(c *fiber.Ctx) error {
	var req LoginRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}
	for _, user := range users {
		if user.Username == req.Username && user.Password == req.Password {
			// Create JWT token
			token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
				"username": user.Username,
				"exp":      time.Now().Add(time.Hour * 24).Unix(),
			})
			t, err := token.SignedString(jwtSecret)
			if err != nil {
				return c.Status(500).JSON(fiber.Map{"error": "Could not login"})
			}
			return c.JSON(fiber.Map{"token": t})
		}
	}
	return c.Status(401).JSON(fiber.Map{"error": "Invalid credentials"})
}

// JWT Middleware
func JWTProtected() fiber.Handler {
	return func(c *fiber.Ctx) error {
		tokenStr := c.Get("Authorization")
		if len(tokenStr) > 7 && tokenStr[:7] == "Bearer " {
			tokenStr = tokenStr[7:]
		}
		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		})
		if err != nil || !token.Valid {
			return c.Status(401).JSON(fiber.Map{"error": "Unauthorized"})
		}
		return c.Next()
	}
}
