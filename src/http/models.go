package http

type HttpRequest struct {
	Body     string
	Path     string
	Protocol string
	Method   string
	Params   map[string]string
	Headers  map[string]string
}

type HttpResponse struct {
	Code     int
	Message  string
	Protocol string
	Body     string
	Headers  map[string]string
}
