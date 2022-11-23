package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"

	auction "github.com/Nickromancer/DISYS-5/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client struct {
	port        int32 // We use ports for id
	serverPort  int32
	lamportTime int32
	ctx         context.Context
}

// Function for incrementing lamport time
func (c *Client) IncrementLamportTime(otherLamportTime int32) {
	var mu sync.Mutex
	defer mu.Unlock()
	mu.Lock()
	if c.lamportTime < otherLamportTime {
		c.lamportTime = otherLamportTime + 1
	} else {
		c.lamportTime++
	}
}

func main() {
	arg1, _ := strconv.ParseInt(os.Args[1], 10, 32)
	arg2, _ := strconv.ParseInt(os.Args[0], 10, 32)

	ownPort := int32(arg1) + 5000
	serverPort := int32(arg2) + 5000

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	c := &Client{
		port:        ownPort,
		serverPort:  serverPort,
		lamportTime: 1,
		ctx:         ctx,
	}

	f, err := os.OpenFile(fmt.Sprintf("logfile.%d", c.port), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v\n", err)
	}
	defer f.Close()
	mw := io.MultiWriter(os.Stdout, f)
	log.SetOutput(mw)

	serverConnection, _ := c.connectToServer()

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		input := scanner.Text()
		splitted := strings.Fields(input)
		if splitted[0] == "/bid" {
			bidAmount, _ := strconv.Atoi(splitted[1])
			log.Printf("Client %d sent a bid with amount: %d (Lamport time %d)\n", c.port, bidAmount, c.lamportTime)
			ack, err := serverConnection.Bid(c.ctx, &auction.Amount{
				LamportTime: c.lamportTime,
				ClientId:    c.port,
				BidAmount:   int32(bidAmount),
			})
			if err != nil {
				log.Fatalf("Could not send a bid to the server.\n")
			}
			log.Printf("Client %d received ack with result %s and Lamport time %d (Client Lamport time %d)\n",
				c.port, ack.Result.String(), ack.LamportTime, c.lamportTime)

		} else if splitted[0] == "/result" {
			outcome, err := serverConnection.Result(ctx, &auction.Empty{})
			if err != nil {
				log.Fatalf("Could not get a result from the server.\n")
			}
			log.Printf("The auction is currently %s.\nThe current winning client: %d with bid: %d\n", outcome.State.String(), outcome.WinnerId, outcome.WinningBid)
		}
	}
}

func (c *Client) connectToServer() (auction.AuctionClient, error) {
	conn, err := grpc.Dial(fmt.Sprintf("localhost:%d", c.serverPort), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Could not connect to port %d\n", c.serverPort)
	}
	log.Printf("Connected to server port %d\n", c.serverPort)
	return auction.NewAuctionClient(conn), nil
}
