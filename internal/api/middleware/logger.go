package logger

import (
	"fmt"
	"time"
	"github.com/gofiber/fiber/v2"
) 

func Logger() fiber.Handler {
	return func (c *fiber.Ctx) error {
		start := time.Now()
		err := c.Next()
		duration := time.Since(start)

		method := c.Method()
		path := c.Path()
		status := c.Response().StatusCode()

		log := fmt.Sprintf(
			"%s[%s] | %d | %s | %.1fms",
			time.Now().Format("15:04:05"),
			path,
			status,
			method,
			float64(duration.Microseconds())/1000,
		)
	fmt.Println(log)
	return err

	}
}