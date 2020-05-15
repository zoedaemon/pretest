package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gavv/httpexpect"
	"github.com/zoedaemon/pretest/routers"
)

func TestGetRoot(t *testing.T) {

	//TODO: init in tests package
	Host := "localhost:8080"

	//call your router init with SimpleAPI
	api := routers.RegisterHandlers(Host)

	//create test server
	server := httptest.NewServer(api.Server)
	defer server.Close()

	//new object for httpexpect test libs
	e := httpexpect.New(t, server.URL)

	//make connection builder
	conn := e.Builder(func(req *httpexpect.Request) {})

	//send request and evaluate response code
	conn.GET("/").
		Expect().
		Status(http.StatusOK).JSON().Object()

	//send request and response suppose to be coz just GET method registered for '/'
	conn.POST("/").
		Expect().
		Status(http.StatusMethodNotAllowed).JSON().Object()
}
