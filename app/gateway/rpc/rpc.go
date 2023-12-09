package rpc

import (
	"ParkNavigate/global"
	"ParkNavigate/pkg/pb/navigation_pb"
	"context"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	ctx    context.Context
	cancel context.CancelFunc

	NavigationClient navigation_pb.NavigationServiceClient
)

func Init() {
	ctx, cancel = context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	// navigation service port
	target := "127.0.0.1:8081"

	conn, err := grpc.DialContext(ctx, target, opts...)
	if err != nil {
		global.Logger.Panic(err)
	}

	NavigationClient = navigation_pb.NewNavigationServiceClient(conn)
}
