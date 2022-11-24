package main

import (
	"context"
	"fmt"
	"io"
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
	id                int32
	clients           map[int32]auction.ReplicationClient
	primaryServer     bool
	primaryServerPort int32
	lamportTime       int32
	auction           Auction
	wg                sync.WaitGroup
	ctx               context.Context
}

func main() {
	arg1, _ := strconv.ParseInt(os.Args[1], 10, 32)
	ownPort := int32(arg1) + 6000

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	//Prints to log file and terminal
	f, err := os.OpenFile(fmt.Sprintf("logfile.%d", ownPort), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()
	mw := io.MultiWriter(os.Stdout, f)
	log.SetOutput(mw)

	p := &peer{
		id:                ownPort,
		clients:           make(map[int32]auction.ReplicationClient),
		primaryServer:     false,
		primaryServerPort: 6000,
		lamportTime:       1,
		auction: Auction{
			state:        auction.Outcome_NOTSTARTED,
			winnerId:     -1,
			winnerAmount: -1,
		},
		wg:  sync.WaitGroup{},
		ctx: ctx,
	}

	if ownPort == p.primaryServerPort {
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

	// Connect to the other replica servers (peers)
	for i := 0; i < 3; i++ {
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

	// If not a primary server, continiously ping the primary server
	go func() {
		for !p.primaryServer {
			log.Printf("Trying to ping primary server. ")
			_, err := p.clients[p.primaryServerPort].Ping(ctx, &auction.PingMessage{Id: p.id})

			// If no reply, elect new leader
			if err != nil {
				log.Printf("The king is dead. Long live the king!\n")
				p.elect()
			} else {
				// Else sleep for 2 seconds
				log.Printf("Succesfully pinged primary server.\n")
				time.Sleep(2000 * time.Millisecond)
			}
		}
	}()

	for {

	}
}

func (p *peer) elect() {
	// Delete old primary port
	delete(p.clients, p.primaryServerPort)

	// Find the replication server with the lowest id, among all alive replication servers.
	lowestId := p.id
	for id, client := range p.clients {
		reply, err := client.Ping(p.ctx, &auction.PingMessage{Id: p.id})
		if err != nil {
			fmt.Printf("Got no reply from %d\n", id)
		}
		fmt.Printf("Got reply from id %v: %v\n", id, reply.Id)
		if lowestId > reply.Id {
			lowestId = reply.Id
		}
	}
	// Declare new primary server
	log.Printf("New king of the hill is %d", lowestId)
	p.primaryServerPort = lowestId
	if lowestId == p.id {
		p.primaryServer = true
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

// Returns to the frontend whether the replication server is the primary server
func (p *peer) IsPrimaryServer(ctx context.Context, in *auction.Empty) (*auction.PrimaryServerResponse, error) {
	return &auction.PrimaryServerResponse{PrimaryServer: p.primaryServer}, nil
}

// Receives ping, and sends back a ping with the servers id
func (p *peer) Ping(ctx context.Context, in *auction.PingMessage) (*auction.PingMessage, error) {
	log.Printf("Received a ping!\n")
	return &auction.PingMessage{
		Id: p.id,
	}, nil
}

// Contains the actual logic for bidding
func (p *peer) BidBackup(ctx context.Context, in *auction.Amount) (*auction.Ack, error) {
	// If this is the primary server, it forwards the call to all the backups servers, and waits for reply
	if p.primaryServer {
		var wg sync.WaitGroup
		for id, client := range p.clients {
			wg.Add(1)
			go func(clientId int32, requestClient auction.ReplicationClient) {
				defer wg.Done()
				log.Printf("Sending request to %d", clientId)
				_, err := requestClient.BidBackup(ctx, in)
				// If a backup server has crashed, the primary server deletes it from its map of peers
				if err != nil {
					fmt.Printf("Backup server %d is not responding", clientId)
					delete(p.clients, clientId)
				}
			}(id, client)
		}
		wg.Wait()
		log.Printf("Received acknowledgements from all backups.\n")
	}

	p.IncrementLamportTime(in.LamportTime)
	log.Printf("Server has received a bid from client %d with amount %d and Lamport time %d (Lamport time %d)\n",
		in.ClientId, in.BidAmount, in.LamportTime, p.lamportTime)
	// If attempted bid is less than or equal to 0, return exception
	//time.Sleep(5000 * time.Millisecond) // Sleep timer to simulate slow program (to test for sequential consistency)
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

// Contains the actual logic for result
func (p *peer) ResultBackup(ctx context.Context, in *auction.Empty) (*auction.Outcome, error) {
	// If this is the primary server, it forwards the call to all the backups servers, and waits for reply
	if p.primaryServer {
		var wg sync.WaitGroup
		for id, client := range p.clients {
			wg.Add(1)
			go func(clientId int32, requestClient auction.ReplicationClient) {
				defer wg.Done()
				log.Printf("Sending request to %d", clientId)
				_, err := requestClient.ResultBackup(ctx, in)
				// If a backup server has crashed, the primary server deletes it from its map of peers
				if err != nil {
					fmt.Printf("Backup server %d is not responding", clientId)
					delete(p.clients, clientId)
				}
			}(id, client)
		}
		wg.Wait()
		log.Printf("Received acknowledgements from all backups.\n")
	}

	// Sends back the state of the action, stored in the auction struct
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
