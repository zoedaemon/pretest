package simpleapi

import (
	"encoding/json"
	"log"
	"net/http"
)

type (

	//HandlerFunc SimpleAPI implementation
	HandlerFunc func(Scope) *Response

	//Scope SimpleAPI all attributes / field
	Scope struct {
		//TODO: should be private attributes, but need implement setter and getter,
		//		for moment keep it simple
		Request    *http.Request
		Response   http.ResponseWriter
		CustomData interface{} //e.g. map objects of files config, database connection, redis, etc
	}

	//Response that should be pass from user defined function HandlerFunc
	Response struct {
		ResponseCode int         `json:"response-code"`
		Error        error       `json:"error"`
		Data         interface{} `json:"data"`
	}
)

//RequestLogFormat default format for log the request
const RequestLogFormat = "\n\tHost : %s \n\tPath : %s"

//New SimpleAPI
func New() *Scope {
	return &Scope{}
}

/*
GetMethod register handler for Get method
path : endpoints path
handler: user defined function for handling response current endpoints
*/
func (s *Scope) GetMethod(path string, handler HandlerFunc) {

	//handle response from http lib
	http.HandleFunc(path, func(writer http.ResponseWriter, request *http.Request) {

		//intercept execution for checking correct http methodd
		if request.Method != http.MethodGet {

			//error response as json
			Data, _ := json.Marshal(&Response{
				ResponseCode: http.StatusMethodNotAllowed,
				Error:        &Error{Detail: "Status Method Not Allowed"},
			})

			//send response to the writer
			writer.WriteHeader(http.StatusMethodNotAllowed)
			writer.Write([]byte(Data))

		} else { //no error

			//execute user defined function
			response := handler(*s)

			//you can cek error from user defined handler and do something like custom log
			//if response.Error != nil {...}

			Data, _ := json.Marshal(response)
			writer.WriteHeader(response.ResponseCode)
			writer.Write([]byte(Data))

			//Print log
			log.Printf(RequestLogFormat, request.Host, request.URL.Path)
		}
	})

}

//PostMethod register handler for Get method
func (s *Scope) PostMethod(path string, h HandlerFunc) {

}

//Serve for host (hostname/IP and port concatenate with ":")
func (s *Scope) Serve(host string) {
	log.Panic(http.ListenAndServe(host, nil))
}
