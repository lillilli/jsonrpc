# jsonrpc

[![Go Report Card](https://goreportcard.com/badge/lillilli/jsonrpc)](https://goreportcard.com/report/lillilli/jsonrpc)
[![GoDoc](https://godoc.org/github.com/lillilli/jsonrpc?status.svg)](https://godoc.org/github.com/lillilli/jsonrpc)
[![GitHub license](https://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/lillilli/jsonrpc/master/LICENSE)

## About

- Simple, Tiny, Flexible.
- No `reflect` package.
- No `Mutex` usages.
- Any method naming allowed.
- Compliance with [JSON-RPC 2.0](http://www.jsonrpc.org/specification).

## Install

```
$ go get -u github.com/lillilli/jsonrpc
```

## Usage

### Basic

```go
package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/lillilli/jsonrpc"
)

func main() {
	address := flag.String("address", ":65534", "")

	s := jsonrpc.NewServer()
	s.Handle("getHealthStatus", healthHandler)

	http.Handle("/jsonrpc/", s.Handler())
	log.Fatal(http.ListenAndServe(*address, nil))
}

func healthHandler(w jsonrpc.ResponseWriter, r *jsonrpc.Request) error {
	w.SetResponse("ok")
	return nil
}
```

### Advanced

```go
package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/lillilli/jsonrpc"
)

func main() {
	address := flag.String("address", ":65534", "")

	s := jsonrpc.NewServer()
	s.Handle("getHealthStatus", healthHandler)

	http.Handle("/", s.Handler())
	log.Fatal(http.ListenAndServe(*address, nil))
}

type healthParams struct {
	Message string `json:"message"`
}

func healthHandler(w jsonrpc.ResponseWriter, r *jsonrpc.Request) error {
	params := new(healthParams)

	// return parse error with code -32700
	if err := r.Unmarshal(params); err != nil {
		return err
	}

	// will return "error": {"code": -32602, "message": "Internal error", "data": "message must be provided"}
	if params.Message == "" {
		w.SetInvalidRequestParamsError("message must be provided")
		return nil
	}

	// will return "error": {"code": -32603, "message": "Internal error", "data": "message too short"}
	if len(params.Message) < 2 {
		w.SetErrorData("message too short")
		return nil
	}

	w.SetResponse(fmt.Sprintf("ok, %s", params.Message))
	return nil
}
```

## Examples

### Invoke the Echo method

```
POST / HTTP/1.1
Accept: application/json, */*
Accept-Encoding: gzip, deflate
Connection: keep-alive
Content-Type: application/jso

{
  "jsonrpc": "2.0",
  "method": "getHealthStatus",
  "params": {
    "name": "John Doe"
  },
  "id": "123"
}

HTTP/1.1 200 OK

Content-Length: 43
Content-Type: application/json
Date: Wed, 27 Feb 2019 10:10:57 GMT

{
  "jsonrpc": "2.0",
  "result": "ok",
  "id": "123"
}
```

### Invoke the notification

```
POST / HTTP/1.1
Accept: application/json, */*
Accept-Encoding: gzip, deflate
Connection: keep-alive
Content-Type: application/jso

{
  "jsonrpc": "2.0",
  "method": "getHealthStatus",
  "params": {
    "name": "John Doe"
  }
}
```

## License

Released under the [MIT License](https://github.com/lillilli/jsonrpc/blob/master/LICENSE).
