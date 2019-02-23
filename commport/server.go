package commport

import (
	"fmt"
	"net"
	"os"
	"sync"

	"github.com/cha87de/tshub/datahub"
)

// NewServer creates a new Server instance
func NewServer(bindType string, bindAddr string, datahub *datahub.Hub) *Server {
	return &Server{
		bindType:   bindType,
		bindAddr:   bindAddr,
		datahub:    datahub,
		connAccess: sync.Mutex{},
	}
}

// Server opens a server socket and accepts incoming connections
type Server struct {
	// config options
	bindType string
	bindAddr string
	datahub  *datahub.Hub

	// runtime params
	listener    net.Listener
	connections []*ServerConn
	connAccess  sync.Mutex
}

// Start opens the server socket and handles incoming connections
func (server *Server) Start() {
	var err error
	server.listener, err = net.Listen(server.bindType, server.bindAddr)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	for {
		// Listen for an incoming connection.
		conn, err := server.listener.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			server.Stop()
			os.Exit(1)
		}
		// Handle connections in a new goroutine.
		serverConn := NewServerConn(server, conn, server.datahub)
		server.connAccess.Lock()
		server.connections = append(server.connections, serverConn)
		server.connAccess.Unlock()
		go serverConn.Read()
	}
}

// removeConnection removes the given serverConn from the list of active connections (called by serverConn on close)
func (server *Server) removeConnection(serverConn *ServerConn) {
	// find and remove from server.connections
	server.connAccess.Lock()
	defer server.connAccess.Unlock()
	for i, sc := range server.connections {
		if sc == serverConn {
			server.connections = append(server.connections[:i], server.connections[i+1:]...)
			break
		}
	}
}

// Stop closes all open connections and terminates the server socket
func (server *Server) Stop() {
	server.connAccess.Lock()
	defer server.connAccess.Unlock()
	for _, serverConn := range server.connections {
		serverConn.Close()
	}
	server.listener.Close()
}
