package apiserver

import (
	"log"

	"github.com/AliasgharHeidari/mobile-numbers-mine/internal/api/handler"
	"github.com/AliasgharHeidari/mobile-numbers-mine/internal/api/middleware"
	"github.com/AliasgharHeidari/mobile-numbers-mine/internal/config"
	"github.com/gofiber/fiber/v2"
)

func Start() {
	app := fiber.New()

	app.Use(middleware.Logger())
	app.Use(middleware.CorsMiddleware())

	app.Post("/user/login", handler.Login)
	// User CRUD routes
	app.Get("/user", middleware.JwtProtectedMiddleware, handler.GetUserList)
	app.Get("/user/:id", middleware.JwtProtectedMiddleware, handler.GetUserByID)
	app.Post("/user", middleware.JwtProtectedMiddleware, handler.CreateUser)
	app.Put("/user/:id", middleware.JwtProtectedMiddleware, handler.UpdateUserByID)
	app.Delete("/user/:id", middleware.JwtProtectedMiddleware, handler.DeleteUserByID)

	// Mobile number routes
	app.Post("/user/:id/mobile-number", middleware.JwtProtectedMiddleware, handler.AddMobileNumber)
	app.Delete("/user/:id/mobile-number/:number", middleware.JwtProtectedMiddleware, handler.DeleteMobileNumber)

	log.Println(app.Listen(config.AppConfig.API.Server.ListenString()))
}
