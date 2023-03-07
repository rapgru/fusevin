// Package main implements a client for Greeter service.
package main

import (
	"context"
	"flag"
	"log"
	"time"
	"io"
	"bufio"
	"fmt"
	"os"

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
		log.Fatalf("could create puppet: %v", err)
	}

	log.Printf("Redirecting filename: %s", r.GetVinFilename())
	defer c.DestroyPuppet(context.Background(), &pb.DestroyPuppetRequest{Id: r.GetId()})

	stream, err := c.StartStdinNotify(context.Background(), &pb.StartStdinNotifyRequest{Id: r.GetId()})
	if err != nil {
		log.Fatalf("could not start stream: %v", err)
	}

	for i:=1; i<=2; i++ {
		_, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("client.StdinNotifyStream failed: %v", err)
		}

		log.Printf("Received StdinNotify")
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("Read occured, enter string: ")
		// text <- channel
		text, err := reader.ReadString('\n')
		if err != nil {
			log.Fatalf("unable to read from terminal")
		}

		c.SupplyStdinContent(context.Background(), &pb.StdinContent{Id: r.GetId(), Payload: text})
	}
}
