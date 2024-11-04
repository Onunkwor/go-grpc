package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/onunkwor/go-grpc/proto"
)

func callSayHelloBidirectionalStreaming(client pb.GreetServiceClient, names *pb.NamesList) {
	log.Printf("Client streaming started")
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	// Start the bidirectional streaming RPC
	stream, err := client.SayHelloBidirectionalStreaming(ctx)
	if err != nil {
		log.Fatalf("Could not start bidirectional streaming: %v", err)
	}

	// Goroutine to send names to the server
	go func() {
		for _, name := range names.Names {
			req := &pb.HelloRequest{
				Name: name,
			}

			// Send the request to the server
			if err := stream.Send(req); err != nil {
				log.Fatalf("Error while sending name: %v", err)
			}
			log.Printf("Sent name: %v", name)

			// Simulate some delay between sending each name
			time.Sleep(2 * time.Second)
		}
		// Close the stream after sending all names
		stream.CloseSend()
	}()

	// Receive responses from the server
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			// Server has closed the stream
			log.Println("Server has finished sending responses")
			break
		}
		if err != nil {
			log.Fatalf("Error while receiving response: %v", err)
		}
		// Log the response message
		log.Printf("Received response: %v", res.Message)
	}
	log.Println("Bidirectional streaming finished")
}
