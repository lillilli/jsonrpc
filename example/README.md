# Example server

1. Build

```bash
go build server.go && ./server
```

2. Send requests to localhost:65534.

For example:

```bash
curl -X POST -H "Content-Type: application/json" --data '{"jsonrpc":"2.0","method":"healthStatus","params":{"message":""},"id":"1234"}' http://localhost:65534
```

or 

```bash
curl -X POST -H "Content-Type: application/json" --data '{"jsonrpc":"2.0","method":"getHealthStatus","params":{"message":"123"},"id":"1234"}' http://localhost:65534
```

or 

```bash
curl -X POST -H "Content-Type: application/json" --data '{"jsonrpc":"2.0","method":"getHealthStatus","id":"1234"}' http://localhost:65534
```

or 

```bash
curl -X POST -H "Content-Type: application/json" --data '{"jsonrpc":"2.0","method":"getHealthStatus","params":{},"id":"1234"}' http://localhost:65534
```