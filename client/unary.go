package main

import (
	"context"
	"log"
	"time"

	pb "github.com/onunkwor/go-grpc/proto"
)
func callSayHello(client pb.GreetServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	res, err := client.SayHello(ctx, &pb.NoParam{})
	if err!=nil {
		log.Fatalf("Failed to say hello: %v", err)
	}
	log.Printf("%s",res.Message)
}