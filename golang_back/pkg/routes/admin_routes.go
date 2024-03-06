package routes

import (
	"aslon1213/customer_support_bot/pkg/handlers"
	"aslon1213/customer_support_bot/pkg/middlewares"

	"github.com/gofiber/fiber/v2"
)

func RegisterAdminDashboardRoutes(fb *fiber.App, md *middlewares.MiddlewaresWrapper, handlers *handlers.HandlersWrapper) {
	client := fb.Group("/client", md.AuthenticationMiddleware)
	// fb.Get("/client", adaptor.HTTPHandlerFunc(webhandlers.IndexHandler))
	client.Get("/dashboard", handlers.ClientHandlers.GetDashboard)

}
