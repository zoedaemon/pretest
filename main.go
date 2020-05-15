package main

import (
	"github.com/zoedaemon/pretest/simpleapi"

	"log"
)

//Root path
func root(scope simpleapi.Scope) *simpleapi.Response {
	return &simpleapi.Response{
		Error: nil,
		Data:  map[string]string{"version": "API v0.1"},
	}
}

//RegisterHandlers register all handlers API
func registerHandlers() {

	//create new SimpleAPI
	SimpleAPI := simpleapi.New()

	//register root path
	//http.HandleFunc("/", root)
	SimpleAPI.GetMethod("/", root)

	//call it with localhost:8080 (without / ) to access root endpoint
	Host := "localhost:8080"
	log.Print("Listening On : ", Host)

	//API Listen and Serve
	SimpleAPI.Serve(Host)
}

func main() {
	log.Printf("Initialize Project...")
	registerHandlers()
}
