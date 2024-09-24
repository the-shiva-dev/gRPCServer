package service

import (
	"gRPCServer/providers"
	"gRPCServer/providers/grpcProvider"
	"gRPCServer/utils"
	"log"
)

var GRPCPortAddress = "50000"

type Service struct {
	// RealTimeClient *grpcProvider.RealtimeHub
	RealtimeChatProvider providers.RealtimeChatHubProvider
	GRPC                 *grpcProvider.GRPCServer
}

func ServiecsInit() *Service {

	utils.LogDebug("Server Init", "Starting the server...", "", nil)

	realtimeChatProvider := grpcProvider.NewRealtimeChatProvider()

	newgRPCServer := grpcProvider.GRPCServerProvider("50000", realtimeChatProvider)

	return &Service{
		RealtimeChatProvider: realtimeChatProvider,
		GRPC:                 newgRPCServer,
	}
}

// Start the gRPC server for Agent communication.
func (srv *Service) StartgRPCServer() {
	utils.LogInfo("StartgRPCServer", "Starting the gRPC server", "", nil)
	log.Printf("starting  the listener: %v", GRPCPortAddress)
	if err := srv.GRPC.GRPCServer.Serve(srv.GRPC.Listener); err != nil {
		log.Fatalf("Unable to serve the listener: %v", GRPCPortAddress)
		utils.LogError("StartgRPCServer", "Unable to serve the listener: %v", GRPCPortAddress, err)
		return
	}
	utils.LogInfo("StartgRPCServer", "gRPC server started", "", nil)
}
