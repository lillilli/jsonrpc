package jsonrpc

import (
	"encoding/json"
	"net/http"
)

// ResponseWriter represents response writer interface.
type ResponseWriter interface {
	SetResponse(v interface{})
	SetErrorData(v interface{})
	SetInvalidRequestParamsError(data interface{})
}

func newErrorResponse(err *respError) *response {
	return &response{
		Version: Version,
		Error:   err,
	}
}

// newResponse generates a JSON-RPC response.
func newResponse(r *Request) *response {
	return &response{
		ID:      r.ID,
		Version: r.Version,
	}
}

// A response represents a JSON-RPC response returned by the server.
type response struct {
	Version string      `json:"jsonrpc"`
	Result  interface{} `json:"result,omitempty"`
	Error   *respError  `json:"error,omitempty"`
	ID      interface{} `json:"id"`
}

// SetResponse - set response result
func (r *response) SetResponse(v interface{}) {
	r.Result = v
}

// SetInvalidRequestParamsError - set response error to invalid req params
func (r *response) SetInvalidRequestParamsError(data interface{}) {
	r.Error = errInvalidParams()
	r.Error.Data = data
}

// SetErrorData - set response error data
func (r *response) SetErrorData(v interface{}) {
	if r.Error == nil {
		r.Error = errInternal()
	}

	r.Error.Data = v
}

// sendResponse writes JSON-RPC response.
func sendResponse(w http.ResponseWriter, resp []*response) error {
	w.Header().Set(contentTypeKey, contentTypeValue)

	if len(resp) > 1 {
		return json.NewEncoder(w).Encode(resp)
	}

	if len(resp) == 1 {
		return json.NewEncoder(w).Encode(resp[0])
	}

	return nil
}
