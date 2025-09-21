package ws

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // Allow all origins in development
	},
}

type Hub struct {
	clients    map[*Client]bool
	register   chan *Client
	unregister chan *Client
	broadcast  chan []byte
}

type Client struct {
	hub      *Hub
	conn     *websocket.Conn
	send     chan []byte
	channels []string
}

func NewHub() *Hub {
	return &Hub{
		clients:    make(map[*Client]bool),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		broadcast:  make(chan []byte),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true
			log.Printf("Client connected. Total clients: %d", len(h.clients))

		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
				log.Printf("Client disconnected. Total clients: %d", len(h.clients))
			}

		case message := <-h.broadcast:
			for client := range h.clients {
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(h.clients, client)
				}
			}
		}
	}
}

func (h *Hub) BroadcastToChannel(channel string, message []byte) {
	for client := range h.clients {
		for _, clientChannel := range client.channels {
			if clientChannel == channel {
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(h.clients, client)
				}
				break
			}
		}
	}
}

func HandleWebSocket(c *gin.Context, hub *Hub) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Printf("WebSocket upgrade error: %v", err)
		return
	}

	client := &Client{
		hub:      hub,
		conn:     conn,
		send:     make(chan []byte, 256),
		channels: []string{},
	}

	client.hub.register <- client

	go client.writePump()
	go client.readPump()
}

func (c *Client) readPump() {
	defer func() {
		c.hub.unregister <- c
		c.conn.Close()
	}()

	for {
		var msg map[string]interface{}
		err := c.conn.ReadJSON(&msg)
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("WebSocket error: %v", err)
			}
			break
		}

		// Handle subscription messages
		if action, ok := msg["action"].(string); ok {
			switch action {
			case "subscribe":
				if channel, ok := msg["channel"].(string); ok {
					c.subscribe(channel)
				}
			case "unsubscribe":
				if channel, ok := msg["channel"].(string); ok {
					c.unsubscribe(channel)
				}
			}
		}
	}
}

func (c *Client) writePump() {
	defer c.conn.Close()

	for message := range c.send {
		if err := c.conn.WriteMessage(websocket.TextMessage, message); err != nil {
			log.Printf("WebSocket write error: %v", err)
			return
		}
	}
	_ = c.conn.WriteMessage(websocket.CloseMessage, []byte{})
}

func (c *Client) subscribe(channel string) {
	// Check if already subscribed
	for _, ch := range c.channels {
		if ch == channel {
			return
		}
	}
	c.channels = append(c.channels, channel)
	log.Printf("Client subscribed to channel: %s", channel)
}

func (c *Client) unsubscribe(channel string) {
	for i, ch := range c.channels {
		if ch == channel {
			c.channels = append(c.channels[:i], c.channels[i+1:]...)
			log.Printf("Client unsubscribed from channel: %s", channel)
			break
		}
	}
}
