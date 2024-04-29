package http

import (
	"fmt"
	"net"
	"os"

	"github.com/Arshia-Izadyar/go-http-server/src"
)

type Server struct {
	Addr     string
	Type     string
	Port     int
	Handlers Handler
}

func (s *Server) ListenAndServe() error {
	l, err := net.Listen(s.Type, fmt.Sprintf("%s:%d", s.Addr, s.Port))
	if err != nil {
		return err
	}

	for {
		c, err := l.Accept()
		if err != nil {
			src.HandleError(err, "error in accept connection.")
		}
		go s.handleConnectionRequests(c)
	}

}

func (s *Server) handleConnectionRequests(c net.Conn) {

	var data = make([]byte, 1024)
	n, err := c.Read(data)
	if err != nil {
		src.HandleError(err, "")
		os.Exit(1)
	}
	httpRequest := ParseTcpSegment(data[:n])

	s.Handlers.ServeHTTP(httpRequest, c)

	c.Close()

}
