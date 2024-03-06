package routes

import (
	"aslon1213/customer_support_bot/pkg/handlers"
	"aslon1213/customer_support_bot/pkg/middlewares"

	"github.com/gofiber/fiber/v2"
)

func RegisterActionsRoutes(fb *fiber.App, md *middlewares.MiddlewaresWrapper, handlers *handlers.HandlersWrapper) {
	ah := handlers.ActionsHandler
	fb.Get("/can", md.ApiKeyMiddleware, ah.Can) // classify and answer ---
	//					 //	 go throught classification process and decide whether actions should be taken or should be given answer to question
	actions := fb.Group("/actions")
	actions.Put("/train", ah.Train)
	actions.Post("/upload", nil)
	actions.Get("/get_all", ah.GetAllActions)
	actions.Get("", ah.QueryActions)
	actions.Put("/append", nil)
}
