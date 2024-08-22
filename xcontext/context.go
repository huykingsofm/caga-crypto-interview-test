package xcontext

import (
	"context"

	"github.com/gorilla/websocket"
)

const (
	WSUpgraderKey    = "wsUpgrader"
	ServerAddressKey = "serverAddress"
)

func WithWSUpgrader(ctx context.Context) context.Context {
	return context.WithValue(ctx, WSUpgraderKey, &websocket.Upgrader{})
}

func WSUpgarder(ctx context.Context) *websocket.Upgrader {
	return ctx.Value(WSUpgraderKey).(*websocket.Upgrader)
}

func WithServerAddress(ctx context.Context, addr string) context.Context {
	return context.WithValue(ctx, ServerAddressKey, addr)
}

func ServerAddress(ctx context.Context) string {
	return ctx.Value(ServerAddressKey).(string)
}
