package server

import (
	"cagacryptotestinterview/xcontext"
	"context"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

func wsHandler(ctx context.Context) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ws, err := xcontext.WSUpgarder(ctx).Upgrade(w, r, nil)
		if err != nil {
			log.Println("Got an error", err)
			return
		}

		defer ws.Close()

		// Currently, unique property is applied to a single connections. So this is
		// a local variable.
		// If the property will be applied to all connections in the future, move
		// this variable to the global scope.
		engine := NewRandomEngine()

		for {
			_, msg, err := ws.ReadMessage()
			if err != nil {
				log.Println("Client closed the connection:", err)
				break
			}
			log.Println("Got a message from", r.RemoteAddr, ":", string(msg))

			n, err := engine.Next()
			if err != nil {
				ws.WriteMessage(websocket.TextMessage, []byte("Something wrong"))
			} else {
				ws.WriteMessage(websocket.TextMessage, []byte(n))
			}
		}
	}
}

func NewServer(ctx context.Context) *http.Server {
	ctx = xcontext.WithWSUpgrader(ctx)

	mux := http.NewServeMux()
	mux.HandleFunc("/ws", wsHandler(ctx))

	server := &http.Server{Addr: xcontext.ServerAddress(ctx), Handler: mux}
	return server
}
