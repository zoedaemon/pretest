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
		//hold server http
		Server *http.ServeMux

		//derivate Context Class
		Context

		//e.g. map objects of files config, database connection, redis, etc
		CustomData interface{}
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
	return &Scope{
		Server: http.NewServeMux(),
	}
}

/*GetMethod register handler for Get method
* path 	: endpoints path
* handler	: user defined function for handling response current endpoints
**/
func (s *Scope) GetMethod(path string, handler HandlerFunc) {
	s.generalMethod(http.MethodGet, path, handler)
}

/*PostMethod register handler for Post method
* path 	: endpoints path
* handler	: user defined function for handling response current endpoints
**/
func (s *Scope) PostMethod(path string, handler HandlerFunc) {
	s.generalMethod(http.MethodPost, path, handler)
}

/*generalMethod register handler for general methods used in internal in this package
* method	: request method to handle, e.g. http.MethodGet or http.MethodPost, etc
* path 		: endpoints path
* handler	: user defined function for handling response current endpoints
**/
func (s *Scope) generalMethod(method string, path string, handler HandlerFunc) {
	//handle response from http lib
	s.Server.HandleFunc(path, func(writer http.ResponseWriter, request *http.Request) {

		//intercept execution for checking correct http methodd
		if request.Method != method {

			//error response as json
			Data, _ := json.Marshal(&Response{
				ResponseCode: http.StatusMethodNotAllowed,
				Error:        &Error{Detail: "Status Method Not Allowed"},
			})

			//send response to the writer
			writer.Header().Set("Content-Type", "application/json")
			writer.WriteHeader(http.StatusMethodNotAllowed)
			writer.Write([]byte(Data))

		} else { //no error

			s.SetContext(writer, request)
			//execute user defined function
			response := handler(*s)

			//you can cek error from user defined handler and do something like custom log
			//if response.Error != nil {...}

			Data, _ := json.Marshal(response)
			writer.Header().Set("Content-Type", "application/json")
			writer.WriteHeader(response.ResponseCode)
			writer.Write([]byte(Data))

			//Print log
			log.Printf(RequestLogFormat, request.Host, request.URL.Path)
		}
	})
}

/*Serve for host (hostname/IP and port concatenate with ":")
* host 	: define your host for example localhost:8080
**/
func (s *Scope) Serve(host string) {
	log.Panic(http.ListenAndServe(host, s.Server))
}
