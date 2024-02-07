package handlers

import (
	"aslon1213/customer_support_bot/pkg/handlers/actions"
	"aslon1213/customer_support_bot/pkg/handlers/authentication"
	"aslon1213/customer_support_bot/pkg/handlers/chat"
	"aslon1213/customer_support_bot/pkg/handlers/client"
	"aslon1213/customer_support_bot/pkg/handlers/documents"
	"aslon1213/customer_support_bot/pkg/initializers"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type HandlersWrapper struct {
	ActionsHandler         *actions.ActionsWrappers
	ChatHandlers           *chat.ChatHandlers
	DocumentHandler        *documents.DocumentHandlers
	AuthenticationHandlers *authentication.AuthenticationHandlers
	ClientHandlers         *client.ClientHandlers
}

func New(ctx context.Context, mongoClient *mongo.Client, redisClient *initializers.RedisClient) *HandlersWrapper {
	// cllect := mongoClient.Database("customer_support_bot").Collection("users"

	clients_collection := mongoClient.Database("customer_support_bot").Collection("clients")
	indexModel := mongo.IndexModel{
		Keys: bson.M{
			"username": 1, // Index key for 'username' field
		},
		Options: options.Index().SetUnique(true), // Set the index as unique
	}

	_, err := clients_collection.Indexes().CreateOne(ctx, indexModel)
	if err != nil {
		panic(err)
	}
	api_keys_collection := mongoClient.Database("customer_support_bot").Collection("api_keys")
	// api_keys_collection - key  should be unique
	indexModel = mongo.IndexModel{
		Keys: bson.M{
			"key": 1, // Index key for 'key' field
		},
		Options: options.Index().SetUnique(true), // Set the index as unique
	}
	_, err = api_keys_collection.Indexes().CreateOne(ctx, indexModel)
	if err != nil {
		panic(err)
	}
	return &HandlersWrapper{
		ActionsHandler:         actions.New(ctx, clients_collection),
		ChatHandlers:           chat.New(ctx, clients_collection),
		DocumentHandler:        documents.New(ctx, clients_collection),
		AuthenticationHandlers: authentication.New(ctx, clients_collection, api_keys_collection, redisClient),
		ClientHandlers:         client.New(ctx, clients_collection, api_keys_collection, redisClient),
	}
}
