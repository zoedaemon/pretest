package routers

import (
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/websocket"
)

func TestWebsocketMessage(t *testing.T) {

	//call your router init with SimpleAPI
	api := RegisterHandlers()

	//create test server
	server := httptest.NewServer(api.Server)
	defer server.Close()

	// Convert http://127.0.0.1 to ws://127.0.0.
	u := "ws" + strings.TrimPrefix(server.URL, "http") + "/ws/messages/echo"

	// Connect to the server
	ws, _, err := websocket.DefaultDialer.Dial(u, nil)
	if err != nil {
		t.Fatalf("%v", err)
	}
	defer ws.Close()

	// Send message to server, read response and check to see if it's what we expect.
	for i := 0; i < 10; i++ {
		if err := ws.WriteMessage(websocket.TextMessage, []byte("hello")); err != nil {
			t.Fatalf("%v", err)
		}
		_, p, err := ws.ReadMessage()
		if err != nil {
			t.Fatalf("%v", err)
		}
		if string(p) != "Your Message : hello" {
			t.Fatalf("bad message")
		}
	}
}
