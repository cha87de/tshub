package tsdb

import (
	"time"

	"github.com/cha87de/tshub/util"
	"gonum.org/v1/gonum/stat"
)

// NewItem returns a new Item with given values and current time
func NewItem(values map[string]interface{}) *Item {
	return &Item{
		Timestamp: time.Now(),
		Values:    values,
		counts:    make(map[string]int64),
	}
}

// Item represents a set of metric values at a timestamp
type Item struct {
	Timestamp time.Time
	Values    map[string]interface{}
	counts    map[string]int64
}

func (item *Item) merge(newItem Item) {
	for metric, value := range newItem.Values {
		valueFloat, err := util.GetFloat(value)
		if err != nil {
			// if not a number, simply set it.
			item.Values[metric] = value
			continue
		}

		// if a number:
		if oldVal, exists := item.Values[metric]; !exists {
			// no previous value availabe, simply set it
			item.Values[metric] = value
		} else {
			// previous value exists, build average
			oldValueFloat, err := util.GetFloat(oldVal)
			if err != nil {
				// this should never happen!
				continue
			}
			oldValueCount, ok := item.counts[metric]
			if !ok {
				oldValueCount = 0
			}
			newAvg := stat.Mean(
				[]float64{oldValueFloat, valueFloat},
				[]float64{float64(oldValueCount), 1},
			)
			newCount := oldValueCount + 1

			// store back new values
			item.Values[metric] = newAvg
			item.counts[metric] = newCount
		}
	}
	// finally update timestamp
	// item.Timestamp = newItem.Timestamp
}
