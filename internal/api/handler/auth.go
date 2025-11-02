package handler

import (
	"time"

	"github.com/AliasgharHeidari/mobile-numbers-mine/internal/config"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)
// Login 		godoc
// @Summary 	Login user
// @Tags 		Authentication
// @Description Authenticate user and return JWT token
// @Accept  	json
// @Produce  	json
// @Param   	request body model.LoginRequest true "login request"
// @Success 	200 {object} model.LoginSuccessResponse "login success"
// @Failure 	400 {object} model.LoginFailureResponse "invalid request body"
// @Router 		/user/login [post]
func Login(c *fiber.Ctx) error {
	type request struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	var body request
	err := c.BodyParser(&body)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request body",
		})
	}

	if body.Username != "Aliasghar" || body.Password != "1234" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "wrong credientials",
		})
	}

	claims := jwt.MapClaims{
		"username": body.Username,
		"exp": time.Now().Add(
			time.Minute * time.Duration(config.AppConfig.API.JWT.TokenDuration)).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(config.AppConfig.API.JWT.SecretKey))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to create token",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"token": tokenString,
	})
}
