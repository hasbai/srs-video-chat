package ws

import (
	"github.com/goccy/go-json"
	"log"
)

type Hub struct {
	Name        string `json:"Name"`
	clients     map[string]*Client
	clientCount int
	broadcast   chan []byte
	register    chan *Client
	unregister  chan *Client
}

func NewHub(name string) *Hub {
	return &Hub{
		Name:        name,
		broadcast:   make(chan []byte),
		register:    make(chan *Client),
		unregister:  make(chan *Client),
		clients:     make(map[string]*Client),
		clientCount: 0,
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client.Name] = client
			h.clientCount++
			log.Printf("Registered %s <-- %s, client count: %d", h.Name, client.Name, h.clientCount)
		case client := <-h.unregister:
			if _, ok := h.clients[client.Name]; ok {
				delete(h.clients, client.Name)
				h.clientCount--
			}
			log.Printf("Unregisted %s <-- %s, client count: %d", h.Name, client.Name, h.clientCount)
			if h.clientCount == 0 {
				log.Println("Closing hub", h.Name)
				delete(Hubs, h.Name)
				return
			}
		case message := <-h.broadcast:
			for name := range h.clients {
				h.clients[name].send <- message
			}
		}
	}
}

func (h *Hub) Clients() map[string]*Client {
	return h.clients
}

func (h *Hub) HasClient(name string) bool {
	_, ok := h.clients[name]
	return ok
}

func (h *Hub) Broadcast(v any) {
	data, err := json.Marshal(v)
	if err != nil {
		log.Println(err)
		return
	}
	h.broadcast <- data
}

func (h *Hub) Register(client *Client) {
	h.register <- client
}

func (h *Hub) Unregister(client *Client) {
	h.unregister <- client
}
