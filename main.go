package main

import (
	"log"

	"github.com/zoedaemon/pretest/routers"
	"github.com/zoedaemon/pretest/simplestorage"
)

func main() {
	log.Printf("Initialize Project...")

	//in Postman call it with localhost:8080 (without / ) to access root endpoint
	Host := "localhost:8080"

	//Create simple API
	SimpleAPI := routers.RegisterHandlers()

	//new simplestorage.Mapper
	DataMap := simplestorage.NewMapper()

	//save it to SimpleAPI so accessible from every handler
	SimpleAPI.SetCustomData(DataMap)

	//API Listen and Serve
	SimpleAPI.Serve(Host)
	log.Print("Listening On : ", Host)
}
