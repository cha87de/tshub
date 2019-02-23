package datahub

import (
	"fmt"
	"sync"

	tsmodels "github.com/cha87de/tsprofiler/models"
)

// TODO cleanup of "old" profiles, which don't appear anymore

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
func (store *Store) GetByName(name string) (tsmodels.TSProfile, error) {
	store.tsprofilesAccess.Lock()
	defer store.tsprofilesAccess.Unlock()
	profile, found := store.tsprofiles[name]
	if !found {
		err := fmt.Errorf("No profile found with name %s", name)
		return tsmodels.TSProfile{}, err
	}
	return profile, nil
}

// GetNames returns the names of available TSProfiles
func (store *Store) GetNames() []string {
	store.tsprofilesAccess.Lock()
	defer store.tsprofilesAccess.Unlock()
	var names []string
	for n := range store.tsprofiles {
		names = append(names, n)
	}
	return names
}
