package chainstats

import (
	"context"

	"google.golang.org/grpc/stats"
)

// ChainStatsHandler is a composite stats.Handler that chains multiple stats.Handler instances.
type ChainStatsHandler struct {
	handlers []stats.Handler
}

// NewChainStatsHandler creates a new ChainStatsHandler.
func NewChainStatsHandler(handlers ...stats.Handler) *ChainStatsHandler {
	return &ChainStatsHandler{handlers: handlers}
}

// TagRPC calls TagRPC on all the chained stats.Handler instances.
func (c *ChainStatsHandler) TagRPC(ctx context.Context, info *stats.RPCTagInfo) context.Context {
	for _, handler := range c.handlers {
		ctx = handler.TagRPC(ctx, info)
	}
	return ctx
}

// HandleRPC calls HandleRPC on all the chained stats.Handler instances.
func (c *ChainStatsHandler) HandleRPC(ctx context.Context, stat stats.RPCStats) {
	for _, handler := range c.handlers {
		handler.HandleRPC(ctx, stat)
	}
}

// TagConn calls TagConn on all the chained stats.Handler instances.
func (c *ChainStatsHandler) TagConn(ctx context.Context, info *stats.ConnTagInfo) context.Context {
	for _, handler := range c.handlers {
		ctx = handler.TagConn(ctx, info)
	}
	return ctx
}

// HandleConn calls HandleConn on all the chained stats.Handler instances.
func (c *ChainStatsHandler) HandleConn(ctx context.Context, stat stats.ConnStats) {
	for _, handler := range c.handlers {
		handler.HandleConn(ctx, stat)
	}
}
