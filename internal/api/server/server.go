package apiserver

import (
	"log"

	"github.com/AliasgharHeidari/mobile-numbers-mine/internal/api/handler"
	requestlogger "github.com/AliasgharHeidari/mobile-numbers-mine/internal/api/middleware/logger"
	"github.com/gofiber/fiber/v2"
)

func Start() {
	app := fiber.New()

	// User CRUD routes
	app.Get("/user",
		requestlogger.RequestLogger,
		handler.GetUserList)

	app.Get("/user/:id",
		requestlogger.RequestLogger,
		handler.GetUserByID)

	app.Post("/user",
		requestlogger.RequestLogger,
		handler.CreateUser)

	app.Put("/user/:id",
		requestlogger.RequestLogger,
		handler.UpdateUserByID)

	app.Delete("/user/:id",
		requestlogger.RequestLogger,
		handler.DeleteUserByID)

	// Mobile number routes
	app.Post("/user/:id/mobile-number",
		requestlogger.RequestLogger,
		handler.AddMobileNumber)

	app.Delete("/user/:id/mobile-number/:number",
		requestlogger.RequestLogger,
		handler.DeleteMobileNumber)

	log.Println(app.Listen(":8080"))
}
