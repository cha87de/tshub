package datahub

import (
	"fmt"

	kvmtopmodels "github.com/cha87de/kvmtop/models"
)

// NewStreamer returns a new instance of Streamer
func NewStreamer(hub *Hub) *Streamer {
	return &Streamer{
		hub: hub,
	}
}

// Streamer handles the distribution of TSData monitoring items
type Streamer struct {
	hub *Hub
}

// Put transmits a new TSData monitoring item to the Streamer
func (streamer *Streamer) Put(tsdata kvmtopmodels.TSData) {
	fmt.Printf("streamer: %+v", tsdata)
}
