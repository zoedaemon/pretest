package routers

import (
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gavv/httpexpect"
	"github.com/zoedaemon/pretest/simplestorage"
)

func TestGetAllMessage(t *testing.T) {

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
	const getAllMessageEndpoint = "/messages"

	////test no content coz no message sent in yet
	conn.GET(getAllMessageEndpoint).
		Expect().
		Status(http.StatusNoContent)

	////Test simulate sent 3 message
	//send 3 message
	for i := 0; i < 3; i++ {
		payload := map[string]interface{}{
			"message": "message number " + strconv.Itoa(i),
		}

		//send request and evaluate response code
		conn.POST(sendMessageEndpoint).
			WithJSON(payload).
			Expect().
			Status(http.StatusCreated).JSON().Object()
	}

	//test size of mapped data is 3 as POST in loop before
	obj := conn.GET(getAllMessageEndpoint).
		Expect().
		Status(http.StatusOK).JSON().Object()
	obj.Value("data").Object().ContainsKey("size").ValueEqual("size", 3)
}
