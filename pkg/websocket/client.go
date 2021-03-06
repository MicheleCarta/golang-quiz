package websocket

import (
	"fmt"
	"log"

	"github.com/gorilla/websocket"
)

type Client struct {
	ID   float64
	Conn *websocket.Conn
	Pool *Pool
}

type Message struct {
	Type int    `json:"type"`
	Body string `json:"body"`
}

func (c *Client) Read(playerId float64) {
	defer func() {
		c.Pool.Unregister <- c
		c.Conn.Close()
	}()

	for {
		messageType, p, err := c.Conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		message := Message{Type: messageType, Body: string(p)}

		if playerId > 0 {
			c.Pool.Unicast <- message
		} else {
			c.Pool.Broadcast <- message
		}

		fmt.Printf("Message Received: %+v\n", message)
	}
}
