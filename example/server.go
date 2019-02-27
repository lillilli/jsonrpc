package main

import (
	"errors"
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
	s.Handle("healthStatus", baseHealthHandler)

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
		return errors.New("lol")
	}

	// will return "error": {"code": -32603, "message": "Internal error", "data": "message too short"}
	if len(params.Message) < 2 {
		w.SetErrorData("message too short")
		return nil
	}

	w.SetResponse(fmt.Sprintf("ok, %s", params.Message))
	return nil
}

func baseHealthHandler(w jsonrpc.ResponseWriter, r *jsonrpc.Request) error {
	w.SetResponse("ok")
	return nil
}
