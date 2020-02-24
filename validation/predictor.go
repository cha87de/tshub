package validation

import (
	"fmt"
	"math"

	tsmodels "github.com/cha87de/tsprofiler/models"
	tspredictor "github.com/cha87de/tsprofiler/predictor"
)

func newPredictor(profile tsmodels.TSProfile) *predictor {
	predictor := &predictor{}

	predictor.tspredictor = tspredictor.NewPredictor(profile)
	//x.SetMode()
	return predictor
}

type predictor struct {
	predictions []tsmodels.TSState
	tspredictor *tspredictor.Predictor
}

func (predictor *predictor) getError(tsstates []tsmodels.TSState, maxstates int) float32 {
	errorTotalSum := float64(0)
	errorTotalCount := 0
	// for each metric
	for _, metricState := range tsstates {
		actualState := float64(metricState.State.Value)

		var prediction tsmodels.TSState
		found := false
		for _, state := range predictor.predictions {
			if state.Metric == metricState.Metric {
				prediction = state
				found = true
				break
			}
		}
		if !found {
			// handle as highest possible error (1)
			errorTotalSum++
			errorTotalCount++
			continue
		}
		predictedState := float64(prediction.State.Value)
		diff := math.Abs(actualState - predictedState)
		ratio := diff / float64(maxstates)
		errorTotalSum += ratio
		errorTotalCount++
	}
	return float32(errorTotalSum / float64(errorTotalCount))
}

func (predictor *predictor) predict(phasePointer int, prevTSStates [][]tsmodels.TSState) {
	prevTSStatesStr := make(map[string]string)
	for _, metricStates := range prevTSStates {
		metricStateHistory := ""
		for _, state := range metricStates {
			if metricStateHistory != "" {
				metricStateHistory = metricStateHistory + "-"
			}
			metricStateHistory = metricStateHistory + fmt.Sprintf("%d", state.State.Value)
		}
		prevTSStatesStr[metricStates[0].Metric] = metricStateHistory
	}
	predictor.tspredictor.SetState(prevTSStatesStr)
	predictor.tspredictor.SetPhase(phasePointer)
	simulation := predictor.tspredictor.SimulateSteps(1)
	predictor.predictions = simulation[0]
}
