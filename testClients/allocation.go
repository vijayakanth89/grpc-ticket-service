package client

import (
	"context"
	"encoding/json"
	"log"
	"time"

	pb "github.com/vijayakanth89/grpc-ticket-service/ticketservice"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewTimeServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r3, err4 := c.AllocationStatus(ctx, &pb.AllocationStatusRequest{TrainNo: "T001", Section: "A"})

	if err4 == nil {
		prettyJSONAllocationStatusRes, err5 := json.MarshalIndent(r3, "", "  ")
		if err5 == nil {
			log.Println(string(prettyJSONAllocationStatusRes))
		} else {
			log.Println(err5)
		}

	} else {
		log.Println(err4)
	}

	alltickets, _ := c.GetAllTickets(ctx, &pb.DummyMessage{})

	prettyJSONAllticket, _ := json.MarshalIndent(alltickets, "", "  ")

	log.Println(string(prettyJSONAllticket))
}
