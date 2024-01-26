package actions

import (
	"aslon1213/customer_support_bot/pkg/grpc/toclassifier"
	"context"
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"

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
	username := c.Query("username")
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
			Query:    q,
			Username: username,
		},
	)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(general_answer)
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

func (a *ActionsWrappers) Train(c *fiber.Ctx) error {
	username := c.Query("username")
	var input []struct {
		Name    string `json:"name"`
		Actions []struct {
			Type    string `json:"type"`
			Element string `json:"element"`
			Value   string `json:"value"`
		}
		CanBeFormatted bool `json:"can_be_formatted"`
	}
	err := c.BodyParser(&input)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
			"input": c.Body(),
		})
	}
	// fmt.Println("got input:", input)
	actions := []*toclassifier.ActionFull{}
	for _, v := range input {
		actionfull := toclassifier.ActionFull{
			Name:           v.Name,
			CanBeFormatted: v.CanBeFormatted,
			Username:       username,
		}
		for _, action := range v.Actions {
			action := toclassifier.Action{
				Type:    action.Type,
				Element: action.Element,
				Value:   action.Value,
			}
			actionfull.Actions = append(actionfull.Actions, &action)
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
	return c.JSON(fiber.Map{
		"message": response,
	})
}
