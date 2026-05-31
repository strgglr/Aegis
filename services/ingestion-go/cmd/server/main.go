package main

import (
	"ingestion-go/internal/middleware"
	"ingestion-go/internal/queue"
	"ingestion-go/internal/server"
	pb "ingestion-go/proto"
	"net"
	"os"
	"os/signal"

	"google.golang.org/grpc"
)

func main() {

	// create a publisher
	publisher := queue.NewRedisPublisher()

	// create a ingestion server which will use the publisher to publish the events and accept the events from the gRPC client
	ingestionServer := server.NewIngestionServer(publisher)

	// next we need is grpc server to listen to the incoming requests and pass it to the ingestion server
	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(middleware.IngestionInterceptor))

	// register the ingestion server with the grpc server
	pb.RegisterEventIngestionServiceServer(grpcServer, ingestionServer)

	// next we need to start the gRPC server and listen to the incoming requests
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}
	go func() {
		err := grpcServer.Serve(listener)
		if err != nil {
			panic(err)
		}
	}()

	shutdownCh := make(chan os.Signal, 1)
	signal.Notify(shutdownCh, os.Interrupt)
	<-shutdownCh
	println("Shutting down server gracefully in 3 seconds...")
	grpcServer.GracefulStop()
}
