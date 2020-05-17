package simpleapi

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type (

	//HandlerFunc SimpleAPI implementation
	HandlerFunc func(Scope) *Response

	//HandlerWebsocketFunc SimpleAPI implementation + gorilla/websocket package
	HandlerWebsocketFunc func(Scope, *websocket.Conn) ([]byte, error)

	//Scope SimpleAPI all attributes / field
	Scope struct {
		//hold server http
		Server *http.ServeMux

		//derivate Context Class
		Context

		//e.g. map objects of files config, database connection, redis, etc
		customData interface{}

		Upgrader websocket.Upgrader
	}

	//Response that should be pass from user defined function HandlerFunc
	Response struct {
		ResponseCode int         `json:"response-code"`
		Error        error       `json:"error"`
		Data         interface{} `json:"data"`
	}
)

//RequestLogFormat default format for log the request
const RequestLogFormat = "\n\tHost : %s \n\tPath : %s\n\tResponse: %s"

//New SimpleAPI
func New() *Scope {
	return &Scope{
		Server: http.NewServeMux(),
		Upgrader: websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
			CheckOrigin: func(r *http.Request) bool {
				//allow ws connection, CONCERN: use wss if use domain name, don't use this option
				return true
			},
		},
	}
}

/*GetMethod register handler for Get method
* path 	: endpoints path
* handler	: user defined function for handling response current endpoints
**/
func (s *Scope) GetMethod(path string, handler HandlerFunc) {
	generalMethod(s, http.MethodGet, path, handler)
}

/*PostMethod register handler for Post method
* path 	: endpoints path
* handler	: user defined function for handling response current endpoints
**/
func (s *Scope) PostMethod(path string, handler HandlerFunc) {
	generalMethod(s, http.MethodPost, path, handler)
}

/*generalMethod register handler for general methods used in internal in this package
* method	: request method to handle, e.g. http.MethodGet or http.MethodPost, etc
* path 		: endpoints path
* handler	: user defined function for handling response current endpoints
**/
func generalMethod(s *Scope, method string, path string, handler HandlerFunc) {
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
			log.Printf(RequestLogFormat, request.Host, request.URL.Path, Data)
		}
	})
}

/*WebsocketMethod register handler for general methods used in internal in this package
* path 		: endpoints path
* handler	: user defined function for handling response current endpoints
**/
func (s *Scope) WebsocketMethod(path string, handler HandlerWebsocketFunc) {
	//handle response from http lib
	s.Server.HandleFunc(path, func(writer http.ResponseWriter, request *http.Request) {

		s.SetContext(writer, request)

		//upgrate to websocket protocol
		ws, err := s.Upgrader.Upgrade(writer, request, nil)
		if err != nil {
			log.Printf("error : " + err.Error())
			return
		}
		defer func() {
			ws.Close()
		}()

		//Print log
		log.Printf("Websocket connection started...")

		for {
			//execute user defined function for websocket
			msg, err := handler(*s, ws)

			//stop it if return value exist (error happen might be)
			if err != nil {
				break
			}

			//Print log
			log.Printf("some message has been receive : " + string(msg))
		}
		//Print log
		log.Printf("Websocket disconnected...bye...")

	})
}

/*Serve for host (hostname/IP and port concatenate with ":")
* host 	: define your host for example localhost:8080
**/
func (s *Scope) Serve(host string) {
	log.Panic(http.ListenAndServe(host, s.Server))
}

/*SetCustomData for user defined data that can be accessed from any handler
	that defined with GetMethod and PostMethod
* host 	: define your host for example localhost:8080
**/
func (s *Scope) SetCustomData(data interface{}) {
	s.customData = data
}

//GetCustomData get user defined data
func (s *Scope) GetCustomData() interface{} {
	return s.customData
}
