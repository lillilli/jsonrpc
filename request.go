package jsonrpc

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

// Request represents a JSON-RPC request received by the server.
type Request struct {
	Version string          `json:"jsonrpc"`
	Method  string          `json:"method"`
	Params  json.RawMessage `json:"params"`
	ID      interface{}     `json:"id"`

	parsingError bool
}

// Unmarshal unmarshal req params to specified structure.
// Unmarshal will set parse error if unmarshaling failed.
func (r *Request) Unmarshal(v interface{}) error {
	if err := json.Unmarshal(r.Params, v); err != nil {
		r.parsingError = true
		return err
	}

	return nil
}

// parseRequest parses a HTTP request to JSON-RPC request.
func parseRequest(r *http.Request) (requests []*Request, reqParseError *respError) {
	// check for content type
	if !strings.HasPrefix(r.Header.Get(contentTypeKey), contentTypeValue) {
		return requests, errInvalidRequest()
	}

	defer func(r *http.Request) {
		if err := r.Body.Close(); err != nil {
			reqParseError = errInternal()
		}
	}(r)

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return requests, errInvalidRequest()
	}

	// check for batch
	if bytes.ContainsRune(b[:1], batchRequestKey) {
		requests, reqParseError = unmarshalBatchReq(b)
	} else {
		requests, reqParseError = unmarshalReq(b)
	}

	return requests, reqParseError
}

func unmarshalBatchReq(b []byte) ([]*Request, *respError) {
	var rs []*Request

	if err := json.Unmarshal(b, &rs); err != nil {
		return nil, errParse()
	}

	return rs, nil
}

func unmarshalReq(b []byte) ([]*Request, *respError) {
	rs := make([]*Request, 1)
	req := new(Request)

	if err := json.Unmarshal(b, &req); err != nil {
		return nil, errParse()
	}

	rs[0] = req
	return rs, nil
}
