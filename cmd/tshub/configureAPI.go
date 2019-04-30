package main

import (
	"strconv"

	"github.com/cha87de/tshub/datahub"
	"github.com/cha87de/tshub/models"
	"github.com/cha87de/tshub/restapi/operations"
	"github.com/cha87de/tshub/util"
	"github.com/cha87de/tsprofiler/predictor"
	"github.com/go-openapi/runtime/middleware"
)

func configureAPI(api *operations.TshubAPI, datahub *datahub.Hub) {

	// implementation for /hosts
	api.GetHostsHandler = operations.GetHostsHandlerFunc(func(params operations.GetHostsParams) middleware.Responder {
		resolution := 0 // TODO get from request
		hostnames := datahub.Store.GetTsNameWithField("host_name")
		hosts := make([]*models.Host, len(hostnames))
		i := 0
		for _, hostname := range hostnames {
			host := datahub.Store.GetTs(hostname, resolution)
			cores, _ := util.GetFloat(host.Latest("cpu_cores"))
			instances, _ := util.GetFloat(host.Latest("instances"))
			ram, _ := util.GetFloat(host.Latest("ram_Total"))
			hostModel := &models.Host{
				Name:          hostname,
				Cores:         int64(cores),
				Instancecount: int64(instances),
				RAM:           int64(ram),
			}
			hosts[i] = hostModel
			i++
		}
		return operations.NewGetHostsOK().WithPayload(hosts)
	})

	// implementation for /host/${hostname}
	api.GetHostHandler = operations.GetHostHandlerFunc(func(params operations.GetHostParams) middleware.Responder {
		hostname := params.Hostname
		hostDetails := &models.HostDetails{
			Name: hostname,
			Overbooking: &models.HostDetailsOverbooking{
				// TODO lookup values
				CPUCores: 0,
				CPUUtil:  0,
				DiskIO:   0,
				NetIO:    0,
				RAM:      0,
			},
		}
		return operations.NewGetHostOK().WithPayload(hostDetails)
	})

	// implementation for /host/{hostname}/plotdata
	api.GetHostPlotdataHandler = operations.GetHostPlotdataHandlerFunc(func(params operations.GetHostPlotdataParams) middleware.Responder {

		// return operations.NewGetDomainPlotdataInternalServerError()
		hostname := params.Hostname
		metric := params.Metric
		// timeframe := params.Timeframe

		timeframe, _ := strconv.Atoi(*params.Timeframe)
		hostTimeframeStore := datahub.Store.GetTs(hostname, timeframe)
		internalMetric := ""
		if metric == "cpu" {
			internalMetric = "cpu_total"
		} else if metric == "diskio" {
			internalMetric = "disk_io"
		} else if metric == "netio" {
			internalMetric = "net_io"
		}
		pastData := hostTimeframeStore.Dump(internalMetric)

		past := make([]*models.PlotDataItem, len(pastData))
		for i, data := range pastData {
			ts := int64(i)
			val, err := util.GetFloat(data)
			if err != nil {
				val = float64(0)
			}
			past[i] = &models.PlotDataItem{
				Timestamp: ts,
				Value:     val,
			}
		}

		plotdata := &models.PlotData{
			Metric: metric,
			Future: []*models.PlotDataItem{},
			Past:   past,
		}

		return operations.NewGetHostPlotdataOK().WithPayload(plotdata)

	})

	// implementaiton for /domains
	api.GetDomainsHandler = operations.GetDomainsHandlerFunc(func(params operations.GetDomainsParams) middleware.Responder {
		resolution := 0 // TODO get from request
		hostname := params.Hostname
		domainnames := datahub.Store.GetTsNameWithField("UUID")
		domains := make([]*models.Domain, 0)
		for _, domainname := range domainnames {
			domain := datahub.Store.GetTs(domainname, resolution)
			if hostname != nil && domain.Latest("host_name").(string) != *hostname {
				// skip, not requested ...
				continue
			}
			domainModel := &models.Domain{
				Cores: 4,
				RAM:   2,
				Name:  domainname,
			}
			domains = append(domains, domainModel)
		}
		return operations.NewGetDomainsOK().WithPayload(domains)
	})

	// implementaiton for /domain/{domainname}
	api.GetDomainHandler = operations.GetDomainHandlerFunc(func(params operations.GetDomainParams) middleware.Responder {
		domainname := params.Domainname
		domainDetails := &models.DomainDetails{
			Cores: 4,
			RAM:   2,
			Name:  domainname,
		}
		return operations.NewGetDomainOK().WithPayload(domainDetails)
	})

	// implementaiton for /domain/{domainname}/plotdata
	api.GetDomainPlotdataHandler = operations.GetDomainPlotdataHandlerFunc(func(params operations.GetDomainPlotdataParams) middleware.Responder {
		domainname := params.Domainname
		metric := params.Metric

		timeframe, _ := strconv.Atoi(*params.Timeframe)
		domainTimeframeStore := datahub.Store.GetTs(domainname, timeframe)

		internalMetric := ""
		if metric == "cpu" {
			internalMetric = "cpu_total"
		} else if metric == "diskio" {
			internalMetric = "disk_io"
		} else if metric == "netio" {
			internalMetric = "net_io"
		}
		pastData := domainTimeframeStore.Dump(internalMetric)

		past := make([]*models.PlotDataItem, len(pastData))
		for i, data := range pastData {
			ts := int64(i)
			val, err := util.GetFloat(data)
			if err != nil {
				val = float64(0)
			}
			past[i] = &models.PlotDataItem{
				Timestamp: ts,
				Value:     val,
			}
		}

		// compute future data
		profile, err := datahub.Store.GetProfile(domainname)
		if err != nil {
			return operations.NewGetDomainPlotdataNotFound()
		}
		predictor := predictor.NewPredictor(profile)
		/*predictor.SetState(map[string]string{
			"metric_0": "0",
		})*/
		simulation := predictor.Simulate(30)
		future := make([]*models.PlotDataItem, len(simulation))
		for i, simstep := range simulation {
			ts := int64(i)
			var data int64

			// find metric in simulation
			for _, x := range simstep {
				if x.Metric == internalMetric {
					data = x.State.Value
				}
			}

			val := float64(data)
			future[i] = &models.PlotDataItem{
				Timestamp: ts,
				Value:     val,
			}
		}

		plotdata := &models.PlotData{
			Metric: metric,
			Future: future,
			Past:   past,
		}

		return operations.NewGetDomainPlotdataOK().WithPayload(plotdata)
	})

	// implementation for /profiles
	api.GetProfileNamesHandler = operations.GetProfileNamesHandlerFunc(func(params operations.GetProfileNamesParams) middleware.Responder {
		names := datahub.Store.GetProfileNames()
		return operations.NewGetProfileNamesOK().WithPayload(names)
	})

	// implementation for /profile/${profilename}
	api.GetProfileHandler = operations.GetProfileHandlerFunc(func(params operations.GetProfileParams) middleware.Responder {
		name := params.Profilename
		profile, err := datahub.Store.GetProfile(name)
		if err != nil {
			return operations.NewGetProfileNotFound()
		}
		return operations.NewGetProfileOK().WithPayload(profile)
	})

}
