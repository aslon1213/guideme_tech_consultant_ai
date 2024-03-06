package routes

import (
	"aslon1213/customer_support_bot/pkg/handlers"
	"aslon1213/customer_support_bot/pkg/middlewares"
	"fmt"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
)

func RegisterTTSRoutes(fb *fiber.App, middleware *middlewares.MiddlewaresWrapper, handlers *handlers.HandlersWrapper) {

	tts_handlers := handlers.TTSHandlers
	// app.Post("/tts", middleware.ApiKeyMiddleware, tts_handlers.TTS)
	fb.Get("/tts", middleware.ApiKeyMiddleware, tts_handlers.TTSFull)
	// app.Get("/tts/:id", middleware.ApiKeyMiddleware, tts_handlers.GetTTSByID)
	// app.Delete("/tts/:id", middleware.ApiKeyMiddleware, tts_handlers.DeleteTTS)
	// app.Get("/tts/download/:id", middleware.ApiKeyMiddleware, tts_handlers.DownloadTTS)

	/// using websocket for streaming the inference of text to speech
	fb.Get("/tts/stream", func(c *fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(c) {
			c.Locals("allowed", true)
			return c.Next()
		}
		fmt.Println("TTS stream")
		return fiber.ErrUpgradeRequired
	})
	fb.Get("/tts/stream", middleware.ApiKeyMiddleware, websocket.New(tts_handlers.TTSStream))
}
