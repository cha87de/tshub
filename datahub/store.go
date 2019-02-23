package datahub

import (
	"sync"

	tsmodels "github.com/cha87de/tsprofiler/models"
)

// NewStore returns a new Store instance
func NewStore(hub *Hub) *Store {
	return &Store{
		hub:              hub,
		tsprofiles:       make(map[string]tsmodels.TSProfile),
		tsprofilesAccess: sync.Mutex{},
	}
}

// Store is the central sink for TSProfiles
type Store struct {
	hub              *Hub
	tsprofiles       map[string]tsmodels.TSProfile
	tsprofilesAccess sync.Mutex
}

// Keep stores a TSProfile
func (store *Store) Keep(tsprofile tsmodels.TSProfile) {
	store.tsprofilesAccess.Lock()
	store.tsprofiles[tsprofile.Name] = tsprofile
	store.tsprofilesAccess.Unlock()
}

// GetByName returns the stored TSProfile with the given name
func (store *Store) GetByName(name string) tsmodels.TSProfile {
	store.tsprofilesAccess.Lock()
	defer store.tsprofilesAccess.Unlock()
	return store.tsprofiles[name]
}
