package app

import (
	"aslon1213/customer_support_bot/pkg/handlers"
	"aslon1213/customer_support_bot/pkg/initializers"
	"aslon1213/customer_support_bot/pkg/middlewares"
	"aslon1213/customer_support_bot/pkg/routes"
	"context"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

type App struct {
	Fiber      *fiber.App
	Ctx        context.Context
	Mongo      *mongo.Client
	Redis      *initializers.RedisClient
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

	// get redis client
	redisClient, err := initializers.NewRedisClient()
	if err != nil {
		panic(err)
	}
	app.Redis = redisClient

	app.Mongo = mongoClient
	fiber := initializers.NewFiber()
	app.Fiber = fiber
	app.Handlers = handlers.New(app.Ctx, app.Mongo, redisClient)
	app.Middleware = middlewares.New(app.Ctx, app.Mongo, redisClient)
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

	routes.RegisterAuthRoutes(app.Fiber, app.Middleware, app.Handlers)

	routes.RegisterWsRoutes(app.Fiber, app.Middleware, app.Handlers)
	routes.RegisterActionsRoutes(app.Fiber, app.Middleware, app.Handlers)
	routes.RegisterChatRoutes(app.Fiber, app.Middleware, app.Handlers)
	routes.RegisterDocumentsRoutes(app.Fiber, app.Middleware, app.Handlers)
	routes.RegisterAdminDashboardRoutes(app.Fiber, app.Middleware, app.Handlers)
}

func (a *App) LoadUsagesFromRedisToDatabase() {
	for {
		time.Sleep(5 * time.Minute)
		fmt.Println("Loading usages from redis to database")

		// read from redis
		// write to database
		// calculate usage
		// write to database
	}
}
