package datahub

import (
	"fmt"
	"sync"

	tsmodels "github.com/cha87de/tsprofiler/models"
)

// TODO cleanup of "old" profiles, which don't appear anymore

func newStoreProfile() *storeProfile {
	return &storeProfile{
		tsprofiles:       make(map[string]tsmodels.TSProfile),
		tsprofilesAccess: sync.Mutex{},
	}
}

// storeProfile is as part of Store responsible for TSProfile storage
type storeProfile struct {
	// storage for TSProfiles
	tsprofiles       map[string]tsmodels.TSProfile
	tsprofilesAccess sync.Mutex
}

// KeepProfile stores a TSProfile
func (storeProfile *storeProfile) KeepProfile(tsprofile tsmodels.TSProfile) {
	storeProfile.tsprofilesAccess.Lock()
	storeProfile.tsprofiles[tsprofile.Name] = tsprofile
	storeProfile.tsprofilesAccess.Unlock()
}

// GetProfile returns the stored TSProfile with the given name
func (storeProfile *storeProfile) GetProfile(name string) (tsmodels.TSProfile, error) {
	storeProfile.tsprofilesAccess.Lock()
	defer storeProfile.tsprofilesAccess.Unlock()
	profile, found := storeProfile.tsprofiles[name]
	if !found {
		err := fmt.Errorf("No profile found with name %s", name)
		return tsmodels.TSProfile{}, err
	}
	return profile, nil
}

// GetProfileNames returns the names of available TSProfiles
func (storeProfile *storeProfile) GetProfileNames() []string {
	storeProfile.tsprofilesAccess.Lock()
	defer storeProfile.tsprofilesAccess.Unlock()
	var names []string
	for n := range storeProfile.tsprofiles {
		names = append(names, n)
	}
	return names
}
