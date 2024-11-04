package main

import (
	"context"
	"log"
	"time"

	pb "github.com/onunkwor/go-grpc/proto"
)
func callSayHelloClientStream(client pb.GreetServiceClient, names *pb.NamesList) {
  log.Printf("CLient streaming started")
  ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
  stream, err := client.SayHelloClientStreaming(ctx, )
  if err != nil {
	log.Fatalf("Could not send name: %v", err)
  }
  for _, name := range names.Names {
	req := &pb.HelloRequest{
		Name: name,
	}
	if err = stream.Send(req); err != nil {
		log.Fatalf("Error while sending name: %v", err)
	}
	log.Printf("Sent the request with name: %v", name)

	time.Sleep(2*time.Second)
  }
  res, err := stream.CloseAndRecv()
    if err != nil {
        log.Fatalf("Error while receiving response: %v", err)
    }
    log.Printf("%v", res.Names)
}