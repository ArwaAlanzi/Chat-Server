package main

import (
	"log"
	"net"
	"os"

	"grpcChatServer/chatserver"

	"google.golang.org/grpc"
)

func main() {
	Port := os.Getenv("PORT")
	if Port == "" {
		Port = "5000"
	}

	listen, err := net.Listen("tcp", ":"+Port)
	if err != nil {
		log.Fatalf("Is not listening %v : %v", Port, err)
	}
	log.Println("Is now listening : " + Port)
	grpcserver := grpc.NewServer()

	//register ChatService
	cs := chatserver.ChatServer{}
	chatserver.RegisterServicesServer(grpcserver, &cs)

	//grpc listen and serve
	err = grpcserver.Serve(listen)
	if err != nil {
		log.Fatalf("Sorry failed to start gRPC Server :: %v", err)
	}
}
