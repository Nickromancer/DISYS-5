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
	port              int32
	replicationClient ReplicationClient
	ctx               context.Context
}

type ReplicationClient struct {
	port       int32
	serverPort int32
	ctx        context.Context
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
		port:              ownPort,
		replicationClient: *c,
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

func (s *Frontend) Bid(ctx context.Context, in *auction.Amount) (*auction.Ack, error) {

	ack, _ := serverConnection.BidBackup(ctx, in)

	return ack, nil

}

// Function to return the current result of the auction
func (s *Frontend) Result(ctx context.Context, in *auction.Empty) (*auction.Outcome, error) {

	outcome, _ := serverConnection.ResultBackup(ctx, in)

	return outcome, nil

}

func (c *ReplicationClient) connectToServer() (auction.ReplicationClient, error) {
	conn, err := grpc.Dial(fmt.Sprintf("localhost:%d", c.serverPort), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Could not connect to port %d\n", c.serverPort)
	}
	log.Printf("Connected to server port %d\n", c.serverPort)
	return auction.NewReplicationClient(conn), nil
}
