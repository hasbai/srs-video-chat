package handler

import (
	"fmt"
	"github.com/goccy/go-json"
	"github.com/hasbai/srs-video-chat/ws"
)

func onConnect(client *ws.Client) {
	onJoin(client)
}

func onDisconnect(client *ws.Client) {
	onLeave(client)
}

func onMessage(client *ws.Client, payload []byte) error {
	var msg Event
	err := json.Unmarshal(payload, &msg)
	if err != nil {
		return err
	}
	switch msg.Event {
	case EventJoin:
		onJoin(client)
	case EventLeave:
		onLeave(client)
	case EventChat:
		var chat Message[Chat]
		err = json.Unmarshal(payload, &chat)
		if err != nil {
			return err
		}
		return onChat(client, chat.Data)
	default:
		return fmt.Errorf("unknown event %s", msg.Event)
	}
	return nil
}

func onJoin(client *ws.Client) {
	for _, peer := range client.Hub.Clients() {
		if peer != client {
			// tell peers that a new client has joined
			peer.Send(message(EventJoin, client))
			// tell the new client about peer clients
			client.Send(message(EventJoin, peer))
		}
	}
}

func onLeave(client *ws.Client) {
	client.Hub.Broadcast(message(EventLeave, client))
}

func onChat(client *ws.Client, chat Chat) error {
	chat.From = client.Name
	if chat.To == "" {
		client.Hub.Broadcast(message(EventChat, chat))
	} else {
		peer, ok := client.Hub.Clients()[chat.To]
		if !ok {
			return fmt.Errorf("user %s not found", chat.To)
		}
		peer.Send(message(EventChat, chat))
	}
	return nil
}
