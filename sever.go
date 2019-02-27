package jsonrpc

import (
	"net/http"
)

// Server - represents jsonrpc server interface
type Server interface {
	Handle(method string, handler Handler)
	Handler() http.Handler
}

// A MethodRepository has JSON-RPC method functions.
type server struct {
	routes map[string]Handler
}

// NewServer - returns new jsonrpc server instance.
func NewServer() Server {
	return &server{
		routes: make(map[string]Handler),
	}
}

func (s *server) Handle(method string, handler Handler) {
	s.routes[method] = handler
}

func (s *server) Handler() http.Handler {
	routes := make(map[string]Handler)

	for method, handler := range s.routes {
		routes[method] = handler
	}

	return &httpHandler{
		routes: routes,
	}
}
