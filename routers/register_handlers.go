package routers

import (
	"github.com/zoedaemon/pretest/simpleapi"
)

//RegisterHandlers register all handlers API
func RegisterHandlers() (SimpleAPI *simpleapi.Scope) {

	//create new SimpleAPI
	SimpleAPI = simpleapi.New()

	//register root path
	SimpleAPI.GetMethod("/", Root)
	SimpleAPI.PostMethod("/messages/send", SendingMessage)

	//NOTE: at moment key id get from query parameters not from path parameters
	//		i.e.  /messages/get?key={id} not /messages/get/{id}
	SimpleAPI.GetMethod("/messages/get", GetMessage)

	SimpleAPI.GetMethod("/messages", GetAllMessage)

	return SimpleAPI
}
