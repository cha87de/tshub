package commport

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"os"

	kvmtopmodels "github.com/cha87de/kvmtop/models"
	"github.com/cha87de/tshub/datahub"
	"github.com/cha87de/tshub/util"
	"github.com/cha87de/tsprofiler/models"
)

// NewServerConn creates and returns a new instance of ServerConn
func NewServerConn(server *Server, conn net.Conn, datahub *datahub.Hub) *ServerConn {
	return &ServerConn{
		server:  server,
		conn:    conn,
		datahub: datahub,
	}
}

// ServerConn handles the communication of a net.Conn connection
type ServerConn struct {
	server  *Server
	conn    net.Conn
	datahub *datahub.Hub
}

func (serverConn *ServerConn) Read() {
	fmt.Printf("Serving %s\n", serverConn.conn.RemoteAddr().String())
	for {
		netData, err := bufio.NewReader(serverConn.conn).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			break
		}

		var message Message
		err = json.Unmarshal([]byte(netData), &message)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error while unmarshalling json: %s:\n%s\n", err, netData)
			continue
		}
		// fmt.Printf("unmarshalled: %+v\n", message)
		if message.Type == MessageTSProfile {
			// store TSProfile in datahub
			serverConn.datahub.Store.KeepProfile(message.TSProfile)
		} else if message.Type == MessageTSData {
			// handle most recent monitoring data from kvmtop
			serverConn.handleKvmtopData(message.TSData)
		} else if message.Type == MessageTSInput {
			// handle most recent monitoring data
			// TODO where is the name coming from??
			serverConn.datahub.Streamer.Put("default-name", message.TSInput)
		} else {
			// cannot handle message type
			continue
		}
	}
	serverConn.Close()
}

// Close closes the net.Conn connection
func (serverConn *ServerConn) Close() {
	fmt.Printf("Good Bye %s\n", serverConn.conn.RemoteAddr().String())
	serverConn.server.removeConnection(serverConn)
	serverConn.conn.Close()
}

// handleKvmtopData transforms the tsdata from kvmtop to generic TSInput data
func (serverConn *ServerConn) handleKvmtopData(tsdata kvmtopmodels.TSData) {
	domainlist := make([]string, len(tsdata.Domains))

	// Step 1: handle domains in kvmtop data

	for i, domain := range tsdata.Domains {
		domainname := domain["UUID"].(string)
		domainlist[i] = domainname

		// generate additional metrics
		m1, _ := util.GetFloat(domain["net_transmittedBytes"])
		m2, _ := util.GetFloat(domain["net_receivedBytes"])
		domain["net_io"] = m1 + m2

		m1, _ = util.GetFloat(domain["disk_stats_rdbytes"])
		m2, _ = util.GetFloat(domain["disk_stats_wrbytes"])
		domain["disk_io"] = m1 + m2

		// transpose data to tsinput
		metrics := make([]models.TSInputMetric, 0)
		for key, value := range domain {
			metrics = append(metrics, models.TSInputMetric{
				Name:  key,
				Value: value,
			})
		}
		tsinput := models.TSInput{
			Metrics: metrics,
		}

		// send domain to streamer
		serverConn.datahub.Streamer.Put(domainname, tsinput)
	}

	// Step 2: handle host data

	tsdata.Host["instances"] = float64(len(tsdata.Domains))
	tsdata.Host["host_domains"] = domainlist
	hostname := tsdata.Host["host_name"].(string)
	// transpose data to tsinput
	metrics := make([]models.TSInputMetric, 0)
	for key, value := range tsdata.Host {
		metrics = append(metrics, models.TSInputMetric{
			Name:  key,
			Value: value,
		})
	}
	tsinput := models.TSInput{
		Metrics: metrics,
	}
	// send domain to streamer
	serverConn.datahub.Streamer.Put(hostname, tsinput)
}
