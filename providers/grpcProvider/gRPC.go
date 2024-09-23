package grpcProvider

import (
	"log"
	"net"
	"gRPCServer/providers"

	"google.golang.org/grpc"
)

type GRPCServer struct {
	ServicesServer
	ServicesClient
	RealtimeChatHubProvider providers.RealtimeChatHubProvider
	Listener                net.Listener
	GRPCServer              *grpc.Server
}

func GRPCServerProvider(address string, realtimeChatHubProvider providers.RealtimeChatHubProvider) *GRPCServer {
	lis, err := net.Listen("tcp", ":"+address)
	if err != nil {
		log.Fatal("unable to start listener at tcp address : " + address)
	}

	newGRPCServer := grpc.NewServer(
	// grpc.KeepaliveEnforcementPolicy(kaep),
	// grpc.KeepaliveParams(kasp),
	)

	RegisterServicesServer(newGRPCServer, &GRPCServer{
		RealtimeChatHubProvider: realtimeChatHubProvider,
		Listener:                lis})

	return &GRPCServer{
		Listener:   lis,
		GRPCServer: newGRPCServer,
	}
}
