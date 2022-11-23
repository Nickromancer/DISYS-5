package main

import (
	"context"
	"fmt"
	"io"
	"os"
	"strconv"
	"sync"

	auction "github.com/Nickromancer/DISYS-5/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"log"
	"net"
)

type Frontend struct {
	auction.UnimplementedAuctionServer
	port              int32
	auction           Auction
	replicationClient ReplicationClient
	wg                sync.WaitGroup
	ctx               context.Context
}

type ReplicationClient struct {
	port       int32
	serverPort int32
	ctx        context.Context
}

type Auction struct {
	state        auction.Outcome_STATE
	winnerId     int32
	winnerAmount int32
}

var serverConnection auction.ReplicationClient

func main() {
	arg1, _ := strconv.ParseInt(os.Args[1], 10, 32)
	arg2, _ := strconv.ParseInt(os.Args[2], 10, 32)
	ownPort := int32(arg1) + 5000
	clientPort := int32(arg2) + 6000
	serverPort := int32(6000)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	c := &ReplicationClient{
		port:       clientPort,
		serverPort: serverPort,
		ctx:        ctx,
	}

	s := &Frontend{
		port: ownPort,
		auction: Auction{
			state:        auction.Outcome_NOTSTARTED,
			winnerId:     -1,
			winnerAmount: -1,
		},
		replicationClient: *c,
		wg:                sync.WaitGroup{},
		ctx:               ctx,
	}

	serverConnection, _ = c.connectToServer()

	//Prints to log file and terminal
	f, err := os.OpenFile(fmt.Sprintf("logfile.%d", s.port), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()
	mw := io.MultiWriter(os.Stdout, f)
	log.SetOutput(mw)

	log.Printf("Server is starting\n")

	launchServer(s)
}

/*
// Function for incrementing lamport time

	func (s *Server) IncrementLamportTime(otherLamportTime int32) {
		var mu sync.Mutex
		defer mu.Unlock()
		mu.Lock()
		if s.lamportTime < otherLamportTime {
			s.lamportTime = otherLamportTime + 1
		} else {
			s.lamportTime++
		}
	}
*/
func launchServer(s *Frontend) {
	grpcServer := grpc.NewServer()

	listener, err := net.Listen("tcp", fmt.Sprintf("192.168.26.195:%d", s.port))

	if err != nil {
		log.Fatalf("Could not create the server %v\n", err)
	}
	log.Printf("Started server at port %d\n", s.port)

	auction.RegisterAuctionServer(grpcServer, s)

	serveError := grpcServer.Serve(listener)
	if serveError != nil {
		log.Fatalf("Could not serve listener\n")
	}
}

func (s *Frontend) Bid(ctx context.Context, in *auction.Amount) (*auction.Ack, error) {

	ackReply, _ := serverConnection.BidBackup(ctx, in)

	return ackReply.Ack, nil
	/*
		s.IncrementLamportTime(in.LamportTime)
		log.Printf("Server has received a bid from client %d with amount %d and Lamport time %d (Lamport time %d)\n",
		in.ClientId, in.BidAmount, in.LamportTime, s.lamportTime)
		// If attempted bid is less than or equal to 0, return exception
		if in.BidAmount <= 0 {
			log.Printf("Recieved bid was less and/or equal to zero, bad bid: %s.\n", auction.Ack_EXCEPTION.String())
			return &auction.Ack{
				LamportTime: s.lamportTime,
				Result:      auction.Ack_EXCEPTION,
			}, nil
		}
		// If auction is not started, start new auction with starting bid set as highest bid
		if s.auction.state == auction.Outcome_NOTSTARTED || s.auction.state == auction.Outcome_FINISHED {
			log.Printf("Start new auction with starting bid %d and winnerId %d.\n", in.BidAmount, in.ClientId)
			s.auction.state = auction.Outcome_ONGOING
			s.auction.winnerId = in.ClientId
			s.auction.winnerAmount = in.BidAmount
			go s.StartAuction()
			return &auction.Ack{
				LamportTime: s.lamportTime,
				Result:      auction.Ack_SUCCESS,
			}, nil
		} else {
			// If auction is ongoing, check if bid is higher than current highest bid
			if s.auction.winnerAmount < in.BidAmount {
				log.Printf("New bid %d from client %d is highest bid.\n", in.BidAmount, in.ClientId)
				log.Printf("Extending auction.\n")
				go s.ExtendAuction()
				s.auction.winnerAmount = in.BidAmount
				s.auction.winnerId = in.ClientId
				return &auction.Ack{
					LamportTime: s.lamportTime,
					Result:      auction.Ack_SUCCESS,
				}, nil
			} else {
				log.Printf("You have to bid higher than the previous bidder, you donut.\n")
				return &auction.Ack{
					LamportTime: s.lamportTime,
					Result:      auction.Ack_FAIL,
				}, nil
			}
		}
	*/
}

// Function to return the current result of the auction
func (s *Frontend) Result(ctx context.Context, in *auction.Empty) (*auction.Outcome, error) {
	log.Printf("Server received result request.\nSent result reply with outcome %s, winnerId %d and winneramount %d\n",
		s.auction.state.String(), s.auction.winnerId, s.auction.winnerAmount)

	outcomeReply, _ := serverConnection.Result(ctx, in)

	return outcomeReply.Outcome, nil
	/*
		return &auction.Outcome{
			State:      s.auction.state,
			WinnerId:   s.auction.winnerId,
			WinningBid: s.auction.winnerAmount,
		}, nil
	*/
}

/*
// Function to start the auction with a default duration of 5 seconds
func (s *Server) StartAuction() {
	log.Printf("Auction has started! Let the bidding begin!\n")
	time.Sleep(10000 * time.Millisecond)
	s.wg.Wait()
	s.auction.state = auction.Outcome_FINISHED
	log.Printf("Auction is done!\n")
}

// Function to start the auction with a default duration of 5 seconds
func (s *Server) ExtendAuction() {
	s.wg.Add(1)
	time.Sleep(5000 * time.Millisecond)
	s.wg.Done()
}
*/

func (c *ReplicationClient) connectToServer() (auction.ReplicationClient, error) {
	conn, err := grpc.Dial(fmt.Sprintf("192.168.26.195:%d", c.serverPort), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Could not connect to port %d\n", c.serverPort)
	}
	log.Printf("Connected to server port %d\n", c.serverPort)
	return auction.NewReplicationClient(conn), nil
}
