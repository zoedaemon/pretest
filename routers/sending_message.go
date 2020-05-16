package routers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/zoedaemon/pretest/simpleapi"
)

//SendingMessage handle incoming message
func SendingMessage(scope simpleapi.Scope) *simpleapi.Response {

	//define data json payload in body
	type Data struct {
		Message string `json:"message"`
	}
	var data Data

	//get request
	req := scope.GetRequest()

	//parse payload and convert to data instance
	err := json.NewDecoder(req.Body).Decode(&data)

	//if error or message is empty
	if err != nil || len(data.Message) <= 0 {
		log.Printf("Parse Error : %v", err)
		return &simpleapi.Response{
			Error:        &simpleapi.Error{Detail: "POST data not valid"},
			ResponseCode: http.StatusUnprocessableEntity,
		}
	}

	//TODO: store the data
	log.Printf("Post data = %+v\n", req)

	//return message passed and created
	return &simpleapi.Response{
		ResponseCode: http.StatusCreated, //http.StatusInternalServerError,
		Data:         map[string]string{"message": data.Message},
	}
}
