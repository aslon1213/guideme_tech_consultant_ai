package authentication

import (
	"aslon1213/customer_support_bot/pkg/initializers"
	"aslon1213/customer_support_bot/pkg/models"
	"context"
	"fmt"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type AuthenticationHandlers struct {
	ctx                 context.Context
	clients_collection  *mongo.Collection
	api_keys_collection *mongo.Collection
	redisClient         *initializers.RedisClient
}

func New(ctx context.Context, clients_collection *mongo.Collection, api_keys_collection *mongo.Collection, redisClient *initializers.RedisClient) *AuthenticationHandlers {

	return &AuthenticationHandlers{ctx, clients_collection, api_keys_collection, redisClient}
}

func (au *AuthenticationHandlers) Register(c *fiber.Ctx) error {

	user := models.User{}
	err := c.BodyParser(&user)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid request",
		})
	}
	user.ID = primitive.NewObjectID()
	pass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Internal server error",
		})
	}
	user.Password = string(pass)
	client := models.Client{
		User: user,
	}
	client.Role = "client"
	client.Documents = []models.Document{}
	client.Actionslist = []models.Action{}
	client.ApiKeys = []string{}
	client.Balance = 0
	client.Transactions = []models.Transaction{}
	client.UsageInfo = models.UsageInfo{
		TotalTokenUsage: 0,
		TotalRequests:   0,
	}
	_, err = au.clients_collection.InsertOne(au.ctx, client)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Internal server error",
			"error":   err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"message": "Registered succesfully",
	})
}

func (au *AuthenticationHandlers) Login(c *fiber.Ctx) error {

	// get user data
	user := models.User{}
	err := c.BodyParser(&user)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid request",
		})
	}
	// check if the user exists
	var client models.Client
	err = au.clients_collection.FindOne(au.ctx, bson.M{
		"username": user.Username,
	}).Decode(&client)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid credentials",
		})
	}
	// check if the password is correct
	err = bcrypt.CompareHashAndPassword([]byte(client.Password), []byte(user.Password))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid credentials",
		})
	}
	expire_time := time.Now().Add(24 * 10 * time.Hour) // 10 days
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		// "foo":  "bar",
		"user":   user.Username,
		"expire": expire_time,
	})

	// Sign and get the complete encoded token as a string using the secret
	fmt.Println("MY_SECRET: ", os.Getenv("MY_SECRET"))
	tokenString, err := token.SignedString([]byte(os.Getenv("MY_SECRET")))
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Internal server error",
		})
	}

	c.Cookie(&fiber.Cookie{
		Name:     "Authorization",
		Value:    tokenString,
		Expires:  expire_time,
		HTTPOnly: true,
		SameSite: "None",
		Secure:   true,
	})

	// generate a jwt token   ///////////////////////////////////////////////////////////////
	return c.JSON(fiber.Map{
		"message": "Login",
		"token":   tokenString,
	})
}

func (au *AuthenticationHandlers) Logout(c *fiber.Ctx) error {
	return c.Status(fiber.StatusNotImplemented).JSON(fiber.Map{
		"message": "not yet implemented - logout",
	})
}

func (au *AuthenticationHandlers) Refresh(c *fiber.Ctx) error {

	return c.Status(fiber.StatusNotImplemented).JSON(fiber.Map{
		"message": "not yet implemented - refresh token",
	})
}
