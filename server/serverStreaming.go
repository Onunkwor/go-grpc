package main

import (
	"log"
	"time"

	pb "github.com/onunkwor/go-grpc/proto"
)

func (s *helloServer) SayHelloServerStreaming(req *pb.NamesList, stream pb.GreetService_SayHelloServerStreamingServer)  error {
   log.Printf("got request with names: %v", req.Names)
   for _, name := range req.Names{
    res := &pb.HelloResponse{
		Message: "Hello " + name,
	}

	// Send each response message to the client
	if err := stream.Send(res); err != nil {
		return err
	}
	time.Sleep(2*time.Second)
   }
   return nil // End of the stream
}