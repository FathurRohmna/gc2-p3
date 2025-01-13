package grpc

import (
	"client/pb"
	"context"
	"fmt"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func NewGRPCClient() (pb.UserServiceClient, error) {
	conn, err := grpc.Dial(os.Getenv("GRPC_SERVER_ADDRESS"), grpc.WithInsecure())
	if err != nil {
		return nil, fmt.Errorf("could not connect to gRPC server: %w", err)
	}

	client := pb.NewUserServiceClient(conn)
	return client, nil
}

func NewGRPCBookClient() (pb.BookServiceClient, error) {
	conn, err := grpc.Dial("localhost:50052", grpc.WithInsecure())
	if err != nil {
		return nil, fmt.Errorf("failed to connect to gRPC server: %w", err)
	}

	return pb.NewBookServiceClient(conn), nil
}

func AddJWTToContext(ctx context.Context, token string) context.Context {
	return metadata.AppendToOutgoingContext(ctx, "authorization", fmt.Sprintf("Bearer %s", token))
}
