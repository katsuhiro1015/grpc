package main

import (
	"log"
	"net"
	"time"

	pb "github.com/jun06t/grpc-sample/server-streaming/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

const (
	port = ":8443"
)

type server struct{}

func (s *server) GetNewFeed(in *pb.Empty, stream pb.Feeder_GetNewFeedServer) error {
	feed := []string{"article1", "article2", "article3"}

	for _, v := range feed {
		// １秒毎にメッセージを送信
		err := stream.Send(&pb.FeedResponse{Message: v})
		if err != nil {
			return err
		}
		time.Sleep(1 * time.Second)
	}

	// RPC終了
	return nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	creds, err := credentials.NewServerTLSFromFile(
		"../../cert/server.crt",
		"../../cert/server.key",
	)
	s := grpc.NewServer(grpc.Creds(creds))
	pb.RegisterFeederServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
