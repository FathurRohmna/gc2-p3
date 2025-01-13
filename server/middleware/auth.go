package middleware

import (
	"context"
	"os"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	"server/utils"
)

func AuthMiddleware() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, status.Errorf(codes.Unauthenticated, "missing metadata")
		}

		token := md["authorization"]
		if len(token) < 1 {
			return nil, status.Errorf(codes.Unauthenticated, "missing authorization token")
		}

		if !strings.HasPrefix(token[0], "Bearer ") {
			return nil, status.Errorf(codes.Unauthenticated, "invalid token format")
		}

		tokenString := strings.TrimPrefix(token[0], "Bearer ")

		secretKey := os.Getenv("SECRET")
		userID, err := utils.ValidateToken(tokenString, secretKey)
		if err != nil {
			return nil, status.Errorf(codes.Unauthenticated, "invalid token: %v", err)
		}

		ctx = context.WithValue(ctx, "user_id", userID)

		return handler(ctx, req)
	}
}
