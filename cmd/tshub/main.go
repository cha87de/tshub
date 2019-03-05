package main

import (
	"fmt"
	"log"
	"os"

	"github.com/cha87de/tshub/commport"
	"github.com/cha87de/tshub/datahub"
	"github.com/cha87de/tshub/restapi"
	"github.com/cha87de/tshub/restapi/operations"
	"github.com/go-openapi/loads"
	flags "github.com/jessevdk/go-flags"
)

var options struct {
	APIPort  int    `long:"apiport" default:"8080"`
	CommType string `long:"comm.type" default:"tcp"`
	CommAddr string `long:"comm.addr" default:"127.0.0.1:12345"`
}

func main() {
	// CONFIGS
	initializeFlags()

	// first: main tshub features
	dh := datahub.NewHub()
	commportServer := commport.NewServer(options.CommType, options.CommAddr, dh)
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
	server.Port = options.APIPort
	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}
}

func initializeFlags() {
	// initialize parser for flags
	parser := flags.NewParser(&options, flags.Default)
	parser.ShortDescription = "tshub"
	parser.LongDescription = "Endpoint for kvmtop and kvmprofiler"
	parser.ArgsRequired = false

	// Parse parameters
	if _, err := parser.Parse(); err != nil {
		code := 1
		if fe, ok := err.(*flags.Error); ok {
			if fe.Type == flags.ErrHelp {
				code = 0
			}
		}
		if code != 0 {
			fmt.Printf("Error parsing flags: %s", err)
		}
		os.Exit(code)
	}

}
