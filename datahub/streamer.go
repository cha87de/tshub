package datahub

import (
	"fmt"

	"github.com/cha87de/tsprofiler/models"
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
func (streamer *Streamer) Put(name string, tsinput models.TSInput) float32 {
	values := make(map[string]interface{})
	for _, metric := range tsinput.Metrics {
		values[metric.Name] = metric.Value
	}
	streamer.hub.Store.KeepTs(name, values)

	profile, err := streamer.hub.Store.GetProfile(name)
	if err != nil {
		fmt.Printf("failed to get profile %s: %s", name, err)
		return 0
	}
	likeliness := streamer.hub.Validator.Validate(profile, values)
	return likeliness
}
