package main

import (
	"log"

	"github.com/cha87de/tshub/commport"
	"github.com/cha87de/tshub/datahub"
	"github.com/cha87de/tshub/restapi"
	"github.com/cha87de/tshub/restapi/operations"
	"github.com/go-openapi/loads"
)

func main() {
	// CONFIGS
	// TODO: read from Flags
	apiPort := 8080
	commType := "tcp"
	commAddr := "127.0.0.1:12345"

	// first: main tshub features
	dh := datahub.NewHub()
	commportServer := commport.NewServer(commType, commAddr, dh)
	go commportServer.Start()

	// second: swagger api
	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
	if err != nil {
		log.Fatalln(err)
	}
	api := operations.NewTshubAPI(swaggerSpec)
	configureAPI(api, dh)
	server := restapi.NewServer(api)
	server.ConfigureAPI()
	defer server.Shutdown()
	server.Port = apiPort
	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}
}
