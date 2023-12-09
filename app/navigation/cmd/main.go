package main

import (
	"ParkNavigate/app/navigation/internal/service"
	"ParkNavigate/global"
	"ParkNavigate/pkg/pb/navigation_pb"
	"ParkNavigate/setting"
	"net"

	"google.golang.org/grpc"
)

func main() {
	setting.InitConfig()
	setting.InitLogger()

	lis, err := net.Listen("tcp", "127.0.0.1:8081")
	if err != nil {
		global.Logger.Fatal("failed to listen: %v", err)
	}
	global.Logger.Println("listeing to %v", lis.Addr())
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)

	navigation_pb.RegisterNavigationServiceServer(grpcServer, newServer())
	grpcServer.Serve(lis)

}

func newServer() *service.NavigationService {
	return service.NewNavigationService()
}
