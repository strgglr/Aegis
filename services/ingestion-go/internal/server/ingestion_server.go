package server

import (
	"context"
	"ingestion-go/internal/queue"
	pb "ingestion-go/proto"
)

type IngestionServer struct {
	publisher queue.Publisher
	pb.UnimplementedEventIngestionServiceServer
	// workerPool WorkerPool
}

func (is *IngestionServer) SendEvent(ctx context.Context, event *pb.Event) (*pb.Ack, error) {
	is.publisher.Publish(event)
	return &pb.Ack{Status: "Event received successfully"}, nil
}

func NewIngestionServer(publisher queue.Publisher) *IngestionServer {
	return &IngestionServer{
		publisher: publisher,
	}
}
