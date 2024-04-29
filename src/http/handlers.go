package http

import (
	"net"
	"strconv"
	"strings"
)

type Handler interface {
	ServeHTTP(req *HttpRequest, c net.Conn)
}

type HandlerFunc func(req *HttpRequest) HttpResponse

func Home(req *HttpRequest) HttpResponse {
	res := HttpResponse{
		Code:     200,
		Message:  "OK",
		Protocol: "HTTP/1.1",
		Body:     "",
		Headers: map[string]string{
			"Host":         "localhost:4221",
			"Content-Type": "text/plain",
		},
	}

	return res
}

func Echo(req *HttpRequest) HttpResponse {
	name, _ := strings.CutPrefix(req.Path, "/echo/")

	// name := req.Params["name"]
	return HttpResponse{
		Code:     200,
		Message:  "OK",
		Protocol: "HTTP/1.1",
		Body:     name,
		Headers: map[string]string{
			"Host":           "localhost:4221",
			"Content-Type":   "text/plain",
			"Content-Length": strconv.Itoa(len(name)),
		},
	}
}

func UserAgent(req *HttpRequest) HttpResponse {
	agent := req.Headers["User-Agent"]
	return HttpResponse{
		Code:     200,
		Message:  "OK",
		Protocol: "HTTP/1.1",
		Body:     agent,
		Headers: map[string]string{
			"Host":           "localhost:4221",
			"Content-Type":   "text/plain",
			"Content-Length": strconv.Itoa(len(agent)),
		},
	}
}
