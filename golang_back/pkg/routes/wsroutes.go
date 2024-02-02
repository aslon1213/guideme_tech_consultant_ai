package routes

import (
	"aslon1213/customer_support_bot/pkg/grpc/toclassifier"
	"aslon1213/customer_support_bot/pkg/handlers"
	"aslon1213/customer_support_bot/pkg/middlewares"
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func RegisterWsRoutes(fb *fiber.App, middlewares *middlewares.MiddlewaresWrapper, handlers *handlers.HandlersWrapper) {
	fb.Use("/chat/over_voice", func(c *fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(c) {
			c.Locals("allowed", true)
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})
	fb.Get("/chat/over_voice/:id", websocket.New(TestWEbsocket))
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
	time.Sleep(10 * time.Second)
	fmt.Println("waiting stopped 1")
	c.WriteJSON(map[string]string{"answer": "audio"})
	// get audio message from python server
	answer, err := client.GetGreetingMessage(context.Background(), &toclassifier.Username{
		Username: username,
	})
	if err != nil {
		log.Println(err)
		c.WriteMessage(websocket.TextMessage, []byte(err.Error()))
	}
	c.WriteMessage(websocket.BinaryMessage, answer.Answer)

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
			if err = c.WriteJSON(map[string]string{"answer": "text"}); err != nil {
				log.Println("write:", err)
				break
			}
			if err = c.WriteMessage(websocket.TextMessage, []byte(answer.Text)); err != nil {
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
	c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
	// close the chat with python server
}

func ReadWsMEssage(c *websocket.Conn) (int, []byte, error) {
	return c.ReadMessage()
}

func WriteMessage(c *websocket.Conn, mt int, msg []byte) error {
	return c.WriteMessage(mt, msg)
}
