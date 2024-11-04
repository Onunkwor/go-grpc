package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/onunkwor/go-grpc/proto"
)
func callSayHelloServerStream(client pb.GreetServiceClient, names *pb.NamesList) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	stream, err := client.SayHelloServerStreaming(ctx, names)
	if err != nil {
		log.Fatalf("Error while performing server streaming: %v", err)
	}
	// Loop through the stream and read each response
	for {
		response, err := stream.Recv() // receives a stream of messages
		if err == io.EOF {
			break // ends the loop if there's no more data
		}
		if err != nil {
			log.Fatalf("Error receiving stream: %v", err)
		}
		log.Printf("Received: %s", response.Message)
	}
	
}