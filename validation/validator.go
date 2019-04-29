package validation

import (
	"fmt"

	"github.com/cha87de/tshub/util"
	tsmodels "github.com/cha87de/tsprofiler/models"
	profiler "github.com/cha87de/tsprofiler/profiler"
	"github.com/cha87de/tsprofiler/profiler/buffer"
	"github.com/cha87de/tsprofiler/profiler/discretizer"
)

// NewValidator creates and returns a new Validator
func NewValidator() *Validator {
	return &Validator{
		phasePointer: 0,
	}
}

// Validator validates ...
type Validator struct {
	phasePointer int
	prevTSStates [][]tsmodels.TSState
}

// Validate validates the new values with the profile and last known state
func (validator *Validator) Validate(profile tsmodels.TSProfile, utilValue map[string]interface{}) float32 {

	metrics := make([]tsmodels.TSInputMetric, 0)
	for metricname, value := range utilValue {
		floatValue, err := util.GetFloat(value)
		if err != nil {
			continue
		}

		var txs []tsmodels.TxMatrix
		txs = profile.PeriodTree.Root.TxMatrix
		var metricTx tsmodels.TxMatrix
		found := false
		for _, tx := range txs {
			if tx.Metric == metricname {
				metricTx = tx
				found = true
				break
			}
		}
		if !found {
			fmt.Printf("cannot find metric %s in phases.\n", metricname)
			continue
		}

		metrics = append(metrics, tsmodels.TSInputMetric{
			Name:     metricname,
			Value:    floatValue,
			FixedMin: metricTx.Stats.Min,
			FixedMax: metricTx.Stats.Max,
		})
	}
	tsinput := tsmodels.TSInput{
		Metrics: metrics,
	}

	tsprofiler := profiler.NewProfiler(tsmodels.Settings{
		Name: "validator",

		BufferSize:    profile.Settings.BufferSize,
		States:        profile.Settings.States,
		FilterStdDevs: profile.Settings.FilterStdDevs,
		History:       profile.Settings.History,
		FixBound:      profile.Settings.FixBound,
		PeriodSize:    profile.Settings.PeriodSize,
	})

	buffer := buffer.NewBuffer(profile.Settings.FilterStdDevs, tsprofiler)
	buffer.Add(tsinput)
	tsbuffers := buffer.Reset()

	discretizer := discretizer.NewDiscretizer(profile.Settings.States, profile.Settings.FixBound, tsprofiler)
	tsstates := discretizer.Discretize(tsbuffers)

	var likeliness float32
	if len(validator.prevTSStates) == 0 {
		likeliness = 1
	} else {
		//likeliness = profile.LikelinessPhase(validator.phasePointer, validator.prevTSStates, tsstates)
		likeliness = profile.Likeliness(validator.prevTSStates, tsstates)
	}
	if len(validator.prevTSStates) >= profile.Settings.History {
		// remove first item from history
		validator.prevTSStates = validator.prevTSStates[1:]
	}
	validator.prevTSStates = append(validator.prevTSStates, tsstates)

	// fmt.Printf("value %.0f (state %v) likeliness %.2f\n", tsbuffers[0].RawData[0], tsstates[0].State.Value, likeliness)

	return likeliness

}
