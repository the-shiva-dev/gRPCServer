package main

import (
	"gRPCServer/service"
)

func main() {

	service := service.ServiecsInit()

	go service.RealtimeChatProvider.Run()

	service.StartgRPCServer()

}
