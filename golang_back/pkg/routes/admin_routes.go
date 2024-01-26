package routes

import (
	"aslon1213/customer_support_bot/pkg/handlers"
	"aslon1213/customer_support_bot/pkg/middlewares"

	"aslon1213/customer_support_bot/pkg/webhandlers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
)

func RegisterAdminDashboardRoutes(fb *fiber.App, md *middlewares.MiddlewaresWrapper, handlers *handlers.HandlersWrapper) {

	fb.Get("/admin", adaptor.HTTPHandlerFunc(webhandlers.IndexHandler))
}
