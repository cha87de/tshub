package datahub

// NewStore returns a new Store instance
func NewStore(hub *Hub) *Store {
	store := &Store{
		hub: hub,
	}
	store.storeProfile = newStoreProfile()
	store.storeTs = newStoreTs()

	return store
}

// Store is the central sink for TSProfiles
type Store struct {
	hub *Hub

	*storeProfile
	*storeTs
}
