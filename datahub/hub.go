package datahub

import "github.com/cha87de/tshub/validation"

// NewHub returns a new Hub instance
func NewHub() *Hub {
	hub := &Hub{}
	hub.Store = *NewStore(hub)
	hub.Streamer = *NewStreamer(hub)
	hub.Validator = *validation.NewValidator()
	return hub
}

// Hub is the central data hub for storing or streaming
type Hub struct {
	Store     Store
	Streamer  Streamer
	Validator validation.Validator
}
