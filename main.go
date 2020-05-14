package main

import (
	"github.com/zoedaemon/pretest/simpleapi"

	"fmt"
	"log"
	"net/http"
)

//RequestLogFormat default format for log the request
const RequestLogFormat = "\n\tRequest : %s \n\tPath : %s"

//Root path
func root(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "API v0.1")
	log.Printf(RequestLogFormat, request.Host, request.URL.Path)
}

//RegisterHandlers register all handlers API
func registerHandlers() {

	SimpleAPI := simpleapi.New()

	//register root path
	//http.HandleFunc("/", root)
	SimpleAPI.GetMethod('/', root)

	//call it with localhost:8080 (without / ) to access root endpoint
	Host := "localhost:8080"
	log.Print("Listening On : ", Host)

	//API Listen and Serve
	log.Panic(http.ListenAndServe(Host, nil))
}

func main() {
	log.Printf("Initialize Project...")
	registerHandlers()
}
