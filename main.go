package main

import (
	"log"

	"github.com/zoedaemon/pretest/routers"
)

func main() {
	log.Printf("Initialize Project...")

	//in Postman call it with localhost:8080 (without / ) to access root endpoint
	Host := "localhost:8080"

	//Create simple API
	SimpleAPI := routers.RegisterHandlers()

	//API Listen and Serve
	SimpleAPI.Serve(Host)
	log.Print("Listening On : ", Host)
}
