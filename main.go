package main

import (
	"cagacryptotestinterview/server"
	"cagacryptotestinterview/xcontext"
	"context"
	"log"
)

func main() {
	ctx := context.Background()
	ctx = xcontext.WithServerAddress(ctx, ":8000")
	log.Println("Server started at", xcontext.ServerAddress(ctx))

	svr := server.NewServer(ctx)
	if err := svr.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
