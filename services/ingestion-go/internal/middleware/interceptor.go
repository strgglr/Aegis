package middleware

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
)

func IngestionInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	startTime := time.Now()
	// Log the incoming request
	log.Printf("Received request for method: %s", info.FullMethod)

	// Call the next handler in the chain
	resp, err := handler(ctx, req)

	if err != nil {
		log.Printf("Error handling request for method: %s, Error: %v", info.FullMethod, err)
	} else {
		log.Printf("Successfully handled request for method: %s in %v", info.FullMethod, time.Since(startTime))
	}

	return resp, err
}
