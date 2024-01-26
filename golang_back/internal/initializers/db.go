package initializers

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewMongo(ctx context.Context) (*mongo.Client, error) {
	fmt.Println("Connecting to mongodb")
	opts := options.Client().ApplyURI(os.Getenv("MONGO_URI"))
	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		return nil, err
	}

	if os.Getenv("FIBER_MODE") != "dev" {
		err = client.Ping(ctx, nil)
		if err != nil {
			return nil, err
		}
	}

	fmt.Println("Connected to MongoDB!")
	return client, nil
}
