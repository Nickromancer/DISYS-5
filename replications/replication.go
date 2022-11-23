package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"

	auction "github.com/Nickromancer/DISYS-5/proto"
	"google.golang.org/grpc"
)

func main() {
	arg1, _ := strconv.ParseInt(os.Args[1], 10, 32)
	ownPort := int32(arg1) + 5000

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	p := &peer{
		id:      ownPort,
		clients: make(map[int32]auction.ReplicationClient),
		ctx:     ctx,
	}

	// Create listener tcp on port ownPort
	list, err := net.Listen("tcp", fmt.Sprintf(":%v", ownPort))
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

	for i := 0; i < 3; i++ {
		port := int32(6000) + int32(i)

		if port == ownPort {
			continue
		}

		var conn *grpc.ClientConn
		fmt.Printf("Trying to dial: %v\n", port)
		conn, err := grpc.Dial(fmt.Sprintf(":%v", port), grpc.WithInsecure(), grpc.WithBlock())
		if err != nil {
			log.Fatalf("Could not connect: %s", err)
		}
		defer conn.Close()
		c := auction.NewReplicationClient(conn)
		p.clients[port] = c
	}
}

type peer struct {
	auction.UnimplementedReplicationServer
	id            int32
	clients       map[int32]auction.ReplicationClient
	primaryServer bool
	ctx           context.Context
}

func (p *peer) Ping(ctx context.Context, in *auction.Empty) (*auction.Acknowledgement, error) {
	log.Printf("Trying to ping primary replication server")
}

func (p *peer) BidBackup(ctx context.Context, in *auction.Amount) (*auction.AckReply, error) {
}

func (p *peer) Result(ctx context.Context, in *auction.Empty) (*auction.OutcomeReply, error) {
}
