package main

import (
	"io"
	"log"

	pb "github.com/onunkwor/go-grpc/proto"
)
func (s *helloServer) SayHelloClientStreaming(stream pb.GreetService_SayHelloClientStreamingServer) error{
	var messages []string
	for {
		response, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.MessageList{Names: messages})
		}
		if err != nil {
			log.Fatalf("Error receiving stream: %v", err)
		}
		log.Printf("Got messages of names %s",response.Name)
		messages = append(messages,"Hello "+ response.Name)
	}
}