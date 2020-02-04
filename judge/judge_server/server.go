package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"strconv"

	"../judgepb"

	"google.golang.org/grpc"
)

type server struct{}

// 房主创建房间以后，server 会根据所发送的 req 返回 开房信息，其中包括房间的密码
func (*server) StartGameService(ctx context.Context, req *judgepb.StartGameRequest) (*judgepb.StartGameResponse, error) {
	playerAmount := req.GetNewGameInfo().GetPlayerAmount()
	gamePattern := req.GetNewGameInfo().GetGamePattern()
	result := "Player amount is " + strconv.Itoa(int(playerAmount)) + ". The game pattern is "

	res := &judgepb.StartGameResponse{
		ResponseStartgame: result,
		PlayerAmount:      playerAmount,
		GamePattern:       gamePattern,
		NewGamePassword: &judgepb.Password{
			Num1: 1,
			Num2: 2,
			Num3: 3,
			Num4: 4,
			Num5: 5,
			Num6: 6,
		},
	}
	return res, nil
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
