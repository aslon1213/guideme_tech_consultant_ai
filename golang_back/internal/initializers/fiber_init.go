package initializers

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func NewFiber() *fiber.App {
	app := fiber.New(fiber.Config{
		BodyLimit: 100 * 1024 * 1024,
	})
	app.Use(logger.New())
	app.Get("/monitor", monitor.New())
	if os.Getenv("FIBER_MODE") != "dev" {
		// recover_config := recover.Config{

		// }
		app.Use(recover.New())
	}
	return app

}
