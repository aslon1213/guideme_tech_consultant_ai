package documents

import (
	"aslon1213/customer_support_bot/pkg/grpc/toclassifier"
	"aslon1213/customer_support_bot/pkg/models"
	"context"
	"os"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type DocumentHandlers struct {
	ctx                context.Context
	clients_collection *mongo.Collection
}

func New(ctx context.Context, clients_collection *mongo.Collection) *DocumentHandlers {
	return &DocumentHandlers{
		ctx:                ctx,
		clients_collection: clients_collection,
	}
}

func (dh *DocumentHandlers) DeleteDocument(c *fiber.Ctx) error {
	client, ok := c.Locals("client").(*models.Client)
	if !ok {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "client not found",
		})
	}

	filename := c.Query("filename")
	if filename == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "document_name is required",
		})
	}
	// check if user exists

	documents := client.Documents
	for _, document := range documents {
		if document.Name == filename {
			//connect to grpc
			conn, err := grpc.Dial(
				os.Getenv("CLASSIFIER"),
				grpc.WithTransportCredentials(insecure.NewCredentials()),
			)
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"error": err.Error(),
				})
			}
			defer conn.Close()

			// get the stub
			clasifier_client := toclassifier.NewToClassifierClient(conn)
			// delete document from db

			_, err = clasifier_client.DeleteDocument(dh.ctx, &toclassifier.Document{
				Username: client.Username,
				Filename: filename,
			})
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"error": err.Error(),
				})
			}
			return c.JSON(fiber.Map{
				"message": "document deleted",
			})

		}
	}

	return c.Status(401).JSON(fiber.Map{
		"error": "Document not found",
	})

}

func (dh *DocumentHandlers) Upload(c *fiber.Ctx) error {

	file, err := c.FormFile("file")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	filename := file.Filename
	file_type := "txt"
	username := c.Query("username")
	// check if user exists

	conn, err := grpc.Dial(
		os.Getenv("CLASSIFIER"),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	defer conn.Close()
	// get file content
	buffer := make([]byte, file.Size)
	// copy contents of file into buffer
	f, err := file.Open()
	f.Read(buffer)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	client := toclassifier.NewToClassifierClient(conn)
	msg := &toclassifier.Document{
		Username:    username,
		FileContent: buffer,
		Filename:    filename,
		FileType:    file_type,
	}
	_, err = client.SaveDocuments(dh.ctx, msg)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return nil
}

func (dh *DocumentHandlers) Train(c *fiber.Ctx) error {

	conn, err := grpc.Dial(
		os.Getenv("CLASSIFIER"),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	defer conn.Close()

	client := toclassifier.NewToClassifierClient(conn)

	if c.QueryBool("json") {
		var json_input []struct {
			Question string `json:"question"`
			Answer   string `json:"answer"`
		}
		err := c.BodyParser(&json_input)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		qa := []*toclassifier.QuestionAnswerJson{}
		for _, v := range json_input {
			qa = append(qa, &toclassifier.QuestionAnswerJson{
				Question: v.Question,
				Answer:   v.Answer,
			})
		}

		msg := &toclassifier.JsonData{
			Username: c.Query("username"),
			Qa:       qa,
		}
		_, err = client.TrainonSavedDocumentsJson(dh.ctx, msg)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		return c.JSON(fiber.Map{
			"message": "training started",
		})

	}

	_, err = client.TrainOnSavedDocuments(dh.ctx, &toclassifier.Username{
		Username: c.Query("username"),
	})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "training started",
	})
}
