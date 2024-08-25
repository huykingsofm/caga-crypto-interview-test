package main

import (
	"cagacryptotestinterview/server"
	"cagacryptotestinterview/xcontext"
	"context"
	"log"
)

func main() {
	ctx := context.Background()

	// Uncomment the following line to set a small max value for the random
	// engine.
	// ctx = xcontext.WithUniqueRandomEngineMaxValue(ctx, 10)

	log.Println("Server started at", xcontext.ServerAddress(ctx))

	svr := server.NewServer(ctx)
	if err := svr.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
