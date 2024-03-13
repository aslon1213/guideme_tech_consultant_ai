package actions

import (
	"aslon1213/customer_support_bot/pkg/grpc/toclassifier"
	"aslon1213/customer_support_bot/pkg/models"
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ActionsWrappers struct {
	ctx              context.Context
	clientCollection *mongo.Collection
}

func New(ctx context.Context, clientCollection *mongo.Collection) *ActionsWrappers {
	return &ActionsWrappers{
		ctx:              ctx,
		clientCollection: clientCollection,
	}
}

func (a *ActionsWrappers) Can(c *fiber.Ctx) error {

	q := c.Query("q")
	// username := c.Query("username")
	chat_id := c.Query("chat_id")
	fmt.Println("in can: ", chat_id)

	// n_results = 1

	conn, err := grpc.Dial(
		os.Getenv("CLASSIFIER"),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error",
		})
	}
	defer conn.Close()

	client := toclassifier.NewToClassifierClient(conn)
	general_answer, err := client.ClassifyAndAnswer(
		a.ctx,
		&toclassifier.Query{
			Query:  q,
			ChatId: chat_id,
		},
	)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	o := map[string]interface{}{}
	err = json.Unmarshal(general_answer.Answer, &o)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(o)
}

func (a *ActionsWrappers) AppendAction(c *fiber.Ctx) error {
	return c.Status(fiber.StatusNotImplemented).JSON(fiber.Map{
		"message": "Not implemented yet--- We are working on it.",
	})
}

func (a *ActionsWrappers) QueryActions(c *fiber.Ctx) error {

	q := c.Query("q")
	username := c.Query("username")
	// n_results := int32(c.QueryInt("n_results"))
	query := &toclassifier.Query{
		Query:    q,
		Username: username,
	}
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
	cc, err := client.QueryActions(a.ctx, query)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(cc)
}

func (a *ActionsWrappers) GetAllActions(c *fiber.Ctx) error {
	client, ok := c.Locals("client").(*models.Client)
	if !ok {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Client not found",
		})
	}

	// get from the database
	actions := []*models.Action{}
	// project only the actions
	cursor, err := a.clientCollection.Find(a.ctx, bson.M{
		"username": client.Username,
	}, options.Find().SetProjection(bson.M{
		"actions": 1,
	}))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	defer cursor.Close(a.ctx)
	err = cursor.All(a.ctx, &actions)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(actions)
}

func (a *ActionsWrappers) Train(c *fiber.Ctx) error {
	username := c.Query("username")
	var input []models.Action
	err := c.BodyParser(&input)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
			"input": c.Body(),
		})
	}
	fmt.Println("input:", input)
	// return c.JSON(input)
	// load actions to chroma database
	actions := []*toclassifier.ActionFull{}
	for _, v := range input {
		actionfull := toclassifier.ActionFull{
			Username:    username,
			Type:        v.Type,
			Description: v.Description,
			Deeplink: &toclassifier.Deeplink{
				Url:    v.Deeplink.Url,
				Params: v.Deeplink.Params,
			},
		}
		actions = append(actions, &actionfull)
	}

	// fmt.Println("actions:", actions)
	con, err := grpc.Dial(
		os.Getenv("CLASSIFIER"),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error",
		})
	}
	defer con.Close()

	client := toclassifier.NewToClassifierClient(con)
	cc, err := client.TrainActions(a.ctx)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	for _, v := range actions {
		err := cc.Send(v)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
	}
	response := &toclassifier.TrainResponse{}
	response, err = cc.CloseAndRecv()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	fmt.Println("Got response from server: ", response)

	// save to database actions - which is client database = not chroma database
	a.clientCollection.UpdateByID(a.ctx, bson.M{
		"username": username,
	},
		bson.M{
			"$set": bson.M{
				"actions": actions,
			},
		},
	)
	return c.JSON(fiber.Map{
		"message": response,
	})
}
