package datahub

import (
	"fmt"
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

// Put transmits a new TSInput monitoring item to the Streamer
func (streamer *Streamer) Put(name string, input map[string]interface{}) float32 {
	streamer.hub.Store.KeepTs(name, input)

	profile, err := streamer.hub.Store.GetProfile(name)
	if err != nil {
		fmt.Printf("failed to get profile %s: %s", name, err)
		return 0
	}
	likeliness := streamer.hub.Validator.Validate(profile, input)
	return likeliness
}
