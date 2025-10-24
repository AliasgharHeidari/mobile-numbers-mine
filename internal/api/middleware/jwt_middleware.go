package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"

	"github.com/AliasgharHeidari/mobile-numbers-mine/internal/config"
	"strings"
	"time"
)

func JwtProtectedMiddleware(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "missing authorization header",
		})
	}


	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	if tokenString == authHeader {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error" : "invalid authorization header",
		})
	}

	token , err := jwt.Parse(tokenString,func (token *jwt.Token) (interface{}, error) {
		return []byte(config.AppConfig.API.JWT.SecretKey), nil
	})

	if err != nil || !token.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error" : "expired or invalid token",
		})
	}
	
	if claims , ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error" : "token expired",
			})
		}
		c.Locals("username", claims["username"])
	}

	return c.Next()
}
