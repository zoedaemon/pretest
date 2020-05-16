package routers

import (
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"log"
	"net/http"

	"github.com/zoedaemon/pretest/simpleapi"
	"github.com/zoedaemon/pretest/simplestorage"
)

//define data json payload in body
type Data struct {
	ID      string `json:"id,omitempty"`
	Message string `json:"message"`
}

//SendingMessage handle incoming message
func SendingMessage(scope simpleapi.Scope) *simpleapi.Response {

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

	//DONE: store the data
	//get custom data that hold Mapper
	CData := scope.GetCustomData().(*simplestorage.Mapper)

	//generate key by message then put it in to Mapper
	Key := generateKey(data.Message)
	CData.Put(string(Key), data.Message)
	data.ID = Key
	CData.PrintData()

	//return message passed and created
	return &simpleapi.Response{
		ResponseCode: http.StatusCreated, //http.StatusInternalServerError,
		Data: map[string]interface{}{
			"detail": data,
			"info":   "success send/store message",
		},
	}
}

//generateKey just hash with sha1 algorithm
func generateKey(input string) string {
	hash := sha1.New()
	hash.Write([]byte(input))
	result := hash.Sum([]byte(nil))
	return base64.URLEncoding.EncodeToString(result)
}
