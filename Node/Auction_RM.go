package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	pb "Replication/Proto"

	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 10000, "The server port")
	id   int32
)

type auctionServer struct {
	pb.UnimplementedAuctionServer
}

type client struct {
	name string
	bid  int32
	time int32
}

var clients []client
var topBid client

func (s *auctionServer) Bid(ctx context.Context, bid *pb.BidRequest) (*pb.Ack, error) {

	return &pb.Ack{}, nil
}

func (s *auctionServer) Result(ctx context.Context, empty *pb.Empty) (*pb.Outcome, error) {

	return &pb.Outcome{}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption

	grpcServer := grpc.NewServer(opts...)
	pb.RegisterAuctionServer(grpcServer, &auctionServer{})
	grpcServer.Serve(lis)
}
