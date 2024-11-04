package main

import (
	"log"
	"net"

	pb "github.com/onunkwor/go-grpc/proto"
	"google.golang.org/grpc"
)

const port = ":8080" // Add colon before port for TCP connections
type helloServer struct{
	pb.GreetServiceServer
}
func main() {
	// Declare err once
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to start the server: %v", err)
	}
	
	grpcServer := grpc.NewServer()
	pb.RegisterGreetServiceServer(grpcServer, &helloServer{})
	// Reuse the declared err variable here without redeclaring
	log.Printf("server started at %v", lis.Addr())
	if err = grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to start: %v", err)
	}
}
