package main

import (
	"log"

	"github.com/cha87de/tshub/commport"
	"github.com/cha87de/tshub/datahub"
	"github.com/cha87de/tshub/restapi"
	"github.com/cha87de/tshub/restapi/operations"
	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime/middleware"
)

func main() {
	// CONFIGS
	// TODO: read from Flags
	apiPort := 8080
	commType := "tcp"
	commAddr := "127.0.0.1:12345"

	// first: main tshub features
	datahub := datahub.NewHub()
	commportServer := commport.NewServer(commType, commAddr, datahub)
	go commportServer.Start()

	// second: swagger api
	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
	if err != nil {
		log.Fatalln(err)
	}
	api := operations.NewTshubAPI(swaggerSpec)
	api.GetProfileNamesHandler = operations.GetProfileNamesHandlerFunc(func(params operations.GetProfileNamesParams) middleware.Responder {
		names := datahub.Store.GetNames()
		return operations.NewGetProfileNamesOK().WithPayload(names)
	})
	api.GetProfileHandler = operations.GetProfileHandlerFunc(func(params operations.GetProfileParams) middleware.Responder {
		name := params.Profilename
		profile, err := datahub.Store.GetByName(name)
		if err != nil {
			return operations.NewGetProfileNotFound()
		}
		return operations.NewGetProfileOK().WithPayload(profile)
	})
	server := restapi.NewServer(api)
	defer server.Shutdown()
	server.Port = apiPort
	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}
}
