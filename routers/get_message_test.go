package routers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gavv/httpexpect"
	"github.com/zoedaemon/pretest/simplestorage"
)

func TestGetMessage(t *testing.T) {

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
	const getMessageEndpoint = "/messages/get"

	payload := map[string]interface{}{
		"message": "this is message",
	}

	//send request and evaluate response code
	obj := conn.POST(sendMessageEndpoint).
		WithJSON(payload).
		Expect().
		Status(http.StatusCreated).JSON().Object()

	//now check is message really stored in mapper
	ID := obj.Value("data").Object().Value("detail").Object().Value("id").String().Raw()
	conn.GET(getMessageEndpoint).
		WithQuery("key", ID).
		Expect().
		Status(http.StatusOK).JSON().Object()

	//test with invalid key value
	conn.GET(getMessageEndpoint).
		WithQuery("key", "invalid").
		Expect().
		Status(http.StatusNotFound).JSON().Object()

	//get message, then response suppose to be method not allowed
	//coz just POST method registered for '/message/get'
	conn.POST(getMessageEndpoint).
		Expect().
		Status(http.StatusMethodNotAllowed).JSON().Object()
}
