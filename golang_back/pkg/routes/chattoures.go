package routes

import (
	"aslon1213/customer_support_bot/pkg/handlers"
	"aslon1213/customer_support_bot/pkg/middlewares"

	"github.com/gofiber/fiber/v2"
)

func RegisterChatRoutes(fb *fiber.App, md *middlewares.MiddlewaresWrapper, handlers *handlers.HandlersWrapper) {

	chat := fb.Group("/chat")
	chath := handlers.ChatHandlers
	chat.Get("/open", chath.OpenChat)
	chat.Get("/query", chath.Query)
	chat.Get("/close", chath.CloseChat)
}
