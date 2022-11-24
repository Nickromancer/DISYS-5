package main

import (
	"context"
	"fmt"
	"io"
	"os"
	"strconv"

	auction "github.com/Nickromancer/DISYS-5/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"log"
	"net"
)

type Frontend struct {
	auction.UnimplementedAuctionServer
	port               int32
	replicationClient  ReplicationClient
	replicationServers map[int32]auction.ReplicationClient
	ctx                context.Context
}

type ReplicationClient struct {
	port              int32
	primaryServerPort int32
	ctx               context.Context
}

func main() {
	arg1, _ := strconv.ParseInt(os.Args[1], 10, 32)
	arg2, _ := strconv.ParseInt(os.Args[2], 10, 32)
	ownPort := int32(arg1) + 5000
	clientPort := int32(arg2) + 6000
	serverPort := int32(6000)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	//Sets up the replication client struct
	c := &ReplicationClient{
		port:              clientPort,
		primaryServerPort: serverPort,
		ctx:               ctx,
	}

	//Sets up the frontend server struct
	s := &Frontend{
		port:               ownPort,
		replicationClient:  *c,
		replicationServers: make(map[int32]auction.ReplicationClient),
		ctx:                ctx,
	}

	//Connects to all the replication servers
	for i := 0; i < 3; i++ {
		port := int32(6000) + int32(i)
		serverConnection, err := c.connectToServer(port)
		if err != nil {
			log.Fatalf("Could not connect to the server with port %d\n", port)
		}
		s.replicationServers[port] = serverConnection
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

	//Launches the auction server (frontend)
	launchServer(s)
}

// Function for launching the auction server (frontend)
func launchServer(s *Frontend) {
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

// In case the primary server stops responding (crashes), frontend finds new primary server.
func (s *Frontend) findNewPrimaryServer() {
	delete(s.replicationServers, s.replicationClient.primaryServerPort)

	for id, client := range s.replicationServers {
		reply, err := client.IsPrimaryServer(s.ctx, &auction.Empty{})
		if err != nil {
			log.Printf("Could not ping %d\n", id)
		} else {
			if reply.PrimaryServer {
				log.Printf("The new primary server is %d\n", id)
				s.replicationClient.primaryServerPort = id
				break
			}
		}
	}
}

// Forwards the Bid request to the backend (replication servers).
// Returns the reply to the client
func (s *Frontend) Bid(ctx context.Context, in *auction.Amount) (*auction.Ack, error) {
	var ack *auction.Ack
	ack, err := s.replicationServers[s.replicationClient.primaryServerPort].BidBackup(ctx, in)

	// If no reply is recieved, find new primary server
	if err != nil {
		log.Printf("Primary server is not responding.\n")
		s.findNewPrimaryServer()
		ack, _ = s.replicationServers[s.replicationClient.primaryServerPort].BidBackup(ctx, in)
	}

	return ack, nil

}

// Forwards the Result request to the backend (replication servers).
// Returns the reply to the client
func (s *Frontend) Result(ctx context.Context, in *auction.Empty) (*auction.Outcome, error) {
	var outcome *auction.Outcome
	outcome, err := s.replicationServers[s.replicationClient.primaryServerPort].ResultBackup(ctx, in)

	// If no reply is recieved, find new primary server
	if err != nil {
		log.Printf("Primary server is not responding.\n")
		s.findNewPrimaryServer()
		outcome, _ = s.replicationServers[s.replicationClient.primaryServerPort].ResultBackup(ctx, in)
	}

	return outcome, nil
}

// Function for connecting to the Replication servers
func (c *ReplicationClient) connectToServer(connectionPort int32) (auction.ReplicationClient, error) {
	conn, err := grpc.Dial(fmt.Sprintf("localhost:%d", connectionPort), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Could not connect to port %d\n", connectionPort)
	}
	log.Printf("Connected to server port %d\n", connectionPort)
	return auction.NewReplicationClient(conn), nil
}
