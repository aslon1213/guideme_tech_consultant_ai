package routes

import (
	"aslon1213/customer_support_bot/pkg/handlers"
	"aslon1213/customer_support_bot/pkg/middlewares"

	"github.com/gofiber/fiber/v2"
)

func RegisterDocumentsRoutes(fb *fiber.App, md *middlewares.MiddlewaresWrapper, handlers *handlers.HandlersWrapper) {
	documents := fb.Group("/documents", md.AuthenticationMiddleware)
	dh := handlers.DocumentHandler
	documents.Post("/upload", dh.Upload)
	documents.Delete("/delete", dh.DeleteDocument)
	documents.Get("/train", dh.Train)
	// documents.Get("", nil) // qeury
	// documents.Put("/append", nil)
	// documents.Put("/train", nil)
}
