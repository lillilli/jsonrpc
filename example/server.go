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

	http.Handle("/", s.Handler())
	log.Fatal(http.ListenAndServe(*address, nil))
}

func healthHandler(w jsonrpc.ResponseWriter, r *jsonrpc.Request) error {
	w.SetResponse("ok")
	return nil
}
