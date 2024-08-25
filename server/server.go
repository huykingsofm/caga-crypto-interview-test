package server

import (
	"cagacryptotestinterview/xcontext"
	"context"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var uniqueEngine = NewUniqueRandomEngineWithMax(10)

func wsHandler(ctx context.Context) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ws, err := xcontext.WSUpgarder(ctx).Upgrade(w, r, nil)
		if err != nil {
			log.Println("Got an error", err)
			return
		}

		defer ws.Close()

		for {
			_, msg, err := ws.ReadMessage()
			if err != nil {
				log.Println("Client closed the connection:", err)
				break
			}

			// Ignore if the message is empty.
			if len(msg) == 0 {
				ws.WriteMessage(websocket.TextMessage, []byte("got an empty message"))
				continue
			}

			log.Println("Got a message from", r.RemoteAddr, ":", string(msg))

			n, err := uniqueEngine.Next()
			if err != nil {
				ws.WriteMessage(
					websocket.TextMessage,
					[]byte("Something wrong, contact to administrator to know more information"),
				)
				log.Println("Got an error from random engine:", err)
			} else {
				// Send the unique random value to the client.
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
