package jsonrpc

import (
	"net/http"
)

// Handler - represents jsonrpc handler
type Handler func(w ResponseWriter, r *Request) error

type httpHandler struct {
	routes map[string]Handler
}

// ServeHTTP provides basic JSON-RPC handling.
func (h *httpHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := h.serveHTTPReq(w, r); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (h *httpHandler) serveHTTPReq(w http.ResponseWriter, r *http.Request) error {
	requests, err := parseRequest(r)
	if err != nil {
		return sendResponse(w, []*response{newErrorResponse(err)})
	}

	responses := make([]*response, 0)

	for _, request := range requests {
		if resp := h.HandleMethod(request); request.ID != nil {
			responses = append(responses, resp)
		}
	}

	return sendResponse(w, responses)
}

// InvokeMethod invokes JSON-RPC method.
func (h *httpHandler) HandleMethod(r *Request) *response {
	res := newResponse(r)

	handler, err := h.GetMethod(r)
	if err != nil {
		res.Error = err
		return res
	}

	if handlerErr := handler(res, r); handlerErr != nil {
		if r.parsingError {
			res.Error = errParse()
		} else {
			res.Error = errInternal()
		}
	}

	return res
}

// TakeMethod takes jsonrpc.Func in MethodRepository.
func (h *httpHandler) GetMethod(r *Request) (Handler, *respError) {
	if r.Method == "" || r.Version != Version {
		return nil, errInvalidRequest()
	}

	handler, ok := h.routes[r.Method]
	if !ok {
		return nil, errMethodNotFound()
	}

	return handler, nil
}
