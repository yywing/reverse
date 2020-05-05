package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	pb "reverse/proto"
	"reverse/rpc"

	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 10001, "The server port")
)

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()

	pb.RegisterReverseServer(grpcServer, rpc.NewReverseService())
	grpcServer.Serve(lis)
}
