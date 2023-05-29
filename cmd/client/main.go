package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"time"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	apipb "github.com/troydai/control-surface/gen/proto/com/troydai/controlsurface/v1"
)

const _server = "localhost:8080"

func main() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatalf("failed to initialize zap logger: %v", err)
	}

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)

	conn, err := grpc.Dial(_server, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.Fatal("failed to dial", zap.Error(err))
	}

	client := apipb.NewControlServiceClient(conn)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	for {
		select {
		case <-sig:
			return
		default:
		}

		time.Sleep(1 * time.Second)

		if err := getParameters(ctx, client, logger); err != nil {
			logger.Error("failed to get parameters", zap.Error(err))
		}
	}
}

func getParameters(ctx context.Context, client apipb.ControlServiceClient, logger *zap.Logger) error {
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	resp, err := client.GetParameters(ctx, &apipb.GetParametersRequest{})
	if err != nil {
		return err
	}

	logger.Info("got parameters", zap.Any("parameters", resp.ParameterValues))
	return nil
}
