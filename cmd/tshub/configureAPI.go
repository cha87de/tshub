package main

import (
	"github.com/cha87de/tshub/datahub"
	"github.com/cha87de/tshub/models"
	"github.com/cha87de/tshub/restapi/operations"
	"github.com/cha87de/tshub/util"
	"github.com/go-openapi/runtime/middleware"
)

func configureAPI(api *operations.TshubAPI, datahub *datahub.Hub) {

	// implementation for /hosts
	api.GetHostsHandler = operations.GetHostsHandlerFunc(func(params operations.GetHostsParams) middleware.Responder {
		hosts := make([]*models.Host, len(datahub.Streamer.Hosts))
		i := 0
		for hostname, host := range datahub.Streamer.Hosts {
			hostModel := &models.Host{
				Name:          hostname,
				Cores:         int64(host[0].Latest("cpu_cores").(float64)),
				Instancecount: int64(host[0].Latest("instances").(int)),
				RAM:           int64(host[0].Latest("ram_Total").(float64)),
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

	// implementaiton for /host/{hostname}/plotdata
	api.GetHostPlotdataHandler = operations.GetHostPlotdataHandlerFunc(func(params operations.GetHostPlotdataParams) middleware.Responder {
		return operations.NewGetHostPlotdataInternalServerError()
	})

	// implementaiton for /domains
	api.GetDomainsHandler = operations.GetDomainsHandlerFunc(func(params operations.GetDomainsParams) middleware.Responder {
		hostname := params.Hostname
		domains := make([]*models.Domain, 0)
		for domainname, domain := range datahub.Streamer.Domains {
			if hostname != nil && domain[0].Latest("host_name").(string) != *hostname {
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
		// return operations.NewGetDomainPlotdataInternalServerError()
		domainname := params.Domainname
		metric := params.Metric
		// timeframe := params.Timeframe

		domainStores, ok := datahub.Streamer.Domains[domainname]
		if !ok {
			return operations.NewGetDomainPlotdataNotFound()
		}
		domainTimeframeStore := domainStores[0] // TODO align with timeframe param
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

		plotdata := &models.PlotData{
			Metric: metric,
			Future: []*models.PlotDataItem{},
			Past:   past,
		}

		return operations.NewGetDomainPlotdataOK().WithPayload(plotdata)

	})

	// implementation for /profiles
	api.GetProfileNamesHandler = operations.GetProfileNamesHandlerFunc(func(params operations.GetProfileNamesParams) middleware.Responder {
		names := datahub.Store.GetNames()
		return operations.NewGetProfileNamesOK().WithPayload(names)
	})

	// implementation for /profile/${profilename}
	api.GetProfileHandler = operations.GetProfileHandlerFunc(func(params operations.GetProfileParams) middleware.Responder {
		name := params.Profilename
		profile, err := datahub.Store.GetByName(name)
		if err != nil {
			return operations.NewGetProfileNotFound()
		}
		return operations.NewGetProfileOK().WithPayload(profile)
	})

}
