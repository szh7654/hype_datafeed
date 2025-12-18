package main

import (
	"context"
	dh "indexer/data_handler"
	dr "indexer/data_retriver"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/rs/zerolog/log"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup

	dh.Start(ctx, &wg)
	dr.Start(ctx, &wg)

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	sig := <-sigChan
	log.Info().Msgf("Got signal %s, shutting down...", sig)
	cancel()
	wg.Wait()
	log.Info().Msgf("Shutdown complete")
}
