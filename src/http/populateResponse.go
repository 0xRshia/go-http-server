package http

import "fmt"

func PopulateResponse(res HttpResponse) string {
	response := fmt.Sprintf("%s %d %s\r\n", res.Protocol, res.Code, res.Message)
	for k, v := range res.Headers {
		response += fmt.Sprintf("%s: %s\r\n", k, v)
	}

	// response += fmt.Sprintf("Content-Length: %d", len([]byte(res.Body)))

	response += "\r\n"
	response += res.Body

	return response
}
