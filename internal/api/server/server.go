package apiserver

import (
	"log"
	"github.com/AliasgharHeidari/mobile-numbers-mine/internal/api/handler"
	"github.com/gofiber/fiber/v2"
)

func Start() {
	app := fiber.New()

	// User CRUD routes
	app.Get("/user", handler.GetUserList)
	app.Get("/user/:id", handler.GetUserByID)
	app.Post("/user", handler.CreateUser)
	app.Put("/user/:id", handler.UpdateUserByID)
	app.Delete("/user/:id", handler.DeleteUserByID)

	// Mobile number routes
	app.Post("/user/:id/mobile-number", handler.AddMobileNumber)
	app.Delete("/user/:id/mobile-number/:number", handler.DeleteMobileNumber)

	log.Println(app.Listen(":8080"))
}
