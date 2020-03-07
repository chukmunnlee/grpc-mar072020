package main

import (
	"context"
	"fmt"
	"net/http"

	"google.golang.org/grpc"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"

	gw "github.com/chukmunnlee/grpc-mar072020/messages"
)

func main() {
	mux := runtime.NewServeMux()

	opts := []grpc.DialOption{grpc.WithInsecure()}

	// Register the gateway
	fmt.Println("Register the gateway")
	checkError(
		gw.RegisterBoardgameServiceHandlerFromEndpoint(
			context.Background(),
			mux,
			"localhost:50051",
			opts,
		),
	)

	// Start the HTTP server
	fmt.Println("Start HTTP ")
	http.ListenAndServe(":8080", mux)

}
