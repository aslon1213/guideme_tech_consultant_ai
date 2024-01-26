package routes

import (
	"aslon1213/customer_support_bot/internal/handlers"
	"aslon1213/customer_support_bot/internal/middlewares"

	"github.com/gofiber/fiber/v2"
)

func RegisterActionsRoutes(fb *fiber.App, md *middlewares.MiddlewaresWrapper, handlers *handlers.HandlersWrapper) {
	ah := handlers.ActionsHandler
	fb.Get("/can", ah.Can)
	actions := fb.Group("/actions")
	actions.Put("/train", ah.Train)
	actions.Post("/upload", nil)
	actions.Get("", ah.QueryActions)
	actions.Put("/append", nil)

}
