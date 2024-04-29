package http

import (
	"errors"
	"log"
	"net"
	"regexp"
	"strings"
)

type route struct {
	Regex   *regexp.Regexp
	Handler HandlerFunc
	Params  map[string]string
}

type Mux struct {
	Handlers []route
}

func NewMux() *Mux {
	return &Mux{
		Handlers: []route{},
	}
}

func (m *Mux) setHandler(path string, h HandlerFunc) error {

	r := route{Handler: h, Params: map[string]string{}}

	pattern := "^" + regexp.QuoteMeta(path) + "$"
	re := regexp.MustCompile(`:([a-zA-Z]+)`)
	found := re.FindAllString(pattern, -1)

	for _, str := range found {
		r.Params[str[1:]] = ""
		pattern = strings.Replace(pattern, str, `([a-zA-Z\-]+)`, 1)
	}

	compiledRegex, err := regexp.Compile(pattern)
	if err != nil {

		return err
	}

	r.Regex = compiledRegex
	m.Handlers = append(m.Handlers, r)
	return nil
}

func (m *Mux) HandlerFunc(path string, h HandlerFunc) error {
	if path == "" {
		return errors.New("path cannot be empty")
	}

	m.setHandler(path, h)

	return nil
}

func (m Mux) ServeHTTP(req *HttpRequest, c net.Conn) {
	for _, route := range m.Handlers {
		if matches := route.Regex.FindStringSubmatch(req.Path); matches != nil {
			if len(matches) > 1 {
				var i = 1

				for k := range route.Params {
					route.Params[k] = matches[i]
					if i < len(matches)-1 {
						i++
					} else {
						break
					}
				}
			}
			req.Params = route.Params
			response := route.Handler(req)
			httpResponse := PopulateResponse(response)
			_, err := c.Write([]byte(httpResponse))
			if err != nil {
				log.Fatal(err)
			}
			return
		}
	}
	httpResponse := PopulateResponse(HttpResponse{
		Code:     404,
		Message:  "Not Found",
		Protocol: "HTTP/1.1",
		Body:     "",
		Headers: map[string]string{
			"Content-Type": "text/plain",
		},
	})
	c.Write([]byte(httpResponse))
}
