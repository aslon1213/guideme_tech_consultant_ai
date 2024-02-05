package middlewares

import (
	"aslon1213/customer_support_bot/pkg/initializers"
	"aslon1213/customer_support_bot/pkg/models"
	"context"
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MiddlewaresWrapper struct {
	ctx                 context.Context
	clients_collection  *mongo.Collection
	api_keys_collection *mongo.Collection
	redisClient         *initializers.RedisClient
}

func New(ctx context.Context, mongoClient *mongo.Client, redisClient *initializers.RedisClient) *MiddlewaresWrapper {
	clients_collection := mongoClient.Database("customer_support_bot").Collection("clients")
	api_keys_collection := mongoClient.Database("customer_support_bot").Collection("api_keys")
	return &MiddlewaresWrapper{ctx, clients_collection, api_keys_collection, redisClient}
}

func (md *MiddlewaresWrapper) AuthenticationMiddleware(c *fiber.Ctx) error {
	fmt.Println("in authentication middleware")

	tokenString := c.Cookies("Authorization", "")
	if tokenString == "" {
		return c.Status(401).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	fmt.Println("tokenString: ", tokenString)

	// check the token is valid or not
	// Parse takes the token string and a function for looking up the key. The latter is especially
	// useful if you use multiple keys for your application.  The standard is to use 'kid' in the
	// head of the token to identify which key to use, but the parsed token (head and claims) is provided
	// to the callback, providing flexibility.
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		mysecret := os.Getenv("MY_SECRET")
		fmt.Println("MY_SECRET: ", mysecret)
		return []byte(mysecret), nil
	})
	// if err != nil {
	// 	return c.Status(401).JSON(fiber.Map{
	// 		"error": "Invalid token",
	// 		"err":   err.Error(),
	// 	})
	// }

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		client := &models.Client{}
		username := claims["user"].(string)
		err := md.clients_collection.FindOne(md.ctx, bson.M{
			"username": username,
		}).Decode(client)
		if err != nil {
			return c.Status(401).JSON(fiber.Map{
				"error": "User not found",
			})
		}
		c.Locals("client", client)
		c.Next()
	} else {
		fmt.Println(err)
		return c.Status(401).JSON(fiber.Map{
			"error": "Invalid token",
		})
	}
	return nil
}

func (md *MiddlewaresWrapper) ApiKeyMiddleware(c *fiber.Ctx) error {
	fmt.Println("in api key middleware")
	apiKey := c.Get("api-key")

	// check in databases
	ok := models.ValidateApiKey(apiKey, md.api_keys_collection, md.redisClient)
	if !ok {
		return c.Status(401).JSON(fiber.Map{
			"error": "Invalid API key",
		})
	}

	return c.Next()
}

func (md *MiddlewaresWrapper) RateLimitMiddleware(c *fiber.Ctx) error {
	return c.Next()
}

func (md *MiddlewaresWrapper) CalculateTokenUsageMiddleware(c *fiber.Ctx) error {
	fmt.Println("Calculating token usage")
	return c.Next()
}

func (md *MiddlewaresWrapper) CheckTokenMiddleware(c *fiber.Ctx) error {
	return c.Next()
}

func (md *MiddlewaresWrapper) CalculateResponseTimeMiddleware(c *fiber.Ctx) error {
	return c.Next()
}
