package routes

import (
	"aslon1213/customer_support_bot/pkg/grpc/toclassifier"
	"aslon1213/customer_support_bot/pkg/handlers"
	"aslon1213/customer_support_bot/pkg/middlewares"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var SavedConnections = make(map[string]*websocket.Conn)

func RegisterWsRoutes(fb *fiber.App, middlewares *middlewares.MiddlewaresWrapper, handlers *handlers.HandlersWrapper) {
	fb.Use("/chat/over_voice", func(c *fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(c) {
			c.Locals("allowed", true)
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})
	fb.Get("/chat/over_voice/", websocket.New(TestWEbsocket))
	fb.Get("/chat/call/:username", MakeACall)
}

func MakeACall(c *fiber.Ctx) error {
	username := c.Params("username")
	if username == "" {
		username = "test"
	}
	chat_id := c.Query("chat_id")
	if chat_id == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "chat_id is required"})
	}

	conn, ok := SavedConnections[chat_id]
	if !ok {
		return c.JSON(fiber.Map{"error": "user is not connected"})
	}
	// get audio message from python server
	conn.WriteJSON(map[string]string{"answer": "text", "action": "call"})
	con, err := grpc.Dial(os.Getenv("CLASSIFIER"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Println(err)
		return c.JSON(fiber.Map{"error": err.Error()})
	}
	defer con.Close()
	client := toclassifier.NewToClassifierClient(con)
	answer, err := client.GetGreetingMessage(context.Background(), &toclassifier.Query{
		Username: username,
		ChatId:   chat_id,
	})
	if err != nil {
		log.Println(err)
		return c.JSON(fiber.Map{"error": err.Error()})
	}
	// parse the answer and send it to the user
	conn.WriteMessage(websocket.TextMessage, answer.Answer)
	return c.JSON(fiber.Map{"message": "call is made"})
}

func ClearConnection(chatid string) {
	delete(SavedConnections, chatid)
	fmt.Println("Deleted the connection of ", chatid)
	fmt.Println("Saved Connections: ", SavedConnections)
}

func TestWEbsocket(c *websocket.Conn) {
	// c.Locals is added to the *websocket.Conn
	log.Println(c.Locals("allowed"))  // true
	log.Println(c.Params("id"))       // 123
	log.Println(c.Query("v"))         // 1.0
	log.Println(c.Cookies("session")) // ""

	username := c.Query("username")
	if username == "" {
		username = "test"
	}
	// websocket.Conn bindings https://pkg.go.dev/github.com/fasthttp/websocket?tab=doc#pkg-index
	var (
		mt  int
		msg []byte
		err error
	)

	con, err := grpc.Dial(os.Getenv("CLASSIFIER"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Println(err)
		return
	}
	defer con.Close()
	client := toclassifier.NewToClassifierClient(con)
	// waiting_stopped := make(chan bool)

	chat_id_client, err := client.OpenChat(context.Background(), &toclassifier.Query{
		Username: username,
	})
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println("waiting stopped 1")
	SavedConnections[chat_id_client.ChatId] = c
	// get audio message from python server
	fmt.Println("chat id:", chat_id_client.ChatId)
	for {

		// read query,answer,question or anthing sentence from client
		if mt, msg, err = c.ReadMessage(); err != nil {
			log.Println("read:", err)

			break
		}
		log.Printf("recv: %s, %d", msg, mt)
		// get audio or text answer from python server
		answer, err := client.GiveAudioAnswerOrJustTextAnswer(context.Background(), &toclassifier.Query{
			Query:    string(msg),
			Username: username,
			ChatId:   chat_id_client.ChatId,
		})
		if err != nil {
			log.Println(err)
			continue
		}

		// if text exists then just send text otherwise send the audio stream
		if answer.Text == "" {
			if err = c.WriteJSON(map[string]string{"answer": "audio"}); err != nil {
				log.Println("write:", err)
				break
			}
			if err = c.WriteMessage(websocket.BinaryMessage, answer.Audio); err != nil {
				log.Println("write:", err)
				break
			}
		} else {
			// Input string
			input := answer.Text
			regex := regexp.MustCompile(`\d{1,3}(,\d{3})*`)

			// Find all matches of numbers in the input string
			matches := regex.FindAllString(input, -1)
			// Replace commas with empty strings in each match
			for _, match := range matches {
				// delete , from the match
				replaced := strings.ReplaceAll(match, ",", "")
				replaced = "'" + replaced + "'"
				// replace the match with the replaced string
				input = strings.Replace(input, match, replaced, 1)
			}

			// Replace commas with empty string
			for _, match := range matches {
				replaced := strings.ReplaceAll(match, ",", "")
				input = strings.Replace(input, match, replaced, 1)
			}
			if err = c.WriteJSON(map[string]string{"answer": "text"}); err != nil {
				log.Println("write:", err)
				break
			}
			log.Println("after modifitcation input:", input)
			if err = c.WriteMessage(websocket.TextMessage, []byte(input)); err != nil {
				log.Println("write:", err)
				break
			}
		}
		// file, err := os.Open("test.wav")
		// if err != nil {
		// 	log.Println(err)
		// 	return
		// }
		// defer file.Close()
		// log when user is diconnected

	}

	_, err = client.CloseChat(context.Background(), &toclassifier.ChatID{
		ChatId: chat_id_client.ChatId,
	})
	if err != nil {
		log.Println(err)
		c.WriteJSON(map[string]string{"text": err.Error()})
	}
	fmt.Println("Closed the chat id:", chat_id_client.ChatId)
	ClearConnection(chat_id_client.ChatId)
	c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
	// close the chat with python server
}

func ReadWsMEssage(c *websocket.Conn) (int, []byte, error) {
	return c.ReadMessage()
}

func WriteMessage(c *websocket.Conn, mt int, msg []byte) error {
	return c.WriteMessage(mt, msg)
}
