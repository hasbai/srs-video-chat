package handler

import "github.com/hasbai/srs-video-chat/ws"

type payload interface {
	[]byte | *ws.Client | Chat
}

type Message[T payload] struct {
	Event Event `json:"event"`
	Data  T     `json:"data"`
}

type Event = string

const (
	EventJoin  Event = "join"
	EventLeave Event = "leave"
	EventChat  Event = "chat"
)

func message[T payload](event Event, data T) Message[T] {
	return Message[T]{Event: event, Data: data}
}
