package httplog

import "net/http"

type Server struct {
}

func NewHTTPServer(addr string) *http.Server {
	http.Server{}
}
