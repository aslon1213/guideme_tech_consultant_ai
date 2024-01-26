package routes

import (
	"aslon1213/customer_support_bot/internal/handlers"
	"aslon1213/customer_support_bot/internal/middlewares"

	"github.com/gofiber/fiber/v2"
)

func RegisterDocumentsRoutes(fb *fiber.App, md *middlewares.MiddlewaresWrapper, handlers *handlers.HandlersWrapper) {
	documents := fb.Group("/documents")
	dh := handlers.DocumentHandler
	documents.Post("/upload", dh.Upload)
	documents.Get("/train", dh.Train)
	documents.Get("", nil) // qeury
	documents.Put("/append", nil)
	documents.Put("/train", nil)
}
