package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"net"
	"os"
	"strconv"
	"time"

	Replication "github.com/Karlo2001/Replication/Proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
)

var (
	//port       = flag.Int("port", 10000, "The server port")
	Id           int32
	Timestamp    int32
	queue        = make([]client, 0)
	inProgress   = false
	leader       = false
	votes        = 0
	voted        = false
	leaderId     int32
	serverQueue  = make([]int32, 0)
	votedForId   = 0
	recievedPing = false
	clientNames  []string
	clientBids   []int32
	serverCon    []*grpc.ClientConn
)

type server struct {
	Replication.UnimplementedAuctionServer
}

type client struct {
	name string
	bid  int32
	//time int32
}
type SecondsTimer struct {
	timer *time.Timer
	end   time.Time
}

func NewSecondsTimer(t time.Duration) *SecondsTimer {
	return &SecondsTimer{time.NewTimer(t), time.Now().Add(t)}
}

func (s *SecondsTimer) Reset(t time.Duration) {
	s.timer.Reset(t)
	s.end = time.Now().Add(t)
}

func (s *SecondsTimer) Stop() {
	s.timer.Stop()
}

//Maybe good idea to add this as a critical section
var clients []client
var serverPorts []*string
var topBid client
var servers []Replication.AuctionClient
var serverIds []string
var timer *SecondsTimer

func (s *server) Bid(ctx context.Context, bid *Replication.BidRequest) (*Replication.Ack, error) {
	queue = append(queue, client{name: bid.Name, bid: bid.Bid})
	for {
		if leader {
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

					for _, s := range clients {
						clientNames = append(clientNames, s.name)
						clientBids = append(clientBids, s.bid)
					}
					for _, s := range servers {
						s.Update(context.Background(), &Replication.UpdateMessage{HighestBidName: topBid.name, HighestBid: topBid.bid, AllBiddersNames: clientNames, AllBids: clientBids, Timestamp: Timestamp, Ping: false})
					}
					clientNames = clientNames[0:]
					clientBids = clientBids[0:]
					return &Replication.Ack{Ack: 1, Timestamp: Timestamp}, nil
				}
				queue = queue[1:]
				return &Replication.Ack{Ack: 0, Timestamp: Timestamp}, nil
			}
		} else {
			for i, s := range serverIds {
				id, _ := strconv.Atoi(s)
				if int32(id) == leaderId {
					servers[i].Bid(context.Background(), bid)
				}
			}
		}
	}
}

func (s *server) Result(ctx context.Context, empty *Replication.Empty) (*Replication.Outcome, error) {
	if !inProgress {
		return &Replication.Outcome{HighestBid: topBid.bid, Name: topBid.name, Timestamp: Timestamp, Ended: true}, nil
	}
	return &Replication.Outcome{HighestBid: topBid.bid, Name: topBid.name, Timestamp: Timestamp, Ended: false}, nil
}

func (s *server) VoteForMe(ctx context.Context, serverInfo *Replication.ServerInfo) (*Replication.Ack, error) {
	serverQueue = append(serverQueue, serverInfo.Id)
	for {
		if serverQueue[0] == serverInfo.Id {
			if voted {
				//votedForId = int(serverInfo.Id)
				serverQueue = serverQueue[1:]
				return &Replication.Ack{Ack: 0, Timestamp: Timestamp}, nil
			}
			if votedForId == 0 {
				voted = true
				votedForId = int(serverInfo.Id)
				log.Println("Voted for: " + strconv.Itoa(votedForId))
				serverQueue = serverQueue[1:]
				return &Replication.Ack{Ack: 1, Timestamp: Timestamp}, nil
			}
			//votes--
			serverQueue = serverQueue[1:]
			return &Replication.Ack{Ack: -1, Timestamp: Timestamp}, nil
		}
	}
}

func (s *server) NewLeader(ctx context.Context, serverInfo *Replication.ServerInfo) (*Replication.Ack, error) {
	leaderId = serverInfo.Id
	votes = 0
	voted = false
	leader = false
	log.Println("New leader has id: " + strconv.Itoa(int(leaderId)))
	go ping()
	return &Replication.Ack{Ack: 1, Timestamp: Timestamp}, nil
}

func (s *server) Update(ctx context.Context, Update *Replication.UpdateMessage) (*Replication.Ack, error) {
	if Update.Ping {
		//log.Println("I GOT PINGED")
		timer.Reset(time.Millisecond * time.Duration(rand.Intn(2000-1000)+1000))
		recievedPing = true
		return &Replication.Ack{Ack: 1, Timestamp: Timestamp}, nil
	}
	topBid.bid = Update.HighestBid
	topBid.name = Update.HighestBidName
	clients = clients[0:]
	for i, c := range Update.AllBids {
		clients = append(clients, client{bid: c, name: Update.AllBiddersNames[i]})
	}
	log.Println(Update.HighestBid)
	return &Replication.Ack{Ack: 1, Timestamp: Timestamp}, nil
}

func ping() {
	timer = NewSecondsTimer(time.Second * 2)
	for {
		//log.Println(leader)
		if leader {
			for _, s := range servers {
				ct, _ := context.WithTimeout(context.Background(), time.Duration(time.Millisecond*200))
				s.Update(ct, &Replication.UpdateMessage{Ping: true})
			}
			time.Sleep(time.Duration(time.Millisecond * 50))
		} else {
			//rand.Seed(time.Now().UnixNano())
			//time.Sleep(time.Duration(time.Millisecond * time.Duration())
			if recievedPing {
				recievedPing = false
				continue
			}
			//log.Println(timer.TimeRemaining())
			if timer.TimeRemaining() <= 0 {
				newElection()
			}
		}
	}
}
func (s *SecondsTimer) TimeRemaining() time.Duration {
	return s.end.Sub(time.Now())
}

func main() {
	flag.Parse()
	if len(os.Args) < 5 {
		log.Println("Please specify the addresses and an id")
		os.Exit(1)
	}

	nodeid, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Println("The specified id is not of type int. Please try again!")
		os.Exit(1)
	}
	Id = int32(nodeid)
	port, _ := strconv.Atoi(string(os.Args[2][len(os.Args[2])-4:]))
	idport := flag.Int("port", port, "The nodes port")
	fmt.Println(os.Args[2][len(os.Args[2])-4:])
	var counter = 0
	for _, c := range os.Args[3:] {
		serverIds = append(serverIds, string(c[len(c)-4]))
		serverPorts = append(serverPorts, flag.String("Server port: "+string(counter), "localhost"+string(c[len(c)-5:]), "Another servers port"))
		counter++
	}
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *idport))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption

	grpcServer := grpc.NewServer(opts...)
	Replication.RegisterAuctionServer(grpcServer, &server{})
	topBid.bid = 0

	go connectServers()
	grpcServer.Serve(lis)
}

func connectServers() {
	for _, idport := range serverPorts {
		log.Println("Searching on port " + *idport)
		conn, _ := grpc.Dial(*idport, grpc.WithInsecure(), grpc.WithBlock())
		c := Replication.NewAuctionClient(conn)
		servers = append(servers, c)
		serverCon = append(serverCon, conn)
		log.Println(*idport + " connected")
	}
	rand.Seed(time.Now().UnixNano())
	time.Sleep(time.Duration(time.Millisecond * time.Duration(rand.Intn(300-150)+150)))
	go ping()
	//newElection()
	inProgress = true
}

func newElection() {
	rand.Seed(time.Now().UnixNano())
	time.Sleep(time.Duration(time.Millisecond * time.Duration(rand.Intn(1000-500)+500)))
	leaderId = 0
	votedForId = 0
	if !voted {
		voted = true
		/*votes++
		votedForId = int(Id)
		log.Println("Id: " + strconv.Itoa(int(Id)) + " voted for self")*/
	}
	for i, s := range servers {
		//ct, _ := context.WithTimeout(context.Background(), time.Duration(time.Millisecond*200))
		//_, err := grpc.Dial(*serverPorts[i], grpc.WithInsecure())

		if serverCon[i].GetState() == connectivity.Ready {

			response, _ := s.VoteForMe(context.Background(), &Replication.ServerInfo{Id: Id, Timestamp: Timestamp})
			log.Println("Id: " + strconv.Itoa(int(Id)) + " got vote: " + strconv.Itoa(int(response.Ack)))
			if response.Ack == 1 {
				votes++
			}
		}
	}
	log.Println(votes)
	if votes+1 >= len(servers)/2+1 {
		for i, s := range servers {
			if serverCon[i].GetState() == connectivity.Ready {
				ct, _ := context.WithTimeout(context.Background(), time.Duration(time.Millisecond*200))
				s.NewLeader(ct, &Replication.ServerInfo{Id: Id, Timestamp: Timestamp})
				leaderId = Id
				leader = true
				go ping()
			}
		}
		return
	}
	if leaderId == 0 {
		voted = false
		votes = 0
		votedForId = 0
		newElection()
	}
}
