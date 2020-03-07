package main

import (
	"context"
	"fmt"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"

	"database/sql"
	_ "github.com/go-sql-driver/mysql"

	pb "github.com/chukmunnlee/grpc-mar072020/messages"
)

type BGGService struct {
	pb.UnimplementedBoardgameServiceServer
	db *sql.DB
}

func (s *BGGService) Open(user string, passwd string) error {
	dsn := fmt.Sprintf("%s:%s@/bgg", user, passwd)
	db, err := sql.Open("mysql", dsn)
	s.db = db
	return err
}

func (s *BGGService) Close() error {
	return s.db.Close()
}

func (s *BGGService) FindGameById(ctx context.Context, req *pb.FindGameByIdRequest) (*pb.FindGameByIdResponse, error) {
	gameId := req.GetGameId()

	fmt.Printf("Gameid: %d\n", gameId)

	rows, err := s.db.Query("select gid, name, ranking, url from game where gid = ?", gameId)
	if nil != err {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("SQL error: %v", err),
		)
	}
	defer rows.Close()

	if !rows.Next() {
		return &pb.FindGameByIdResponse{
				Status: pb.FindGameByIdResponse_NOT_FOUND,
				Text:   fmt.Sprintf("Gameid %d does not exists", gameId),
			},
			nil
	}

	g := pb.Boardgame{}
	err = rows.Scan(&g.GameId, &g.Name, &g.Ranking, &g.Url)
	if nil != err {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Cannot read record: %v", err),
		)
	}

	return &pb.FindGameByIdResponse{
		Boardgame: &g,
		Status:    pb.FindGameByIdResponse_FOUND,
		Text:      "Found your game",
	}, nil
}

func main() {
	// Listen to a port
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	checkError(err)

	// Create a gGRPC server
	grpcServ := grpc.NewServer()

	reflection.Register(grpcServ)

	// Pass your service
	bggSvc := BGGService{}

	fmt.Println("Open connection to database...")
	checkError(bggSvc.Open("fred", "fred"))
	defer bggSvc.Close()

	// Register with your Server
	pb.RegisterBoardgameServiceServer(grpcServ, &bggSvc)

	// Start the server
	fmt.Println("Staring Boardgame Service")
	checkError(grpcServ.Serve(lis))

}
