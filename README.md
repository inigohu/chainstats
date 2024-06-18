# ChainStats

ChainStats is a simple Go library that allows you to chain multiple `grpc.StatsHandler` instances.

## Installation

```sh
go get github.com/inigohu/chainstats
```

### Usage

```go

package main

import (
 "context"
 "google.golang.org/grpc"
 "google.golang.org/grpc/stats"
 "github.com/inigohu/chainstats"
)

// Example custom stats handler
type customStatsHandler struct{}

func (h *customStatsHandler) TagRPC(ctx context.Context, info *stats.RPCTagInfo) context.Context {
 // Add custom logic here
 return ctx
}

func (h *customStatsHandler) HandleRPC(ctx context.Context, stat stats.RPCStats) {
 // Add custom logic here
}

func (h *customStatsHandler) TagConn(ctx context.Context, info *stats.ConnTagInfo) context.Context {
 // Add custom logic here
 return ctx
}

func (h *customStatsHandler) HandleConn(ctx context.Context, stat stats.ConnStats) {
 // Add custom logic here
}

func main() {
 customHandler1 := &customStatsHandler{}
 customHandler2 := &customStatsHandler{}

 statsHandler := chainstats.NewChainStatsHandler(customHandler1, customHandler2)

 server := grpc.NewServer(grpc.StatsHandler(statsHandler))

 // Register your gRPC services and start the server
}
```
