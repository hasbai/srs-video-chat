package ws

import (
	"github.com/goccy/go-json"
	"golang.org/x/net/websocket"
	"log"
	"strings"
)

type Client struct {
	Name         string `json:"name"`
	Hub          *Hub   `json:"-"`
	conn         *websocket.Conn
	send         chan []byte
	close        chan bool
	OnConnect    func(*Client)                      `json:"-"`
	OnDisconnect func(*Client)                      `json:"-"`
	OnMessage    func(c *Client, data []byte) error `json:"-"`
}

func NewClient(conn *websocket.Conn, hub *Hub, name string) *Client {
	return &Client{
		Name:  name,
		conn:  conn,
		Hub:   hub,
		send:  make(chan []byte),
		close: make(chan bool),
	}
}

func (c *Client) readLoop() {
	for {
		data := make([]byte, 1024)
		n, err := c.conn.Read(data)
		if err != nil {
			c.Close()
			return
		}
		err = c.OnMessage(c, data[:n])
		if err != nil {
			c.Close(err.Error())
			return
		}
	}
}

func (c *Client) writeLoop() {
	for {
		select {
		case message := <-c.send:
			_, err := c.conn.Write(message)
			if err != nil {
				log.Println(err)
				c.Close()
				return
			}
		case <-c.close:
			return
		}
	}
}

func (c *Client) Close(message ...string) {
	if len(message) > 0 {
		//goland:noinspection GoUnhandledErrorResult
		c.conn.Write([]byte(strings.Join(message, " ")))
	}
	c.close <- true
	c.Hub.unregister <- c
}

func (c *Client) Send(v any) {
	data, err := json.Marshal(v)
	if err != nil {
		log.Println(err)
		return
	}
	c.send <- data
}

func (c *Client) Loop() {
	go c.readLoop()
	go c.writeLoop()
	c.Hub.register <- c

	c.OnConnect(c)
	<-c.close
	c.OnDisconnect(c)
	log.Println("Client closed")
}
