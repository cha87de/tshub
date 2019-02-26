package commport

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"os"

	"github.com/cha87de/tshub/datahub"
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
			serverConn.datahub.Store.Keep(message.TSProfile)
		} else if message.Type == MessageTSData {
			// handle most recent monitoring data
			serverConn.datahub.Streamer.Put(message.TSData)
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
