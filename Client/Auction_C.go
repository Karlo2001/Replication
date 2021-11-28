package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"

	pb "Replication/Proto"

	"google.golang.org/grpc"
)

var (
	serverAddr = flag.String("server_addr", "localhost:10000", "The server address in the format of host:port")
	name       string
)

var time int32

func main() {
	flag.Parse()
	var opts []grpc.DialOption

	opts = append(opts, grpc.WithInsecure())
	conn, err := grpc.Dial(*serverAddr, opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := pb.NewAuctionClient(conn)

	for {
		fmt.Printf("Enter name: ")
		sc := bufio.NewScanner(os.Stdin)
		if sc.Scan() {
			name = sc.Text()
		}
		if name != "" {
			break
		}
		fmt.Println("*** Invalid name")
	}

	var in string
	rBid, _ := regexp.Compile("bid /d+")
	rAmount, _ := regexp.Compile("/d+")
	for {
		sc := bufio.NewScanner(os.Stdin)
		if sc.Scan() {
			in = sc.Text()

			// To bid
			if rBid.MatchString(in) {
				amountI, err := strconv.ParseInt(rAmount.FindString(in), 10, 32)
				if err != nil {
					fmt.Println("*** Failed to bid")
				} else {
					amount := int32(amountI)
					client.Bid(context.Background(),
						&pb.BidRequest{Name: name, Timestamp: time, Bid: amount})
				}
			} else if in == "status" {
				status, err := client.Result(context.Background(), &pb.Empty{})
				if err != nil {
					fmt.Println("*** Failed retrieve status")
				} else {
					if status.Ended {
						fmt.Println("Status: Ended")
					} else {
						fmt.Println("Status: Ongoing")
					}
					fmt.Println("Highest bidder: " + status.Name)
					fmt.Println("Bid:", status.HighestBid)
				}
			}
		}
	}
}
