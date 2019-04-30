package datahub

import (
	"sync"

	"github.com/cha87de/tshub/tsdb"
)

func newStoreTs() *storeTs {
	return &storeTs{
		tsdata:       make(map[string][]*tsdb.Store),
		tsdataAccess: sync.Mutex{},
	}
}

// StoreTs is as part of Store responsible for storing time series data
type storeTs struct {
	tsdata       map[string][]*tsdb.Store
	tsdataAccess sync.Mutex
}

// KeepTs stores a time series data item
func (storeTs *storeTs) KeepTs(name string, values map[string]interface{}) {
	storeTs.tsdataAccess.Lock()
	defer storeTs.tsdataAccess.Unlock()

	if _, exists := storeTs.tsdata[name]; !exists {
		storeTs.tsdata[name] = make([]*tsdb.Store, len(resolutions))
		for i, resolution := range resolutions {
			storeTs.tsdata[name][i] = tsdb.NewStore(30, resolution)
		}
	}

	for i := range resolutions {
		storeTs.tsdata[name][i].Add(values)
	}
}

func (storeTs *storeTs) GetTs(name string, resolution int) *tsdb.Store {
	return storeTs.tsdata[name][resolution]
}

func (storeTs *storeTs) GetTsNameWithField(fieldname string) []string {
	result := make([]string, 0)
	defaultResolution := 0

	storeTs.tsdataAccess.Lock()
	for name, stores := range storeTs.tsdata {
		store := stores[defaultResolution]
		field := store.Latest(fieldname)
		if field != nil {
			result = append(result, name)
		}
	}
	storeTs.tsdataAccess.Unlock()

	return result
}
