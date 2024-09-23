package grpcProvider

import (
	"gRPCServer/models"
	"gRPCServer/providers"
	"time"
)

type NewClient struct {
	Name         string
	Hostname     string
	IPAddress    string
	Platform     string
	OSAndVersion string

	Stream      Services_ConnectServer
	HUB         *RealtimeHub
	Send        chan models.SendMessage
	Timer       time.Timer
	Ticker      *time.Ticker
	RealtimeHub providers.RealtimeChatHubProvider
}

// NewClientStream gets the metaData and stream of agent and keep it in the agent HUB map.
func (s *GRPCServer) NewClientStream(hub *RealtimeHub, stream Services_ConnectServer, clientContext *models.ClientContext, realtimeHub providers.RealtimeChatHubProvider) *NewClient {

	ticker := time.NewTicker(10 * time.Second)
	return &NewClient{
		Name:        clientContext.Name,
		Hostname:    clientContext.Hostname,
		Platform:    clientContext.Platform,
		Stream:      stream,
		HUB:         hub,
		Send:        make(chan models.SendMessage, 1),
		Timer:       time.Timer{},
		Ticker:      ticker,
		RealtimeHub: realtimeHub,
	}
}

func (newClient *NewClient) Get() *NewClient {
	return newClient
}

func (newClient *NewClient) Register() {
	newClient.HUB.register <- newClient
}

func (newClient *NewClient) Unregister() {
	newClient.HUB.deregister <- newClient
}
