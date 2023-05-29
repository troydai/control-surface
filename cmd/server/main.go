package main

import (
	"log"
	"net"

	"go.uber.org/zap"
	"google.golang.org/grpc"

	apipb "github.com/troydai/control-surface/gen/proto/com/troydai/controlsurface/v1"
	"github.com/troydai/control-surface/internal/controlserver"
)

func main() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatalf("failed to initialize zap logger: %v", err)
	}

	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		logger.Fatal("failed to listen", zap.Error(err))
	}

	server := grpc.NewServer()
	apipb.RegisterControlServiceServer(server, controlserver.New())

	_ = server.Serve(lis)
}
