package routers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gavv/httpexpect"
	"github.com/zoedaemon/pretest/simplestorage"
)

func TestSendingMessage(t *testing.T) {

	//call your router init with SimpleAPI
	api := RegisterHandlers()

	//new simplestorage.Mapper
	DataMap := simplestorage.NewMapper()

	//save it to SimpleAPI so accessible from every handler
	api.SetCustomData(DataMap)

	//create test server
	server := httptest.NewServer(api.Server)
	defer server.Close()

	//new object for httpexpect test libs
	e := httpexpect.New(t, server.URL)

	//make connection builder
	conn := e.Builder(func(req *httpexpect.Request) {})

	const sendMessageEndpoint = "/messages/send"

	payload := map[string]interface{}{
		"message": "this is message",
	}

	//send request and evaluate response code
	conn.POST(sendMessageEndpoint).
		WithJSON(payload).
		Expect().
		Status(http.StatusCreated).JSON().Object()

	//send request, then response suppose to be method not allowed
	//coz just GET method registered for '/'
	conn.GET(sendMessageEndpoint).
		Expect().
		Status(http.StatusMethodNotAllowed).JSON().Object()

	//payload key not valid
	payload = map[string]interface{}{
		"message_invalid": "this message must be invalid",
	}
	conn.POST(sendMessageEndpoint).
		WithJSON(payload).
		Expect().
		Status(http.StatusUnprocessableEntity).JSON().Object()

}
