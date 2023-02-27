package handler

import "github.com/hasbai/srs-video-chat/ws"

type payload interface {
	*ws.Client | Chat
}

type Message[T payload] struct {
	Event
	Data T `json:"data"`
}

type Event struct {
	Event string `json:"event"`
}

const (
	EventJoin  = "join"
	EventLeave = "leave"
	EventChat  = "chat"
)

func message[T payload](event string, data T) Message[T] {
	return Message[T]{Event: Event{Event: event}, Data: data}
}
