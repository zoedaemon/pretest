package routers

import (
	"github.com/gorilla/websocket"
	"github.com/zoedaemon/pretest/simpleapi"
	"github.com/zoedaemon/pretest/simplestorage"
)

//WebsocketMessageWriter special handler for ws connection
func WebsocketMessageWriter(scope simpleapi.Scope, ws *websocket.Conn) ([]byte, error) {

	//DONE: store the data
	//get custom data that hold Mapper
	CData := scope.GetCustomData().(*simplestorage.Mapper)

	//get incoming ws message
	writer, message, err := ws.ReadMessage()
	if err != nil {
		//end of session
		return nil, err
	}

	//generate key by message from websocket then put it in to Mapper
	Key := generateKey(string(message))
	CData.Put(string(Key), message)

	//response to client
	err = ws.WriteMessage(writer, []byte(string(message)+" ("+Key+")"))
	if err != nil {
		//end of session; loop will be break
		return nil, err
	}

	//if nil loop still waiting next ws request current session
	return message, nil
}
