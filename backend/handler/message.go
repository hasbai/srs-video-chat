package handler

import "github.com/hasbai/srs-video-chat/ws"

type Payload interface {
	*ws.Client | Chat
}

type Message[T Payload] struct {
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

func message[T Payload](event string, data T) Message[T] {
	return Message[T]{Event: Event{Event: event}, Data: data}
}
