package main

import (
	"context"
	"fmt"
	"os"
	"strconv"

	"google.golang.org/grpc"

	pb "github.com/chukmunnlee/grpc-mar072020/messages"
)

func main() {

	// Setup the connection
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	checkError(err)
	defer conn.Close()

	// Read gameId from command line
	gameId, err := strconv.Atoi(os.Args[1])
	checkError(err)

	// Create the client
	c := pb.NewBoardgameServiceClient(conn)

	// Create the request
	req := pb.FindGameByIdRequest{GameId: int32(gameId)}
	resp, err := c.FindGameById(context.Background(), &req)
	checkError(err)

	fmt.Printf("Response: %v\n", resp)
}
