package main

import (
	"bufio"
	"context"
	"fmt"
	"grpcChatServer/chatserver"
	"log"
	"os"
	"strings"

	"google.golang.org/grpc"
)

func main() {

	fmt.Println("Please enter the Server IP an d port as (IP:Port):")
	reader := bufio.NewReader(os.Stdin)
	serverID, err := reader.ReadString('\n')

	if err != nil {
		log.Printf("Sorry failed to read from console, try again:: %v", err)
	}
	serverID = strings.Trim(serverID, "\r\n")

	log.Println("The server is now connecting : " + serverID)
	//connect to grpc server
	conn, err := grpc.Dial(serverID, grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Sorry faile to conncet to gRPC server : %v", err)
	}
	defer conn.Close()

	//call ChatService to create a stream
	client := chatserver.NewServicesClient(conn)

	stream, err := client.ChatService(context.Background())
	if err != nil {
		log.Fatalf("Sorry failed to call ChatService : %v", err)
	}

	// implement communication with gRPC server
	ch := clienthandle{stream: stream}
	ch.clientConfig()
	go ch.sendMessage()
	go ch.receiveMessage()

	//blocker
	bl := make(chan bool)
	<-bl
}

// clienthandle
type clienthandle struct {
	stream     chatserver.Services_ChatServiceClient
	clientName string
}

func (ch *clienthandle) clientConfig() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Username : ")
	name, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("Sorry failed to read from console :: %v", err)
	}
	ch.clientName = strings.Trim(name, "\r\n")

}

// send message
func (ch *clienthandle) sendMessage() {

	// create a loop
	for {

		reader := bufio.NewReader(os.Stdin)
		clientMessage, err := reader.ReadString('\n')
		if err != nil {
			log.Fatalf("Sorry failed to read from console :: %v", err)
		}
		clientMessage = strings.Trim(clientMessage, "\r\n")

		clientMessageBox := &chatserver.FromClient{
			Name: ch.clientName,
			Body: clientMessage,
		}

		err = ch.stream.Send(clientMessageBox)

		if err != nil {
			log.Printf("Error while sending message to server :: %v", err)
		}

	}
}

// receive message
func (ch *clienthandle) receiveMessage() {

	//create a loop
	for {
		mssg, err := ch.stream.Recv()
		if err != nil {
			log.Printf("Error in receiving message from server :: %v", err)
		}

		//print message to console
		fmt.Printf("%s : %s \n", mssg.Name, mssg.Body)

	}
}
