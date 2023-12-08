package rpc

import (
	"ParkNavigate/global"
	"context"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	ctx    context.Context
	cancel context.CancelFunc

	//navigationClient
)

func Init() {
	ctx, cancel = context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	// navigation service port
	target := "127.0.0.1:50051"

	conn, err := grpc.DialContext(ctx, target, opts...)
	if err != nil {
		global.Logger.Panic(err)
	}

	//navigationClient := service.NewRouteGuideClient(conn)
}
