package main

import (
	"context"
	dh "indexer/data_handler"
	dr "indexer/data_retriver"
	grpc_service "indexer/grpc"
	_ "indexer/util"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup

	dh.Start(ctx, &wg)
	dr.Start(ctx, &wg)
	log.Info().Msg("start grpc server1")
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to listen")
	}

	grpcServer := grpc.NewServer()
	grpc_service.RegisterServerServer(grpcServer, &grpc_service.ServerImpl{})

	go func() {
		log.Info().Msg("Starting gRPC server on :50051")
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatal().Err(err).Msg("Failed to serve gRPC")
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	sig := <-sigChan
	log.Info().Msgf("Got signal %s, shutting down...", sig)

	grpcServer.GracefulStop()
	cancel()
	wg.Wait()
	log.Info().Msgf("Shutdown complete")
}
