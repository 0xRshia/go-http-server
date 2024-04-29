package http

import (
	"strings"
)

func ParseTcpSegment(buf []byte) *HttpRequest {
	req := &HttpRequest{
		Headers: map[string]string{},
	}
	input := string(buf)
	requestAndBody := strings.Split(input, "\r\n\r\n")
	req.Body = requestAndBody[1]
	requesAndHeader := strings.Split(requestAndBody[0], "\r\n")
	request := requesAndHeader[0]
	headers := requesAndHeader[1:]
	for _, header := range headers {
		segmentedHeader := strings.Split(header, ": ")
		req.Headers[segmentedHeader[0]] = segmentedHeader[1]
	}
	newRequest := strings.Split(request, " ")
	req.Method = newRequest[0]
	req.Path = newRequest[1]
	req.Protocol = newRequest[2]
	return req
}
