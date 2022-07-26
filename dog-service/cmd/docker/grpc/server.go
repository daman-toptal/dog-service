 package main

import (
	"dog-service/internal/service"
	"time"

	grpcDog "protobuf-v1/golang/service"

	ggrpc "google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

var (
	server    *ggrpc.Server
	dogServer grpcDog.DogServiceServer
)

func initGRPCServices() {
	dogServer = service.NewDogService()
}

func initGRPCServer() {
	server = ggrpc.NewServer(
		ggrpc.KeepaliveParams(keepalive.ServerParameters{
			Timeout: 100 * time.Second,
		}),
	)
	registerGRPCServerServices()
}

func registerGRPCServerServices() {
	grpcDog.RegisterDogServiceServer(server, dogServer)
}
