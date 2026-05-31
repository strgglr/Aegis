package main

import (
	"context"
	pb "ingestion-go/proto"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// create a gRPC client
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewEventIngestionServiceClient(conn)

	// call the gRPC method
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	resp, err := client.SendEvent(ctx, &pb.Event{EventId: "123", EventType: "test_event", Source: "test_source", Severity: "Normal", Timestamp: time.Now().Format("2006-01-02 15:04:05"), Payload: "This is a test event"})
	if err != nil {
		log.Fatalf("Error calling SendEvent: %v", err)
	}
	log.Printf("Response: %v", resp)
}
