package tsdb

import (
	"fmt"
	"time"
)

// NewStore instantiates and returns a new tsdb Store
func NewStore(historySize int, resolution time.Duration) *Store {
	minPerSlot := resolution.Seconds() / float64(historySize)
	dur := time.Duration(minPerSlot) * time.Second
	return &Store{
		HistorySize: historySize,
		Resolution:  resolution,
		items:       make([]Item, 0),

		durationPerSlot: dur,
	}
}

// Store holds time series data with maximum number of `HistorySize` and groups
// items by `Resolution` when adding values.
type Store struct {
	HistorySize     int
	Resolution      time.Duration
	durationPerSlot time.Duration

	items []Item
}

// Add adds the given values to the store items, with respect to the configured
// resolution, values may be merged using average with most recent values in
// store.
func (store *Store) Add(values map[string]interface{}) {
	item := NewItem(values)
	// check if latest value is old enough to rotate items
	store.rotate(*item)
	// add given new value to the last items
	store.append(*item)
}

func (store *Store) rotate(item Item) {
	emptyValues := make(map[string]interface{})
	if len(store.items) == 0 {
		fmt.Printf("add first empty item ... \n")
		store.items = append(store.items, *NewItem(emptyValues))
		return
	}
	lastItem := store.items[len(store.items)-1]
	rotateTime := lastItem.Timestamp.Add(store.durationPerSlot)
	fmt.Printf("rotate at %+v\n", rotateTime)
	if rotateTime.Before(item.Timestamp) {
		fmt.Printf("time to rotate ... \n")
		if len(store.items) < store.HistorySize {
			// still growing ... append new item
			store.items = append(store.items, *NewItem(emptyValues))
		} else {
			// time to rotate :-)
			store.items = store.items[1:]
		}

	}
}

func (store *Store) append(item Item) {
	// merge into last one
	store.items[len(store.items)-1].merge(item)
}

// Latest returns the latest value for given key
func (store *Store) Latest(key string) interface{} {
	val, ok := store.items[len(store.items)-1].Values[key]
	if ok {
		return val
	}
	return nil
}

// Dump returns an array of all stored measurements for given key
func (store *Store) Dump(key string) []interface{} {
	dump := make([]interface{}, len(store.items))
	for i, values := range store.items {
		measurement, ok := values.Values[key]
		if !ok {
			dump[i] = nil
			continue
		}
		dump[i] = measurement
	}
	return dump
}
