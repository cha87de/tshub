package main

import (
	"github.com/cha87de/tshub/commport"
	"github.com/cha87de/tshub/datahub"
)

func main() {

	datahub := datahub.NewHub()

	commportServer := commport.NewServer("tcp", "127.0.0.1:12345", datahub)
	commportServer.Start()

}
