package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/cha87de/tshub/datahub"
	"github.com/cha87de/tsprofiler/models"
	"github.com/cha87de/tsprofiler/utils"
	flags "github.com/jessevdk/go-flags"
)

var options struct {
	ProfileFile         string `long:"in.profile" default:"profile.json"`
	TSFile              string `long:"in.ts" default:"tsinput.csv"`
	ProbFile            string `long:"out.probs" default:""`
	PhasesFile          string `long:"out.phases" default:""`
	PredictionErrorFile string `long:"out.predictionerror" default:""`
}

var profileName = "unknownProfile"
var probfile *os.File
var phasesfile *os.File
var predictionErrorFile *os.File

func main() {
	initializeFlags()

	tsprofile := utils.ReadProfileFromFile(options.ProfileFile)
	profileName = tsprofile.Name

	// create and initialize the TSHub
	datahub := initTshub(tsprofile)

	// create & open output file
	if options.ProbFile != "" && options.ProbFile != "-" {
		var err error
		probfile, err = os.OpenFile(options.ProbFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatal(err)
		}
		defer probfile.Close()
	}
	// create & open output file
	if options.PhasesFile != "" && options.PhasesFile != "-" {
		var err error
		phasesfile, err = os.OpenFile(options.PhasesFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatal(err)
		}
		defer phasesfile.Close()
	}
	// create & open output file
	if options.PredictionErrorFile != "" && options.PredictionErrorFile != "-" {
		var err error
		predictionErrorFile, err = os.OpenFile(options.PredictionErrorFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatal(err)
		}
		defer predictionErrorFile.Close()
	}

	// read file line by line
	processTSFile(datahub, options.TSFile)
}

func initializeFlags() {
	// initialize parser for flags
	parser := flags.NewParser(&options, flags.Default)
	parser.ShortDescription = "tshub-fs"
	parser.LongDescription = "Reads a TSProfile and raw time series from file system and compares the TS stream with the TSProfile."

	// Parse parameters
	if _, err := parser.Parse(); err != nil {
		code := 1
		if fe, ok := err.(*flags.Error); ok {
			if fe.Type == flags.ErrHelp {
				code = 0
			}
		}
		if code != 0 {
			fmt.Printf("Error parsing flags: %s\n", err)
		}
		os.Exit(code)
	}
}

func initTshub(tsprofile models.TSProfile) *datahub.Hub {
	datahub := datahub.NewHub()
	datahub.Store.KeepProfile(tsprofile)
	return datahub
}

func processTSFile(datahub *datahub.Hub, filename string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	i := 0
	for {
		i++
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		var utilValues []string
		for _, rawValue := range record {
			//utilValue, err := strconv.ParseFloat(rawValue, 64)
			//if err != nil {
			//	continue
			//}
			utilValues = append(utilValues, rawValue)
		}
		putMeasurement(datahub, utilValues)
	}
	fmt.Printf("processed %d lines\n", i)
	likeliness := likelinessSum / float64(likelinessCount)
	predictionError := predictionerrorSum / float64(predictionerrorCount)
	fmt.Printf("likeliness of CSV to profile: %.6f\n", likeliness)
	fmt.Printf("prediction error of CSV to profile: %.6f\n", predictionError)

}

var likelinessSum float64
var likelinessCount int32
var predictionerrorSum float64
var predictionerrorCount int32

func putMeasurement(datahub *datahub.Hub, utilValue []string) {
	metrics := make(map[string]interface{})
	for i, value := range utilValue {
		metricName := fmt.Sprintf("metric_%d", i)
		metrics[metricName] = value
	}
	likeliness, predictionError, phaseid := datahub.Streamer.Put(profileName, metrics)
	likelinessSum += float64(likeliness)
	likelinessCount++
	predictionerrorSum += float64(predictionError)
	predictionerrorCount++

	var row string

	// handle likeliness output
	row = fmt.Sprintf("%.6f\n", likeliness)
	if options.ProbFile == "-" {
		// use stdout
		fmt.Print(row)
	} else if options.ProbFile == "" {
		// ignore likeliness
	} else {
		// print to file
		if _, err := probfile.Write([]byte(row)); err != nil {
			log.Fatal(err)
		}
	}

	// handle prediction error output
	row = fmt.Sprintf("%.6f\n", predictionError)
	if options.PredictionErrorFile == "-" {
		// use stdout
		fmt.Print(row)
	} else if options.PredictionErrorFile == "" {
		// ignore likeliness
	} else {
		// print to file
		if _, err := predictionErrorFile.Write([]byte(row)); err != nil {
			log.Fatal(err)
		}
	}

	// handle phases output
	row = fmt.Sprintf("%d\n", phaseid)
	if options.PhasesFile == "-" {
		// use stdout
		fmt.Print(row)
	} else if options.PhasesFile == "" {
		// ignore likeliness
	} else {
		// print to file
		if _, err := phasesfile.Write([]byte(row)); err != nil {
			log.Fatal(err)
		}
	}

}
