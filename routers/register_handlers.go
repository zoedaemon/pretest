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

	return SimpleAPI
}
