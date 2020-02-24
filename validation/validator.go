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
		predictors: make(map[string]*predictor),
	}
}

// Validator validates ...
type Validator struct {
	phasePointer int
	prevTSStates [][]tsmodels.TSState

	predictors map[string]*predictor
}

// Validate validates the new values with the profile and last known state
func (validator *Validator) Validate(profile tsmodels.TSProfile, utilValue map[string]interface{}) (float32, float32, int) {

	// handle utilValue and transform to TSState array
	//phaseThresholdLikeliness := profile.Settings.PhaseChangeLikeliness
	//phaseThresholdLikeliness := float32(0.90)
	tsinput := validator.extractTSInput(profile, utilValue)
	tsstates := validator.extractTSStates(profile, tsinput)

	// compute likeliness
	var likeliness float32
	if len(validator.prevTSStates) == 0 {
		likeliness = 1
	} else {
		likeliness = profile.Likeliness(validator.prevTSStates, tsstates)
		/*likeliness = profile.LikelinessPhase(validator.phasePointer, validator.prevTSStates, tsstates)
		if likeliness < phaseThresholdLikeliness {
			//fmt.Printf("find other phase than %d\n", validator.phasePointer)
			nextPhaseSteps := profile.Phases.Tx.Transitions[fmt.Sprintf("%d", validator.phasePointer)]
			for nextPhase, nextPhaseProb := range nextPhaseSteps.NextStateProbs {
				l := profile.LikelinessPhase(nextPhase, validator.prevTSStates, tsstates)
				l2 := l * (float32(nextPhaseProb) / float32(100))
				if l2 > likeliness {
					likeliness = l2
					validator.phasePointer = nextPhase
				}
			}
			//fmt.Printf("found new phase %d\n", validator.phasePointer)
		}
		*/
	}

	// compute prediction error
	_, exists := validator.predictors[profile.Name]
	if !exists {
		validator.predictors[profile.Name] = newPredictor(profile)
	}
	predictionError := validator.predictors[profile.Name].getError(tsstates, profile.Settings.States)

	// Update state history
	if len(validator.prevTSStates) >= profile.Settings.History {
		// remove first item from history
		validator.prevTSStates = validator.prevTSStates[1:]
	}
	validator.prevTSStates = append(validator.prevTSStates, tsstates)

	// update prediction to fit new history
	validator.predictors[profile.Name].predict(validator.phasePointer, validator.prevTSStates)

	return likeliness, predictionError, validator.phasePointer

}

func (validator *Validator) extractTSInput(profile tsmodels.TSProfile, utilValue map[string]interface{}) tsmodels.TSInput {
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
			Name:  metricname,
			Value: floatValue,
			// THIS MIN MAX MAY CHANGE depending on the metricTx!
			FixedMin: metricTx.Stats.Min,
			FixedMax: metricTx.Stats.Max,
		})
	}
	tsinput := tsmodels.TSInput{
		Metrics: metrics,
	}
	return tsinput
}

func (validator  *Validator) extractTSStates(profile tsmodels.TSProfile, tsinput tsmodels.TSInput) []tsmodels.TSState {
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
	return tsstates	
}