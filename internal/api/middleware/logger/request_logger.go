package requestlogger

import "github.com/gofiber/fiber/v2"

func RequestLogger(context *fiber.Ctx) error {
	// Log the request method and URL
	method := context.Method()
	url := context.OriginalURL()
	println("Request:", method, url)

	// Proceed to the next middleware or handler
	return context.Next()
}
