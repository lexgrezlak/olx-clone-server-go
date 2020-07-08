package server

import (
	"fmt"
	"net"
	"strconv"
)

type Server struct {
	ip       string
	port     string
	listener net.Listener
}

// creates a new server listening on the address provided
// if no port is given it choses one randomly
func New(port string) (*Server, error) {
	addr := fmt.Sprintf(":" + port)
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return nil, fmt.Errorf("failed to create listener on %s: %w", addr, err)
	}

	return &Server{
		ip:   listener.Addr().(*net.TCPAddr).IP.String(),
		port: strconv.Itoa(listener.Addr().(*net.TCPAddr).Port),
		listener: listener,
	}, nil
}
