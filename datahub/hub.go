package datahub

// NewHub returns a new Hub instance
func NewHub() *Hub {
	hub := &Hub{}
	hub.Store = *NewStore(hub)
	hub.Streamer = *NewStreamer(hub)
	return hub
}

// Hub is the central data hub for storing or streaming
type Hub struct {
	Store    Store
	Streamer Streamer
}
