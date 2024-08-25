package xcontext

import (
	"context"

	"github.com/gorilla/websocket"
)

type xcontextKey int

const (
	wsUpgraderKey xcontextKey = iota
	serverAddressKey
	uniqueRandomEngineMaxValueKey
)

func WithWSUpgrader(ctx context.Context) context.Context {
	return context.WithValue(ctx, wsUpgraderKey, &websocket.Upgrader{})
}

func WSUpgarder(ctx context.Context) *websocket.Upgrader {
	upgrader := ctx.Value(wsUpgraderKey)
	if upgrader == nil {
		return &websocket.Upgrader{}
	}

	return upgrader.(*websocket.Upgrader)
}

func WithServerAddress(ctx context.Context, addr string) context.Context {
	return context.WithValue(ctx, serverAddressKey, addr)
}

func ServerAddress(ctx context.Context) string {
	svaddr := ctx.Value(serverAddressKey)
	if svaddr == nil {
		return ":8000"
	}

	return svaddr.(string)
}
