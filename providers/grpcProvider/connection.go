package grpcProvider

import (
	"encoding/json"
	"errors"
	"fmt"
	"gRPCServer/models"
	"gRPCServer/utils"
	"io"
	"log"
	"time"
)

// func (g *GRPCServer) Connect(stream Services_ConnectServer) error {

// 	byteString := []byte("csdv")
// 	err := stream.Send(&Message{Message: byteString, MessageType: "hi there!"})
// 	if err != nil {
// 		log.Println("Connect", "error sending message to client ", "", err)
// 		return err
// 	}

// 	return nil
// }

func (g *GRPCServer) Connect(stream Services_ConnectServer) error {
	var client models.ClientContext
	message, err := stream.Recv()
	if err != nil {
		log.Println("Connect", "error receiving the first request message for metadata from the client", "", err)
		return err
	}

	// Decode the data from the message.
	err = json.Unmarshal(message.Message, &client)
	if err != nil {
		log.Println("clientRegistration", "error decoding the request message from the client", "", err)
		return err
	}

	newClient := g.NewClientStream(g.RealtimeChatHubProvider.Get().(*RealtimeHub), stream, &client, g.RealtimeChatHubProvider)

	fmt.Println("RealtimeChatHubProvider.Get()", g.RealtimeChatHubProvider.Get().(*RealtimeHub).clients)
	fmt.Println("RealtimeChatHubProvider.Get() len", len(g.RealtimeChatHubProvider.Get().(*RealtimeHub).clients))

	newClient.Register()

	go newClient.WritePump()

	newClient.ReadPump()

	return nil
}

func (n *NewClient) WritePump() {
	count := 0
	for {
		fmt.Println("iteration", count)

		count++
		select {

		// wait till send channel is empty. after the write operation on send channel stream sends the message to server.
		case sendMessage, ok := <-n.Send:
			if ok {
				err := n.Stream.Send(&Message{
					MessageType: string(sendMessage.MessageType),
					Message:     sendMessage.Message,
				})
				if err != nil {
					utils.LogError("client.go", "WritePump :error sending messages to the client", n.Name, err)
				}
				utils.LogInfo("Write", "sent the message successfully", sendMessage.MessageType, nil)
			}
		case <-n.Timer.C:
			utils.LogInfo("client.go", "WritePump :ping timer finished", n.Name, nil)
			return
		}
	}
}

func (n *NewClient) ReadPump() {
	count := 0
	for {
		count++
		clientMsg, err := n.Stream.Recv()
		if err != nil {
			if err == io.EOF {
				continue
			}
			utils.LogError("client.go", "ReadPump: error getting message from server stream.", "", err)
			break
		} else {

			utils.LogInfo("Read", "read the message successfully", clientMsg.MessageType, nil)

			// comment the log below if dont want to see the messages on console
			log.Printf("Receving message : %s\n", clientMsg.MessageType)

			go n.ProcessClientMessaging(clientMsg.MessageType, clientMsg.Message)

		}
	}
}

func (n *NewClient) ProcessClientMessaging(messageType string, message []byte) {

	switch messageType {

	case models.PingMessageType:
		go n.ProcessPing()

	default:
		err := errors.New("invalid message type")
		utils.LogWarning("ProcessClientMessaging", err.Error(), messageType, err)

	}

}

func (nc *NewClient) ProcessPing() {
	log.Printf("Sending  message : %s\n", models.PongMessage)
	nc.Timer = time.Timer{C: time.After(10 * time.Second)}
	nc.Send <- models.SendMessage{
		Message:     []byte(models.PongMessage),
		MessageType: models.PongMessageType,
	}
}
