package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"io"
	"log"

	pb "github.com/jun06t/grpc-sample/server-streaming/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

const (
	address     = "localhost:8443"
	defaultName = "world"
)

func main() {
	// Set up a connection to the server.
	creds := credentials.NewTLS(&tls.Config{InsecureSkipVerify: true})
	//certFile := "../../cert/server.crt"
	//creds, err := credentials.NewClientTLSFromFile(certFile, "")

	//conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewFeederClient(conn)

	stream, err := client.GetNewFeed(context.Background(), new(pb.Empty))
	if err != nil {
		log.Fatal(err)
	}
	for {
		article, err := stream.Recv()
		// RPCの終了を検知
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(article)
	}
}
