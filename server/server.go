package main

import (
	"context"
	"fmt"
	"io"
	"os"
	"strconv"
	"sync"
	"time"

	auction "github.com/Nickromancer/DISYS-5/proto"
	"google.golang.org/grpc"

	"log"
	"net"
)

type Server struct {
	auction.UnimplementedAuctionServer
	port        int32
	lamportTime int32
	auction     Auction
	ctx         context.Context
}

type Auction struct {
	state        auction.Outcome_STATE
	winnerId     int32
	winnerAmount int32
}

func main() {
	arg1, _ := strconv.ParseInt(os.Args[1], 10, 32)
	ownPort := int32(arg1) + 5000

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	s := &Server{
		port:        ownPort,
		lamportTime: 1,
		auction: Auction{
			state:        auction.Outcome_NOTSTARTED,
			winnerId:     -1,
			winnerAmount: -1,
		},
		ctx: ctx,
	}

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

func launchServer(s *Server) {
	grpcServer := grpc.NewServer()

	listener, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", s.port))

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

func (s *Server) Bid(ctx context.Context, in *auction.Amount) (*auction.Ack, error) {
	s.IncrementLamportTime(in.LamportTime)
	// If attempted bid is less than or equal to 0, return exception
	if in.BidAmount <= 0 {
		return &auction.Ack{
			LamportTime: s.lamportTime,
			Result:      auction.Ack_EXCEPTION,
		}, nil
	}
	// If auction is not started, start new auction with starting bid set as highest bid
	if s.auction.state == auction.Outcome_NOTSTARTED || s.auction.state == auction.Outcome_FINISHED {
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
			s.auction.winnerAmount = in.BidAmount
			s.auction.winnerId = in.ClientId
			return &auction.Ack{
				LamportTime: s.lamportTime,
				Result:      auction.Ack_SUCCESS,
			}, nil
		} else {
			return &auction.Ack{
				LamportTime: s.lamportTime,
				Result:      auction.Ack_FAIL,
			}, nil
		}
	}
}

// Function to return the current result of the auction
func (s *Server) Result(ctx context.Context, in *auction.Empty) (*auction.Outcome, error) {
	log.Printf("Server received result request.")
	return &auction.Outcome{
		State:      s.auction.state,
		WinnerId:   s.auction.winnerId,
		WinningBid: s.auction.winnerAmount,
	}, nil
}

// Function to start the auction with a default duration of 5 seconds
func (s *Server) StartAuction() {
	time.Sleep(5000 * time.Millisecond)
	s.auction.state = auction.Outcome_FINISHED
}
