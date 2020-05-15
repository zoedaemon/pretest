package routers

import (
	"net/http"

	"github.com/zoedaemon/pretest/simpleapi"
)

//Root path
func Root(scope simpleapi.Scope) *simpleapi.Response {
	return &simpleapi.Response{
		//GOTCHA : errors.New(..) cannot show up in response, implement custom Error datatype instead
		Error:        nil,           //&simpleapi.Error{Detail: "error sample"},
		ResponseCode: http.StatusOK, //http.StatusInternalServerError,
		Data:         map[string]string{"version": "API v0.1"},
	}
}
