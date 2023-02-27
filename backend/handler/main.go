package handler

import (
	"errors"
	"fmt"
	"github.com/hasbai/srs-video-chat/ws"
	"golang.org/x/net/websocket"
	"log"
	"net/http"
)

func handshake(config *websocket.Config, req *http.Request) error {
	query := config.Location.Query()
	room := query.Get("room")
	user := query.Get("user")

	if room == "" || user == "" {
		return errors.New("room and user are required")
	}
	const maxNameSize = 16
	if len(room) > maxNameSize || len(user) > maxNameSize {
		return fmt.Errorf("room and user must be less than %d characters", maxNameSize)
	}

	hub, ok := ws.Hubs[room]
	if !ok {
		hub = ws.NewHub(room)
		ws.Hubs[room] = hub
		go hub.Run()
	} else {
		if hub.HasClient(user) {
			return fmt.Errorf("user %s already exists", user)
		}
	}
	return nil
}

func handleWS(conn *websocket.Conn) {
	conn.PayloadType = websocket.TextFrame

	query := conn.Request().URL.Query()
	room := query.Get("room")
	user := query.Get("user")
	hub := ws.Hubs[room]

	client := ws.NewClient(conn, hub, user)
	client.OnMessage = onMessage
	client.OnConnect = onConnect
	client.OnDisconnect = onDisconnect

	client.Loop()
}

func Serve(addr string) {
	log.Println("HTTP listening on", addr)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatalf("Error serving http: %s", err)
	}
}

func init() {
	http.Handle("/ws", websocket.Server{
		Handler: handleWS,
		Handshake: handshake,
	})
}
