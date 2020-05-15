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
		Error error
		Data  interface{}
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

			writer.WriteHeader(http.StatusMethodNotAllowed)
			writer.Write([]byte("{\"error\":\"Status Method Not Allowed\"}"))

		} else { //no error

			//execute user defined function
			response := handler(*s)

			//cek if no error from user defined handler
			if response.Error == nil {

				//json format data
				MapData := map[string]interface{}{"error": nil, "data": response.Data}
				Data, _ := json.Marshal(MapData)

				//send response to the writer
				writer.WriteHeader(http.StatusOK)
				writer.Write([]byte(Data))

			} //else... TODO error handling

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
