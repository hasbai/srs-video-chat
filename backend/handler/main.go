package handler

import (
	"fmt"
	"github.com/hasbai/srs-video-chat/ws"
	"golang.org/x/net/websocket"
	"log"
	"net/http"
)

func wsSend(conn *websocket.Conn, message string) {
	//goland:noinspection GoUnhandledErrorResult
	conn.Write([]byte(message))
}

func handleWS(conn *websocket.Conn) {
	conn.PayloadType = websocket.TextFrame
	query := conn.Request().URL.Query()
	room := query.Get("room")
	user := query.Get("user")
	if room == "" || user == "" {
		wsSend(conn, "room and user are required")
		return
	}
	const maxNameSize = 16
	if len(room) > maxNameSize || len(user) > maxNameSize {
		wsSend(conn, fmt.Sprintf("room and user must be less than %d characters", maxNameSize))
		return
	}

	hub, ok := ws.Hubs[room]
	if !ok {
		hub = ws.NewHub(room)
		ws.Hubs[room] = hub
		go hub.Run()
	} else {
		if hub.HasClient(user) {
			wsSend(conn, fmt.Sprintf("user %s already exists", user))
			return
		}
	}

	client := ws.NewClient(conn, hub, user)
	client.OnMessage = onMessage
	client.OnConnect = onConnect
	client.OnDisconnect = onDisconnect

	client.Loop()
}

func init() {
	http.Handle("/ws", websocket.Server{
		Handler: handleWS,
	})
}

func Serve(addr string) {
	log.Println("HTTP listening on", addr)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatalf("Error serving http: %s", err)
	}
}
