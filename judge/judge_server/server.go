package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"../judgepb"

	"google.golang.org/grpc"
)

type server struct{}

func (*server) StartGameService(ctx context.Context, req *judgepb.StartGameRequest) (*judgepb.StartGameResponse, error) {
	return nil, nil
}

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
