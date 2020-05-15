package routers

import (
	"github.com/zoedaemon/pretest/simpleapi"
)

/*RegisterHandlers register all handlers API
* host : define your host, e.g. localhost:8080
**/
func RegisterHandlers(host string) (SimpleAPI *simpleapi.Scope) {

	//create new SimpleAPI
	SimpleAPI = simpleapi.New()

	//register root path
	SimpleAPI.GetMethod("/", Root)

	return SimpleAPI
}
