package routers

import (
	"github.com/gorilla/websocket"
	"github.com/zoedaemon/pretest/simpleapi"
)

//WebsocketMessage special handler for ws connection
func WebsocketMessage(scope simpleapi.Scope, ws *websocket.Conn) error {

	//get incoming ws message
	writer, message, err := ws.ReadMessage()
	if err != nil {
		//end of session
		return err
	}

	//response to client
	err = ws.WriteMessage(writer, []byte("Your Message : "+string(message)))
	if err != nil {
		//end of session; loop will be break
		return err
	}

	//if nil loop still waiting next ws request current session
	return nil
}
