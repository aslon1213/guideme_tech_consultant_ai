package routes

import (
	"aslon1213/customer_support_bot/pkg/handlers"
	"aslon1213/customer_support_bot/pkg/middlewares"
	"fmt"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
)

func RegisterChatRoutes(fb *fiber.App, md *middlewares.MiddlewaresWrapper, handlers *handlers.HandlersWrapper) {

	chat := fb.Group("/chat", md.ApiKeyMiddleware)
	chath := handlers.ChatHandlers
	chat.Get("/open", chath.OpenChat)
	chat.Get("/query", chath.Query)
	chat.Get("/can", handlers.ActionsHandler.Can) // classify and answer --- "can"
	//					 //	 go throught classification process and decide whether actions should be taken or should be given answer to question
	chat.Get("/close", chath.CloseChat)

	// websocket for chat streaming
	fb.Get("/chat/ws", md.ApiKeyMiddleware, func(c *fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(c) {
			c.Locals("allowed", true)
			return c.Next()
		}
		fmt.Println("Chat can be build with websocket")
		return fiber.ErrUpgradeRequired
	})

	fb.Get("/chat/ws", md.ApiKeyMiddleware, websocket.New(chath.ChatUsingWebsocket))
}
