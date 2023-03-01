package handler

import (
	"github.com/goccy/go-json"
	"github.com/hasbai/srs-video-chat/ws"
	"testing"
)

func testMarshal[T Payload](v T, t *testing.T) {
	data := message(EventJoin, v)
	_, err := json.Marshal(data)
	if err != nil {
		t.Errorf("failed to marshal message: %v", err)
	}
}

func TestMarshalMessage(t *testing.T) {
	var client ws.Client
	testMarshal(&client, t)

	var chat Chat
	testMarshal(chat, t)
}
