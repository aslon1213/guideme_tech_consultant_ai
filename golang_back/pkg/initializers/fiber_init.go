package initializers

import (
	"fmt"
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
	if os.Getenv("FIBER_MODE") == "production" {
		// recover_config := recover.Config{

		// }
		fmt.Println("Using recover middleware")
		app.Use(recover.New())
	}
	app.Static("/static", "./pkg/web/static", fiber.Static{
		Browse: true,
	})
	return app

}
