package routers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gavv/httpexpect"
)

func TestSendingMessage(t *testing.T) {

	//call your router init with SimpleAPI
	api := RegisterHandlers()

	//create test server
	server := httptest.NewServer(api.Server)
	defer server.Close()

	//new object for httpexpect test libs
	e := httpexpect.New(t, server.URL)

	//make connection builder
	conn := e.Builder(func(req *httpexpect.Request) {})

	const MessageEndpoint = "/messages/send"

	payload := map[string]interface{}{
		"message": "this is message",
	}

	//send request and evaluate response code
	conn.POST(MessageEndpoint).
		WithJSON(payload).
		Expect().
		Status(http.StatusCreated).JSON().Object()

	//send request and response suppose to be coz just GET method registered for '/'
	conn.GET(MessageEndpoint).
		Expect().
		Status(http.StatusMethodNotAllowed).JSON().Object()

	payload = map[string]interface{}{
		"message_invalid": "this message must be invalid",
	}
	conn.POST(MessageEndpoint).
		WithJSON(payload).
		Expect().
		Status(http.StatusUnprocessableEntity).JSON().Object()

}
