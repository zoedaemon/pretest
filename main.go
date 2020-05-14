package main

import (
	"fmt"
	"log"
	"net/http"
)

//Root path
func root(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "API v0.1")
	log.Printf("Request : %s \nPath : %s", request.Host, request.URL.Path)
}

//RegisterHandlers register all handlers API
func registerHandlers() {

	//register root path
	http.HandleFunc("/", root)

	//call it with localhost:8080 (without / ) to access root endpoint
	Host := "localhost:8080"
	log.Print("Listening On : ", Host)
	log.Panic(http.ListenAndServe(Host, nil))
}

func main() {
	log.Printf("Initialize Project...")
	registerHandlers()
}
