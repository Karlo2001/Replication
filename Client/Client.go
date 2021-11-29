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

	Replication "github.com/Karlo2001/Replication/Proto"
	"google.golang.org/grpc"
)

var (
	serverAddr = flag.String("server_addr", "localhost:10000", "The server address in the format of host:port")
	name       string
	time       int32
)

func main() {
	flag.Parse()
	var opts []grpc.DialOption

	opts = append(opts, grpc.WithInsecure())
	conn, err := grpc.Dial(*serverAddr, opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := Replication.NewAuctionClient(conn)
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
	rBid, _ := regexp.Compile("bid \\d+")
	rAmount, _ := regexp.Compile("\\d+")
	for {
		sc := bufio.NewScanner(os.Stdin)
		if sc.Scan() {
			in = sc.Text()
			//fmt.Println(rAmount.MatchString(in))
			// To bid
			if rBid.MatchString(in) {
				amountI, err := strconv.ParseInt(rAmount.FindString(in), 10, 32)
				if err != nil {
					fmt.Println("*** Failed to bid")
				} else {
					amount := int32(amountI)
					fmt.Println(amount)
					ack, _ := client.Bid(context.Background(),
						&Replication.BidRequest{Name: name, Timestamp: time, Bid: amount})
					if ack.Ack == 1 {
						fmt.Println("Bid accepted!")
					} else if ack.Ack == 0 {
						fmt.Println("Bid is too low!")
					} else {
						fmt.Println("You can't bid on an auction which has already ended!")
					}
				}
			} else if in == "status" {
				status, err := client.Result(context.Background(), &Replication.Empty{})
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
