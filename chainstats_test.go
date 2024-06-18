package chainstats

import (
	"context"
	"testing"

	"google.golang.org/grpc/stats"
)

type mockStatsHandler struct {
	tagRPCCalled     bool
	handleRPCCalled  bool
	tagConnCalled    bool
	handleConnCalled bool
}

func (m *mockStatsHandler) TagRPC(ctx context.Context, info *stats.RPCTagInfo) context.Context {
	m.tagRPCCalled = true
	return ctx
}

func (m *mockStatsHandler) HandleRPC(ctx context.Context, stat stats.RPCStats) {
	m.handleRPCCalled = true
}

func (m *mockStatsHandler) TagConn(ctx context.Context, info *stats.ConnTagInfo) context.Context {
	m.tagConnCalled = true
	return ctx
}

func (m *mockStatsHandler) HandleConn(ctx context.Context, stat stats.ConnStats) {
	m.handleConnCalled = true
}

func TestChainStatsHandler(t *testing.T) {
	handler1 := &mockStatsHandler{}
	handler2 := &mockStatsHandler{}
	chainHandler := NewChainStatsHandler(handler1, handler2)

	ctx := context.Background()
	info := &stats.RPCTagInfo{}
	stat := &stats.InPayload{}

	// Test TagRPC
	ctx = chainHandler.TagRPC(ctx, info)
	if !handler1.tagRPCCalled || !handler2.tagRPCCalled {
		t.Errorf("TagRPC was not called on all handlers")
	}

	// Test HandleRPC
	chainHandler.HandleRPC(ctx, stat)
	if !handler1.handleRPCCalled || !handler2.handleRPCCalled {
		t.Errorf("HandleRPC was not called on all handlers")
	}

	// Test TagConn
	connInfo := &stats.ConnTagInfo{}
	ctx = chainHandler.TagConn(ctx, connInfo)
	if !handler1.tagConnCalled || !handler2.tagConnCalled {
		t.Errorf("TagConn was not called on all handlers")
	}

	// Test HandleConn
	connStat := &stats.ConnBegin{}
	chainHandler.HandleConn(ctx, connStat)
	if !handler1.handleConnCalled || !handler2.handleConnCalled {
		t.Errorf("HandleConn was not called on all handlers")
	}
}
