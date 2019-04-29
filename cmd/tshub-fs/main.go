package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"

	"github.com/cha87de/tshub/datahub"
	"github.com/cha87de/tsprofiler/models"
	"github.com/cha87de/tsprofiler/utils"
	flags "github.com/jessevdk/go-flags"
)

var options struct {
	ProfileFile string `long:"in.profile" default:"profile.json"`
	TSFile      string `long:"in.ts" default:"tsinput.csv"`
}

var profileName = "unknownProfile"

func main() {
	initializeFlags()

	tsprofile := utils.ReadProfileFromFile(options.ProfileFile)
	profileName = tsprofile.Name

	// create and initialize the TSHub
	datahub := initTshub(tsprofile)

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

		var utilValues []float64
		for _, rawValue := range record {
			utilValue, err := strconv.ParseFloat(rawValue, 64)
			if err != nil {
				continue
			}
			utilValues = append(utilValues, utilValue)
		}
		putMeasurement(datahub, utilValues)
	}
	fmt.Printf("processed %d lines\n", i)
	likeliness := likelinessSum / float64(likelinessCount)
	fmt.Printf("likeliness of CSV to profile: %.6f", likeliness)
}

var likelinessSum float64
var likelinessCount int32

func putMeasurement(datahub *datahub.Hub, utilValue []float64) {
	metrics := make([]models.TSInputMetric, 0)
	for i, value := range utilValue {
		metrics = append(metrics, models.TSInputMetric{
			Name:  fmt.Sprintf("metric_%d", i),
			Value: value,
			//FixedMin: options.FixedMin,
			//FixedMax: options.FixedMax,
		})
	}
	tsinput := models.TSInput{
		Metrics: metrics,
	}
	likeliness := datahub.Streamer.Put(profileName, tsinput)
	likelinessSum += float64(likeliness)
	likelinessCount++

	//data := datahub.Store.GetTs("infile", 0).Dump("metric_0")
	//fmt.Printf("%f\n", data)

	// TODO extend validator / predictor

}
