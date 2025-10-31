package apiserver

import (
	"log"

	_ "github.com/AliasgharHeidari/mobile-numbers-mine/docs"
	"github.com/AliasgharHeidari/mobile-numbers-mine/internal/api/handler"
	"github.com/AliasgharHeidari/mobile-numbers-mine/internal/api/middleware"
	"github.com/AliasgharHeidari/mobile-numbers-mine/internal/config"
	"github.com/gofiber/fiber/v2"
	swagger "github.com/gofiber/swagger"
)

func Start() {
	app := fiber.New()

	app.Use(middleware.Logger())
	app.Use(middleware.CorsMiddleware())

	app.Post("/user/login", handler.Login)
	// User CRUD routes
	app.Get("/user", middleware.ValidateToken, handler.GetUserList)
	app.Get("/user/:id", middleware.ValidateToken, handler.GetUserByID)
	app.Post("/user", middleware.ValidateToken, handler.CreateUser)
	app.Put("/user/:id", middleware.ValidateToken, handler.UpdateUserByID)
	app.Delete("/user/:id", middleware.ValidateToken, handler.DeleteUserByID)

	// Mobile number routes
	app.Post("/user/:id/mobile-number", middleware.ValidateToken, handler.AddMobileNumber)
	app.Delete("/user/:id/mobile-number/:number", middleware.ValidateToken, handler.DeleteMobileNumber)

	// swagger
	app.Get("/swagger/*", swagger.HandlerDefault)

	log.Println(app.Listen(config.AppConfig.API.Server.ListenString()))
}
