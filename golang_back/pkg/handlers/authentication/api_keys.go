package authentication

import (
	"aslon1213/customer_support_bot/pkg/models"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
)

func (au *AuthenticationHandlers) CreateKey(c *fiber.Ctx) error {
	client, ok := c.Locals("client").(*models.Client)
	if !ok {
		return c.Status(500).JSON(fiber.Map{
			"error": "Internal server error",
		})
	}
	expire := c.QueryInt("expire")
	if expire != 30 {
		if expire != 60 {
			if expire != 90 {
				return c.Status(400).JSON(fiber.Map{
					"error": "Invalid expire value - it should be one of 30, 60, 90",
				})
			}
		}
	}

	// get the user information
	api_key := &models.ApiKey{}
	// generate a key - random using uuid
	key_uuid, err := uuid.NewRandom()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Internal server error ----> " + err.Error(),
		})
	}

	// format the key
	prefix := "skg_"
	key_string := strings.Replace(key_uuid.String(), "-", "", -1)
	key_string = prefix + key_string

	api_key.Key = key_string
	api_key.Owner = client.ID
	api_key.Name = c.Query("name")
	api_key.Active = true
	api_key.CreatedAt = time.Now()
	api_key.ExpiresAt = time.Now().Add(time.Hour * 24 * time.Duration(expire))
	// push to mongo - api_keys collection and client_collection
	_, err = au.api_keys_collection.InsertOne(au.ctx, api_key)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Internal server error ----> " + err.Error(),
		})
	}
	// add the key to the client
	_, err = au.clients_collection.UpdateByID(au.ctx, client.ID, bson.M{
		"$push": bson.M{
			"api_keys": key_string,
		},
	})
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Internal server error ----> " + err.Error(),
		})
	}
	// set it to the redis cache
	res, err := au.redisClient.Set(au.ctx, key_string, api_key, time.Until(api_key.ExpiresAt)).Result()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Internal server error ----> " + err.Error(),
			"res":   res,
		})
	}
	return c.JSON(fiber.Map{
		"message": "Key created successfully",
	})
}

func (au *AuthenticationHandlers) DeleteKey(c *fiber.Ctx) error {

	key_string := c.Get("api-key")
	if key_string == "" {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid request - api-key is required",
		})
	}
	// get the user information
	client := c.Locals("client").(*models.Client)
	// remove the key from the client
	_, err := au.clients_collection.UpdateByID(au.ctx, client.ID, bson.M{
		"$pull": bson.M{
			"api_keys": key_string,
		},
	})
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Internal server error ----> " + err.Error(),
		})
	}
	// remove the key from the redis cache
	res, err := au.redisClient.Del(au.ctx, key_string).Result()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Internal server error ----> " + err.Error(),
			"res":   res,
		})
	}
	// remove the key from the mongo collection
	_, err = au.api_keys_collection.DeleteOne(au.ctx, bson.M{
		"key": key_string,
	})
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Internal server error ----> " + err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Key deleted successfully",
	})
}

func (au *AuthenticationHandlers) GetInfo(c *fiber.Ctx) error {

	api_keys := []models.ApiKey{}
	client := c.Locals("client").(*models.Client)
	// get the keys from the client
	res, err := au.api_keys_collection.Find(au.ctx, bson.M{
		"owner": client.ID,
	})
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Internal server error ----> " + err.Error(),
		})
	}
	err = res.All(au.ctx, &api_keys)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Internal server error ----> " + err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"keys": api_keys,
	})
}

func (au *AuthenticationHandlers) ListTheKeys(c *fiber.Ctx) error {
	client, ok := c.Locals("client").(*models.Client)
	if !ok {
		return c.Status(500).JSON(fiber.Map{
			"error": "Internal server error",
		})
	}
	keys := client.ApiKeys

	return c.JSON(fiber.Map{
		"keys": keys,
	})
}
