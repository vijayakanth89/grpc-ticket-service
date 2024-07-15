package ticketservice

import (
	context "context"
	"errors"
	"fmt"
	"strconv"

	db "github.com/vijayakanth89/grpc-ticket-service/datastore"
)

const (
	SERVER_PORT = ":50051"
)

type TicketServer struct {
	UnimplementedTimeServiceServer
}

func (s *TicketServer) SeatReallocate(ctx context.Context, in *TicketEntry) (*TicketReallocResMsg, error) {

	ticket, err := db.Tickets.Exists(in.TicketNo)

	oldSeatNumber := ticket.GetSeatNumber()

	if err != nil {
		return &TicketReallocResMsg{}, err
	}

	err1 := ticket.SeatReallocate()

	if err1 != nil {
		return &TicketReallocResMsg{}, err1
	}

	return &TicketReallocResMsg{NewSeatNo: ticket.GetSeatNumber(), OldSeatNo: oldSeatNumber}, nil
}

func (s *TicketServer) CancelTicket(ctx context.Context, in *CancelTicketRequest) (*CancelTicketResponse, error) {
	ticket, err := db.Tickets.Exists(in.TicketNo)

	if err != nil {
		return &CancelTicketResponse{Status: "F"}, err
	}

	err = ticket.Cancel()

	if err != nil {
		return &CancelTicketResponse{Status: "F"}, err
	}

	delete(db.Tickets.Map, ticket.Number)

	return &CancelTicketResponse{Status: "S"}, nil
}

func (s *TicketServer) AllocationStatus(ctx context.Context, in *AllocationStatusRequest) (*AllocationStatusResponse, error) {
	var t *db.Train
	found := false
	for _, train := range db.AvailableTrains {

		if in.TrainNo == train.Number {
			t = &train
			found = true
			break
		}
	}
	if !found {
		return &AllocationStatusResponse{}, errors.New("train not found")
	}

	var m *map[int]string

	if in.Section == "A" {
		m = &t.SectionA.M
	} else if in.Section == "B" {
		m = &t.SectionB.M
	} else {
		return &AllocationStatusResponse{}, errors.New("invalid section")
	}

	res := AllocationStatusResponse{
		TrainNo: in.TrainNo,
		Section: in.Section,
		Entries: []*SeatEntry{},
	}

	for seatNo, userId := range *m {
		userObj, err := db.Users.GetUser(userId)

		userInfo := &User{UserId: userObj.Id}
		if err == nil {
			userInfo.Email = userObj.EmailId

			entry := &SeatEntry{SeatNumber: strconv.Itoa(seatNo), UserInfo: userInfo}
			res.Entries = append(res.Entries, entry)
		} else {
			fmt.Println(err)
		}

	}

	return &res, nil
}

func (s *TicketServer) TicketPurchaseService(ctx context.Context, in *PurchaseRequestMsg) (*TicketPurchaseResMsg, error) {
	//&pb.TicketPurchaseResMsg{FirstName: "hello", LastName: "dude"}, nil

	ticket, err := db.PurchaseTicket(in.FirstName, in.LastName, in.Email, in.TrainNo)

	if err == nil {

		return &TicketPurchaseResMsg{FirstName: ticket.UserObj.FirstName,
			LastName: ticket.UserObj.LastName,
			TicketNo: ticket.Number,
			Email:    ticket.UserObj.EmailId,
			Section:  ticket.Section,
			From:     ticket.TrainObj.From,
			To:       ticket.TrainObj.To,
			TrainNo:  ticket.TrainObj.Number,
			Price:    ticket.TrainObj.Fare,
			SeatNo:   strconv.Itoa(ticket.SeatNo),
		}, nil
	}

	return &TicketPurchaseResMsg{}, err
}

func (s *TicketServer) GetAllTickets(ctx context.Context, in *DummyMessage) (*TicketsMinListRes, error) {

	res := TicketsMinListRes{
		Tickets: []*TicketEntry{},
	}

	for _, v := range db.Tickets.Map {
		e := TicketEntry{TicketNo: v.Number, UserId: v.UserObj.Id, SeatNumber: strconv.Itoa(v.SeatNo), Section: v.Section}
		res.Tickets = append(res.Tickets, &e)
	}

	return &res, nil

}

func (s *TicketServer) GetReceipt(ctx context.Context, in *Tickets) (*TicketPurchaseResMsg, error) {

	ticket, exists := db.Tickets.Map[in.TicketNo]

	if exists {
		t := TicketPurchaseResMsg{FirstName: ticket.UserObj.FirstName,
			LastName: ticket.UserObj.LastName,
			TicketNo: ticket.Number,
			Email:    ticket.UserObj.EmailId,
			Section:  ticket.Section,
			From:     ticket.TrainObj.From,
			To:       ticket.TrainObj.To,
			TrainNo:  ticket.TrainObj.Number,
			Price:    ticket.TrainObj.Fare,
			SeatNo:   strconv.Itoa(ticket.SeatNo),
		}
		return &t, nil
	}
	return &TicketPurchaseResMsg{}, errors.New("ticket is invalid")
}
