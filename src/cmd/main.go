package main

import (
	"fmt"
	"log"

	"github.com/Arshia-Izadyar/go-http-server/src"
	"github.com/Arshia-Izadyar/go-http-server/src/http"
)

func main() {
	mx := http.NewMux()
	mx.HandlerFunc("/echo/:name", http.Echo)
	mx.HandlerFunc("/echo/:name/:test", http.Echo)
	mx.HandlerFunc("/user-agent", http.UserAgent)
	mx.HandlerFunc("/", http.Home)
	server := http.Server{
		Addr:     "localhost",
		Type:     "tcp",
		Port:     4221,
		Handlers: mx,
	}
	log.Default().Printf("starting tcp server on port %d and host %s", server.Port, server.Addr)
	err := server.ListenAndServe()
	if err != nil {
		src.HandleError(err, fmt.Sprintf("Error starting tcp server on port %d and host %s", server.Port, server.Addr))
	}
}
