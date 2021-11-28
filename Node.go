package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"

	Replication "github.com/Karlo2001/Replication/Proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type ack int32

const (
	Exception ack = -1
	Fail      ack = 0
	Success   ack = 1
)

type Node struct {
	Id        int32
	Timestamp int32
	Replication.UnimplementedReplicationServiceServer
}

var clients []Replication.ReplicationServiceClient
var clientports []*string

func main() {
	//reader := bufio.NewReader(os.Stdin)
	//nodeid, _ := reader.ReadString('\n')
	if len(os.Args) < 5 {
		log.Println("Please specify the addresses and an id")
		os.Exit(1)
	}

	nodeid, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Println("The specified id is not of type int. Please try again!")
		os.Exit(1)
	}
	port, _ := strconv.Atoi(string(os.Args[2][len(os.Args[2])-4:]))
	idport := flag.Int("port", port, "The nodes port")
	fmt.Println(*idport)
	var counter = 0
	for _, c := range os.Args[3:] {
		clientports = append(clientports, flag.String("client port"+string(counter), "localhost"+string(c[len(c)-5:]), "Another clients port"))
		counter++
	}
	//nodeid := 9000
	//idport := ":9000"

	//Make the server
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *idport))
	if err != nil {
		log.Fatalf("Failed to listen on port %d: %v", nodeid, err)
	}
	s := Node{Id: int32(nodeid), Timestamp: 0}
	grpcServer := grpc.NewServer()
	Replication.RegisterReplicationServiceServer(grpcServer, &s)

	//Wait on all servers to join
	//time.Sleep(time.Duration(17) * time.Second)

	go client(&s)
	grpcServer.Serve(lis)
	/*if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC server over port 9000: %v", err)
	}*/

}

func client(s *Node) {

	//Make the client
	createclient()
	//defer conn.Close()
}

func createclient() {
	var idports []*string

	//idports = append(idports, clientport1)
	//idports = append(idports, clientport2)
	idports = append(idports, clientports...)
	for _, idport := range idports {
		log.Println("Searching on port " + *idport)
		conn, _ := grpc.Dial(*idport, grpc.WithInsecure(), grpc.WithBlock())
		c := Replication.NewReplicationServiceClient(conn)
		clients = append(clients, c)
		log.Println("Port " + *idport + " connected")
	}
}

/*func funkygine(s *Node) {
	//--Send request to all Nodes in the system

	responses := make(map[DistributedMutualExclusion.CommunicationServiceClient]*DistributedMutualExclusion.Response)
	for _, c := range clients {
		response, errr := c.RequestAccess(context.Background(), &request)
		if errr != nil {
			log.Fatalf("Error when calling RequestAccess: %s", errr)
		}
		responses[c] = response
	}

	//--Wait until you have received all responses
	//--Look though all validate
	//--If zero -> Continue
	//--If One -> send AccessGranted() to these nodes when entering CriticalSection

	accepted := true
	for _, response := range responses {
		if response.Validate == 2 {
			accepted = false
		}
	}

	// Wait until you have received an AccessGranted from another node, and then send another request with the same timestamp
	if !accepted {
		for {
			time.Sleep(time.Second * time.Duration(1))
			if s.Status == "TRYAGAIN" {
				break
			}
		}
		funkygine(s)
	}
}*/

func (s *Node) Bid(ctx context.Context, request *Replication.BidRequest) (*Replication.Ack, error) {
	return &Replication.Ack{}, nil
}

func (s *Node) Result(ctx context.Context, empty *Replication.Empty) (*Replication.Outcome, error) {
	return &Replication.Outcome{}, nil
}
