package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"

	Replication "github.com/Karlo2001/Replication/Proto"
	"google.golang.org/grpc"
)

var (
	//port       = flag.Int("port", 10000, "The server port")
	Id         int32
	Timestamp  int32
	queue      = make([]client, 0)
	inProgress = true
)

type server struct {
	Replication.UnimplementedAuctionServer
}

type client struct {
	name string
	bid  int32
	//time int32
}

//Maybe good idea to add this as a critical section
var clients []client
var topBid client

func (s *server) Bid(ctx context.Context, bid *Replication.BidRequest) (*Replication.Ack, error) {
	queue = append(queue, client{name: bid.Name, bid: bid.Bid})
	for {
		if queue[0].name == bid.Name {
			if !inProgress {
				queue = queue[1:]
				return &Replication.Ack{Ack: -1, Timestamp: Timestamp}, nil
			}
			prevBidder := false
			for _, c := range clients {
				if c == queue[0] {
					prevBidder = true
					break
				}
			}
			if !prevBidder {
				clients = append(clients, queue[0])
			}
			if queue[0].bid > topBid.bid {
				topBid = queue[0]
				queue = queue[1:]
				return &Replication.Ack{Ack: 1, Timestamp: Timestamp}, nil
			}
			queue = queue[1:]
			return &Replication.Ack{Ack: 0, Timestamp: Timestamp}, nil
		}
	}
}

func (s *server) Result(ctx context.Context, empty *Replication.Empty) (*Replication.Outcome, error) {
	if !inProgress {
		return &Replication.Outcome{HighestBid: topBid.bid, Name: topBid.name, Timestamp: Timestamp, Ended: true}, nil
	}
	return &Replication.Outcome{HighestBid: topBid.bid, Name: topBid.name, Timestamp: Timestamp, Ended: false}, nil
}

func main() {
	flag.Parse()
	listenPort, _ := strconv.Atoi(os.Args[1])
	port := flag.Int("port", listenPort, "The server port")
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption

	grpcServer := grpc.NewServer(opts...)
	Replication.RegisterAuctionServer(grpcServer, &server{})
	inProgress = true
	topBid.bid = 0
	grpcServer.Serve(lis)
}
