package websocket

import "fmt"

type Pool struct {
	Register   chan *Client
	Unregister chan *Client
	Clients    map[*Client]bool
	Broadcast  chan Message
	Unicast    chan Message
}

func NewPool() *Pool {
	return &Pool{
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Clients:    make(map[*Client]bool),
		Broadcast:  make(chan Message),
		Unicast:    make(chan Message),
	}
}
func (pool *Pool) Start(id float64) {
	for {
		select {
		case client := <-pool.Register:
			pool.Clients[client] = true
			fmt.Println("Register Size of Connection Pool: ", len(pool.Clients), " id ", client.ID)
			for client, _ := range pool.Clients {
				fmt.Println("client ", client.ID)
				body := fmt.Sprintf("%s%f", "New User Joined... with id ", client.ID)
				client.Conn.WriteJSON(Message{Type: int(id), Body: body})
			}
			break
		case client := <-pool.Unregister:
			delete(pool.Clients, client)
			fmt.Println("Unregister Size of Connection Pool: ", len(pool.Clients))
			for client, _ := range pool.Clients {
				client.Conn.WriteJSON(Message{Type: 1, Body: "User Disconnected..."})
			}
			break
		case message := <-pool.Broadcast:
			for client, _ := range pool.Clients {
				if err := client.Conn.WriteJSON(message); err != nil {
					fmt.Println(err)
					return
				}

			}
		case message := <-pool.Unicast:
			fmt.Println("Sending Unicast id value ", id)
			for client, _ := range pool.Clients {
				if client.ID == id {
					if err := client.Conn.WriteJSON(message); err != nil {
						fmt.Println(err)
						return
					}
				}

			}

		}

	}
}
