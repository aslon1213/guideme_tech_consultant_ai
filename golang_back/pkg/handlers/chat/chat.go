package chat

import (
	"aslon1213/customer_support_bot/pkg/grpc/toclassifier"
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ChatHandlers struct {
	ctx                context.Context
	clients_collection *mongo.Collection
}

func New(ctx context.Context, clients_collection *mongo.Collection) *ChatHandlers {
	return &ChatHandlers{
		ctx:                ctx,
		clients_collection: clients_collection,
	}
}

func (ch *ChatHandlers) OpenChat(c *fiber.Ctx) error {

	username := c.Query("username")
	// check if user exists

	con, err := grpc.Dial(
		os.Getenv("CLASSIFIER"),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	defer con.Close()
	client := toclassifier.NewToClassifierClient(con)
	query := &toclassifier.Query{
		Query:    "open chat",
		Username: username,
	}
	chat_id, err := client.OpenChat(ch.ctx, query)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"chat_id": chat_id.ChatId,
	})
}

func (ch *ChatHandlers) Query(c *fiber.Ctx) error {

	chat_id := c.Query("chat_id")
	q := c.Query("q")
	fmt.Println("Got Query", q)
	con, err := grpc.Dial(
		os.Getenv("CLASSIFIER"),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	defer con.Close()
	client := toclassifier.NewToClassifierClient(con)
	query := &toclassifier.Query{
		Query:  q,
		ChatId: chat_id,
	}
	answer, err := client.ClassifyAndAnswer(ch.ctx, query)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	o := map[string]interface{}{}
	err = json.Unmarshal(answer.Answer, &o)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	fmt.Println("Got results: ", o)

	return c.JSON(o)
}

func (ch *ChatHandlers) CloseChat(c *fiber.Ctx) error {

	chat_id := c.Query("chat_id")

	con, err := grpc.Dial(
		os.Getenv("CLASSIFIER"),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	defer con.Close()
	client := toclassifier.NewToClassifierClient(con)
	answer, err := client.CloseChat(ch.ctx, &toclassifier.ChatID{
		ChatId: chat_id,
	})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	a := map[string]string{}
	err = json.Unmarshal(answer.Answer, &a)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(a)
}

func (ch *ChatHandlers) ChatUsingWebsocket(c *websocket.Conn) {
	var chat_id string
	input := struct {
		Text     string
		Audio    []byte
		Language string
		UseTTS   bool
		UseSTT   bool
	}{}
	fmt.Println(input)
	fmt.Println(chat_id)
	// fmt.Println("hello test 2")

	for {
		messageType, p, err := c.ReadMessage()
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(messageType, p)
		c.WriteMessage(messageType, p)
	}

}
