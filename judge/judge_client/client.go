package main

import (
	"fmt"
	"log"

	"../judgepb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("This is client part of werewolf_be")
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}

	defer conn.Close()

	client := judgepb.NewStartNewGameClient(conn)
	fmt.Printf("Created client: %f\n", client)
}
