package main

import (
	"log"
	"net"

	db "github.com/vijayakanth89/grpc-ticket-service/datastore"
	pb "github.com/vijayakanth89/grpc-ticket-service/ticketservice"
	"google.golang.org/grpc"
)

func main() {

	db.TestPurchase(10)

	lis, err := net.Listen("tcp", pb.SERVER_PORT)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterTimeServiceServer(s, &pb.TicketServer{})

	log.Printf("gRPC server listening on port %s", pb.SERVER_PORT)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
