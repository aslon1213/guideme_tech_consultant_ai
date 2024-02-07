package client

import (
	"aslon1213/customer_support_bot/pkg/initializers"
	"aslon1213/customer_support_bot/pkg/models"
	"context"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ClientHandlers struct {
	ctx                 context.Context
	clients_collection  *mongo.Collection
	api_keys_collection *mongo.Collection
	redisClient         *initializers.RedisClient
}

func New(ctx context.Context, clients_collection *mongo.Collection, api_keys_collection *mongo.Collection, redisClient *initializers.RedisClient) *ClientHandlers {
	return &ClientHandlers{ctx, clients_collection, api_keys_collection, redisClient}
}

func (cl *ClientHandlers) GetClientInfo(c *fiber.Ctx) error {

	return c.SendString("Hello, World 👋!")

}

func (cl *ClientHandlers) GetDashboard(c *fiber.Ctx) error {
	return c.SendString("Hello, World 👋!")
}

func (cl *ClientHandlers) CreateAPromptTemplate(c *fiber.Ctx) error {

	template := models.PromptTemplate{}

	if err := c.BodyParser(&template); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid request",
		})
	}

	client, ok := c.Locals("client").(*models.Client)

	if !ok {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid request",
		})
	}

	_, err := cl.clients_collection.UpdateByID(cl.ctx, client.ID, bson.M{"$push": bson.M{"prompt_templates": template}})
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Internal server error",
		})
	}
	return c.JSON(fiber.Map{
		"message":  "Prompt template created successfully",
		"template": template,
	})
}

func (cl *ClientHandlers) GetPromptTemplates(c *fiber.Ctx) error {

	client, ok := c.Locals("client").(*models.Client)

	if !ok {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid request",
		})
	}

	return c.JSON(fiber.Map{
		"templates": client.PromptTemplates,
	})
}

func (cl *ClientHandlers) GetUsageInfo(c *fiber.Ctx) error {

	client, ok := c.Locals("client").(*models.Client)

	if !ok {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid request",
		})
	}

	return c.JSON(fiber.Map{
		"usage_info": client.UsageInfo,
	})

}

func (cl *ClientHandlers) GetBalance(c *fiber.Ctx) error {

	client, ok := c.Locals("client").(*models.Client)

	if !ok {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid request",
		})
	}

	return c.JSON(fiber.Map{
		"balance": client.Balance,
	})

}

func (cl *ClientHandlers) GetTransactions(c *fiber.Ctx) error {
	client, ok := c.Locals("client").(*models.Client)

	if !ok {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid request",
		})
	}

	return c.JSON(fiber.Map{
		"transactions": client.Transactions,
	})
}
