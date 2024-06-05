package demo

// doc: https://golangbot.com/go-websocket-server/
import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/gorilla/websocket"
)

type webSocketHandler struct {
	upgrader websocket.Upgrader
}

func (wsh webSocketHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c, err := wsh.upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("error %s when upgrading connection to websocket", err)
		return
	}
	defer func() {
		log.Println("closing connection")
		c.Close()
	}()
	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			return
		}
		if mt == websocket.BinaryMessage {
			err = c.WriteMessage(websocket.TextMessage, []byte("server doesn't support binary messages"))
			if err != nil {
				log.Printf("Error %s when sending message to client", err)
			}
			return
		}
		log.Printf("Receive message %s", string(message))
		if strings.Trim(string(message), "\n") != "start" {
			err = c.WriteMessage(websocket.TextMessage, []byte("You did not say the magic word!"))
			if err != nil {
				log.Printf("Error %s when sending message to client", err)
				return
			}
			continue
		}
		log.Println("start responding to client...")
		i := 1
		for {
			response := fmt.Sprintf("Notification %d", i)
			err = c.WriteMessage(websocket.TextMessage, []byte(response))
			if err != nil {
				return
			}
			i = i + 1
			time.Sleep(2 * time.Second)
		}
	}
}

func TestServer(t *testing.T) {
	/*
		curl -i --header "Upgrade: websocket" \
			--header "Connection: Upgrade" \
			--header "Sec-WebSocket-Key: YSBzYW1wbGUgMTYgYnl0ZQ==" \
			--header "Sec-Websocket-Version: 13" \
			localhost:8080
	*/
	webSocketHandler := webSocketHandler{
		upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
	}
	http.Handle("/", webSocketHandler)
	log.Print("Starting server...")
	// log.Fatal(http.ListenAndServe("localhost:8080", nil))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
