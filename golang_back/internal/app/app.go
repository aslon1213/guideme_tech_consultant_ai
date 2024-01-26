package app

import (
	"aslon1213/customer_support_bot/internal/handlers"
	"aslon1213/customer_support_bot/internal/initializers"
	"aslon1213/customer_support_bot/internal/middlewares"
	"aslon1213/customer_support_bot/internal/routes"
	"context"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

type App struct {
	Fiber      *fiber.App
	Ctx        context.Context
	Mongo      *mongo.Client
	Handlers   *handlers.HandlersWrapper
	Middleware *middlewares.MiddlewaresWrapper
}

func New() *App {

	err := initializers.LoadEnvs()
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	app := &App{}
	app.Ctx = ctx

	mongoClient, err := initializers.NewMongo(app.Ctx)
	if err != nil {
		panic(err)
	}
	app.Mongo = mongoClient
	fiber := initializers.NewFiber()
	app.Fiber = fiber
	app.Handlers = handlers.New(app.Ctx, app.Mongo)
	app.RegisterRoutes()
	return app
}

func (app *App) Run() {
	app.Fiber.Listen(":9000")
}

func (app *App) Close() {
	err := app.Mongo.Disconnect(app.Ctx)
	if err != nil {
		panic(err)
	}
	err = app.Fiber.Shutdown()
	if err != nil {
		panic(err)
	}
}

func (app *App) RegisterRoutes() {

	routes.RegisterActionsRoutes(app.Fiber, app.Middleware, app.Handlers)
	routes.RegisterChatRoutes(app.Fiber, app.Middleware, app.Handlers)
	routes.RegisterDocumentsRoutes(app.Fiber, app.Middleware, app.Handlers)
}
