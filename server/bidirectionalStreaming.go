package main

import (
	"io"
	"log"
	"time"

	pb "github.com/onunkwor/go-grpc/proto"
)

func (s *helloServer) SayHelloBidirectionalStreaming (stream pb.GreetService_SayHelloBidirectionalStreamingServer) error{
	log.Println("Started bidirectional streaming")
for {
	// Receive a message from the client
	req, err := stream.Recv()
	if err == io.EOF {
		// If the client has finished sending, end the response stream
		log.Println("Client has finished sending data")
		break
	}
	if err != nil {
		log.Fatalf("Error receiving data from client: %v", err)
		return err
	}
	// Log the received name from the client
	log.Printf("Received name from client: %s", req.Name)

	// Create a response
	res := &pb.HelloResponse{
		Message: "Hello " + req.Name,
	}
	// Send the response back to the client
	if err := stream.Send(res); err != nil {
		log.Fatalf("Error sending data to client: %v", err)
		return err
	}
// Simulate some processing delay
time.Sleep(2 * time.Second)
}
log.Println("Finished bidirectional streaming")
	return nil
}