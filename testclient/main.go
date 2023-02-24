// Package main implements a client for Greeter service.
package main

import (
	"context"
	"flag"
	"log"
	"time"
	"io"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "tuwien.ac.at/fusevin/vin"
)

const (
	// defaultName = "world"
)

var (
	addr = flag.String("addr", "127.0.0.1:50051", "the address to connect to")
)

func main() {
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewFuseVinClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.CreatePuppet(ctx, &pb.Empty{})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Redirecting filename: %s", r.GetVinFilename())

	ctx = context.Background()
	stream, err := c.StartStdinNotify(ctx, &pb.StartStdinNotifyRequest{Id: r.GetId()})
	if err != nil {
		log.Fatalf("could not start stream: %v", err)
	}

	for {
		_, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("client.StdinNotifyStream failed: %v", err)
		}

		log.Printf("Received StdinNotify")
	}
}
