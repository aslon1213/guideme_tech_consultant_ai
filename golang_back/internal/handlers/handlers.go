package handlers

import (
	"aslon1213/customer_support_bot/internal/handlers/actions"
	"aslon1213/customer_support_bot/internal/handlers/chat"
	"aslon1213/customer_support_bot/internal/handlers/documents"
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type HandlersWrapper struct {
	ActionsHandler  *actions.ActionsWrappers
	ChatHandlers    *chat.ChatHandlers
	DocumentHandler *documents.DocumentHandlers
}

func New(ctx context.Context, mongoClient *mongo.Client) *HandlersWrapper {
	// cllect := mongoClient.Database("customer_support_bot").Collection("users"

	clients_collection := mongoClient.Database("customer_support_bot").Collection("clients")

	return &HandlersWrapper{
		ActionsHandler:  actions.New(ctx, clients_collection),
		ChatHandlers:    chat.New(ctx, clients_collection),
		DocumentHandler: documents.New(ctx, clients_collection),
	}
}
