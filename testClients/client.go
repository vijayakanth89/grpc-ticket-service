package main

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

	r, err := c.TicketPurchaseService(ctx, &pb.PurchaseRequestMsg{FirstName: "neethu", LastName: "twist", Email: "vijayakanth@gmail.com", TrainNo: "T003"})

	if err == nil {

		prettyJSON, err2 := json.MarshalIndent(r, "", "  ")

		if err2 == nil {

			log.Println(string(prettyJSON))
		} else {
			log.Printf(err2.Error())
		}
	} else {
		log.Printf(err.Error())
	}

	r2, err3 := c.GetReceipt(ctx, &pb.Tickets{UserId: "1", TicketNo: r.TicketNo})

	if err == nil {
		log.Printf("%#v", r2)
	} else {
		log.Println(err3)
	}

	res10, err10 := c.CancelTicket(ctx, &pb.CancelTicketRequest{TicketNo: r.TicketNo, UserId: "5"})

	if err10 == nil {
		log.Println(res10)
	} else {
		log.Fatal(err10)
	}

	r2, err3 = c.GetReceipt(ctx, &pb.Tickets{UserId: "1", TicketNo: r.TicketNo})

	if err == nil {
		log.Printf("%#v", r2)
	} else {
		log.Println(err3)
	}

	r, err = c.TicketPurchaseService(ctx, &pb.PurchaseRequestMsg{FirstName: "neethu", LastName: "twist", Email: "vijayakanth@gmail.com", TrainNo: "T003"})

	if err == nil {

		prettyJSON, err2 := json.MarshalIndent(r, "", "  ")

		if err2 == nil {

			log.Println(string(prettyJSON))
		} else {
			log.Printf(err2.Error())
		}
	} else {
		log.Printf(err.Error())
	}

	res11, _ := c.SeatReallocate(ctx, &pb.TicketEntry{TicketNo: r.TicketNo, UserId: "1"})

	log.Println(res11)

	r2, err3 = c.GetReceipt(ctx, &pb.Tickets{UserId: "1", TicketNo: r.TicketNo})

	if err == nil {
		log.Printf("%#v", r2)
	} else {
		log.Println(err3)
	}

}
