package routes

import (
	"aslon1213/customer_support_bot/pkg/handlers"
	"aslon1213/customer_support_bot/pkg/middlewares"

	"github.com/gofiber/fiber/v2"
)

func RegisterSttRoutes(fb *fiber.App, middleware *middlewares.MiddlewaresWrapper, handlers *handlers.HandlersWrapper) {
	stt_handlers := handlers.STTHandlers
	stt := fb.Group("/stt")
	stt.Use(middleware.ApiKeyMiddleware)
	stt.Get("/transcribe", stt_handlers.Transcribe)
	// stt.Get("", nil)
	// stt.Get("/stream", nil)
}
