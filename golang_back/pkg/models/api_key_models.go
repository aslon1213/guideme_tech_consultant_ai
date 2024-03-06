package models

import (
	"aslon1213/customer_support_bot/pkg/initializers"
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ApiKey struct {
	Key       string             `json:"key" bson:"key"`
	Name      string             `json:"name" bson:"name"`
	Owner     primitive.ObjectID `json:"owner" bson:"owner"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
	ExpiresAt time.Time          `json:"expires_at" bson:"expires_at"`
	Active    bool               `json:"active" bson:"active"`
}

func (a *ApiKey) UnmarshalBinary(b []byte) error {
	return json.Unmarshal(b, a)
}
func (a *ApiKey) MarshalBinary() ([]byte, error) {
	return json.Marshal(a)
}

func ValidateApiKey(apiKey string, api_keys_collection *mongo.Collection, redisClient *initializers.RedisClient) bool {
	// check if the api key is in the redis cache
	ctx := context.Background()
	_, err := redisClient.Get(ctx, apiKey).Result()
	if err == redis.Nil {
		// if the key is not in the cache, check the database
		fmt.Println("Key is not in redis")
		var key ApiKey
		err := api_keys_collection.FindOne(ctx, map[string]string{"key": apiKey}).Decode(&key)
		if err != nil {
			fmt.Println("Key is not in the database")
			return false
		}
		fmt.Println("Key is in the database")
		// if the key is in the database, add it to the cache
		redisClient.Set(ctx, apiKey, key, time.Until(key.ExpiresAt))
		return true
	}
	fmt.Println("Key is in the cache")
	return true
}
