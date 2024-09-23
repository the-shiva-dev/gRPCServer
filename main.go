package main

import (
	"gRPCServer/service"
)

func main() {

	service := service.ServiecsInit()

	go service.RealTimeClient.Run()

	service.StartgRPCServer()

}
