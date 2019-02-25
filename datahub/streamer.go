package datahub

import (
	"sync"

	kvmtopmodels "github.com/cha87de/kvmtop/models"
	"github.com/cha87de/tshub/tsdb"
	"github.com/cha87de/tshub/util"
)

// NewStreamer returns a new instance of Streamer
func NewStreamer(hub *Hub) *Streamer {
	return &Streamer{
		hub: hub,

		// collections
		access:     sync.Mutex{},
		HostDomain: make(map[string][]string),

		Hosts:   make(map[string][]*tsdb.Store),
		Domains: make(map[string][]*tsdb.Store),
	}
}

// Streamer handles the distribution of TSData monitoring items
type Streamer struct {
	hub *Hub

	// collections
	access     sync.Mutex
	HostDomain map[string][]string
	Hosts      map[string][]*tsdb.Store
	Domains    map[string][]*tsdb.Store
}

// Put transmits a new TSData monitoring item to the Streamer
func (streamer *Streamer) Put(tsdata kvmtopmodels.TSData) {
	hostname := tsdata.Host["host_name"].(string)
	for _, domain := range tsdata.Domains {
		domainname := domain["UUID"].(string)

		// add domain to host map
		streamer.addHostDomainMapping(hostname, domainname)

		// generate additional metrics
		m1, _ := util.GetFloat(domain["net_transmittedBytes"])
		m2, _ := util.GetFloat(domain["net_receivedBytes"])
		domain["net_io"] = m1 + m2
		m1, _ = util.GetFloat(domain["disk_stats_rdbytes"])
		m2, _ = util.GetFloat(domain["disk_stats_wrbytes"])
		domain["disk_io"] = m1 + m2

		// update Domain measurements
		if _, exists := streamer.Domains[domainname]; !exists {
			streamer.Domains[domainname] = make([]*tsdb.Store, 5)
			streamer.Domains[domainname][0] = tsdb.NewStore(30, resolution1min)
			streamer.Domains[domainname][1] = tsdb.NewStore(30, resolution30min)
			streamer.Domains[domainname][2] = tsdb.NewStore(30, resolution1h)
			streamer.Domains[domainname][3] = tsdb.NewStore(30, resolution12h)
			streamer.Domains[domainname][4] = tsdb.NewStore(30, resolution24h)
		}
		streamer.Domains[domainname][0].Add(domain)
		streamer.Domains[domainname][1].Add(domain)
		streamer.Domains[domainname][2].Add(domain)
		streamer.Domains[domainname][3].Add(domain)
		streamer.Domains[domainname][4].Add(domain)
	}

	// update host measurements
	tsdata.Host["instances"] = len(tsdata.Domains)
	if _, exists := streamer.Hosts[hostname]; !exists {
		streamer.Hosts[hostname] = make([]*tsdb.Store, 1)
		streamer.Hosts[hostname][0] = tsdb.NewStore(30, resolution1min)
		streamer.Hosts[hostname][1] = tsdb.NewStore(30, resolution30min)
		streamer.Hosts[hostname][2] = tsdb.NewStore(30, resolution1h)
		streamer.Hosts[hostname][3] = tsdb.NewStore(30, resolution12h)
		streamer.Hosts[hostname][4] = tsdb.NewStore(30, resolution24h)
	}
	streamer.Hosts[hostname][0].Add(tsdata.Host)
	streamer.Hosts[hostname][1].Add(tsdata.Host)
	streamer.Hosts[hostname][2].Add(tsdata.Host)
	streamer.Hosts[hostname][3].Add(tsdata.Host)
	streamer.Hosts[hostname][4].Add(tsdata.Host)
}

func (streamer *Streamer) addHostDomainMapping(hostname string, domainname string) {
	found := false
	if domains, ok := streamer.HostDomain[hostname]; ok {
		for _, s := range domains {
			if s == domainname {
				found = true
				break
			}
		}
	} else {
		streamer.HostDomain[hostname] = make([]string, 0)
	}
	if !found {
		streamer.HostDomain[hostname] = append(streamer.HostDomain[hostname], domainname)
	}
}
