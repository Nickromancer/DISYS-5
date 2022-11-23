package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"sync"
	"time"

	auction "github.com/Nickromancer/DISYS-5/proto"
	"google.golang.org/grpc"
)

type Auction struct {
	state        auction.Outcome_STATE
	winnerId     int32
	winnerAmount int32
}

type peer struct {
	auction.UnimplementedReplicationServer
	id            int32
	clients       map[int32]auction.ReplicationClient
	primaryServer bool
	lamportTime   int32
	auction       Auction
	wg            sync.WaitGroup
	ctx           context.Context
}

func main() {
	arg1, _ := strconv.ParseInt(os.Args[1], 10, 32)
	ownPort := int32(arg1) + 6000

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	p := &peer{
		id:            ownPort,
		clients:       make(map[int32]auction.ReplicationClient),
		primaryServer: false,
		lamportTime:   1,
		auction: Auction{
			state:        auction.Outcome_NOTSTARTED,
			winnerId:     -1,
			winnerAmount: -1,
		},
		wg:  sync.WaitGroup{},
		ctx: ctx,
	}

	if ownPort == 6000 {
		p.primaryServer = true
	}
	// Create listener tcp on port ownPort
	list, err := net.Listen("tcp", fmt.Sprintf("localhost:%v", ownPort))
	if err != nil {
		log.Fatalf("Failed to listen on port: %v", err)
	}
	grpcServer := grpc.NewServer()
	auction.RegisterReplicationServer(grpcServer, p)

	go func() {
		if err := grpcServer.Serve(list); err != nil {
			log.Fatalf("failed to server %v", err)
		}
	}()

	for i := 0; i < 1; i++ {
		port := int32(6000) + int32(i)
		if port == ownPort {
			continue
		}

		var conn *grpc.ClientConn
		fmt.Printf("Trying to dial: %v\n", port)
		conn, err := grpc.Dial(fmt.Sprintf("localhost:%v", port), grpc.WithInsecure(), grpc.WithBlock())
		if err != nil {
			log.Fatalf("Could not connect: %s", err)
		}
		defer conn.Close()
		c := auction.NewReplicationClient(conn)
		p.clients[port] = c
	}

	go func() {
		for !p.primaryServer {
			state, err := p.clients[6000].Ping(ctx, &auction.Empty{})
			if err != nil {
				log.Printf("The king is dead. Long live the king!\n")
			}
			if state.Result == auction.Acknowledgement_SUCCESS {
				time.Sleep(2000 * time.Millisecond)
			} else {
				//start election
			}
		}
	}()

	for {

	}
}

// Function for incrementing lamport time
func (p *peer) IncrementLamportTime(otherLamportTime int32) {
	var mu sync.Mutex
	defer mu.Unlock()
	mu.Lock()
	if p.lamportTime < otherLamportTime {
		p.lamportTime = otherLamportTime + 1
	} else {
		p.lamportTime++
	}
}

func (p *peer) Ping(ctx context.Context, in *auction.Empty) (*auction.Acknowledgement, error) {
	log.Printf("Received a ping!\n")
	return &auction.Acknowledgement{
		Result: auction.Acknowledgement_SUCCESS,
	}, nil
}

func (p *peer) BidBackup(ctx context.Context, in *auction.Amount) (*auction.Ack, error) {

	p.IncrementLamportTime(in.LamportTime)
	log.Printf("Server has received a bid from client %d with amount %d and Lamport time %d (Lamport time %d)\n",
		in.ClientId, in.BidAmount, in.LamportTime, p.lamportTime)
	// If attempted bid is less than or equal to 0, return exception
	if in.BidAmount <= 0 {
		log.Printf("Recieved bid was less and/or equal to zero, bad bid: %s.\n", auction.Ack_EXCEPTION.String())
		return &auction.Ack{
			LamportTime: p.lamportTime,
			Result:      auction.Ack_EXCEPTION,
		}, nil
	}
	// If auction is not started, start new auction with starting bid set as highest bid
	if p.auction.state == auction.Outcome_NOTSTARTED || p.auction.state == auction.Outcome_FINISHED {
		log.Printf("Start new auction with starting bid %d and winnerId %d.\n", in.BidAmount, in.ClientId)
		p.auction.state = auction.Outcome_ONGOING
		p.auction.winnerId = in.ClientId
		p.auction.winnerAmount = in.BidAmount
		go p.StartAuction()
		return &auction.Ack{
			LamportTime: p.lamportTime,
			Result:      auction.Ack_SUCCESS,
		}, nil
	} else {
		// If auction is ongoing, check if bid is higher than current highest bid
		if p.auction.winnerAmount < in.BidAmount {
			log.Printf("New bid %d from client %d is highest bid.\n", in.BidAmount, in.ClientId)
			log.Printf("Extending auction.\n")
			go p.ExtendAuction()
			p.auction.winnerAmount = in.BidAmount
			p.auction.winnerId = in.ClientId
			return &auction.Ack{
				LamportTime: p.lamportTime,
				Result:      auction.Ack_SUCCESS,
			}, nil
		} else {
			log.Printf("You have to bid higher than the previous bidder, you donut.\n")
			return &auction.Ack{
				LamportTime: p.lamportTime,
				Result:      auction.Ack_FAIL,
			}, nil
		}
	}
}

func (p *peer) ResultBackup(ctx context.Context, in *auction.Empty) (*auction.Outcome, error) {

	log.Printf("Server received result request.\nSent result reply with outcome %s, winnerId %d and winneramount %d\n",
		p.auction.state.String(), p.auction.winnerId, p.auction.winnerAmount)
	return &auction.Outcome{
		State:      p.auction.state,
		WinnerId:   p.auction.winnerId,
		WinningBid: p.auction.winnerAmount,
	}, nil

}

// Function to start the auction with a default duration of 5 seconds
func (s *peer) StartAuction() {
	log.Printf("Auction has started! Let the bidding begin!\n")
	time.Sleep(10000 * time.Millisecond)
	s.wg.Wait()
	s.auction.state = auction.Outcome_FINISHED
	log.Printf("Auction is done!\n")
}

// Function to start the auction with a default duration of 5 seconds
func (s *peer) ExtendAuction() {
	s.wg.Add(1)
	time.Sleep(5000 * time.Millisecond)
	s.wg.Done()
}
