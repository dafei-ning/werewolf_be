package main

import (
	"fmt"
	"log"
	"net"

	"../judgepb"

	"google.golang.org/grpc"
)

type server struct{}

func main() {
	fmt.Println("This is the server part of werewolf_be")
	lis, err := net.Listen("tcp", "0.0.0.0.50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	judgepb.RegisterStartNewGameServer(s, &server{})

	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

}
