package routes

import (
	"aslon1213/customer_support_bot/pkg/handlers"
	"aslon1213/customer_support_bot/pkg/middlewares"

	"github.com/gofiber/fiber/v2"
)

func RegisterSttRoutes(fb *fiber.App, middleware *middlewares.MiddlewaresWrapper, handlers *handlers.HandlersWrapper) {

	stt := fb.Group("/stt")
	stt.Use(middleware.ApiKeyMiddleware)
	// stt.Get("", nil)
	// stt.Get("/stream", nil)
}
