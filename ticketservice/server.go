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

	ticket, exists := db.Tickets.Map[in.TicketNo]

	if !exists {
		return &TicketReallocResMsg{}, errors.New(fmt.Sprintf("ticket %s not found", in.TicketNo))
	}

	var allo1, allo2 *db.Allocations
	var newSeatNo int
	var err1 error
	var newSection string

	if ticket.Section == "A" {
		allo1 = &ticket.TrainObj.SectionA
		allo2 = &ticket.TrainObj.SectionB
	} else if ticket.Section == "B" {
		allo1 = &ticket.TrainObj.SectionB
		allo2 = &ticket.TrainObj.SectionA
	} else {
		return &TicketReallocResMsg{}, errors.New(fmt.Sprintf("invalid section: %s", ticket.Section))
	}

	newSeatNo, newSection, err1 = allo1.FindNewSeat()

	if newSeatNo == 0 || err1 != nil {
		newSeatNo, newSection, err1 = allo2.FindNewSeat()
	} else {
		if newSeatNo == 0 || err1 != nil {
			return &TicketReallocResMsg{}, errors.New(db.ERROR_NO_SEATS)
		}
	}

	res := TicketReallocResMsg{NewSeatNo: strconv.Itoa(newSeatNo), OldSeatNo: strconv.Itoa(ticket.SeatNo)}

	ticket.Section = newSection
	ticket.SeatNo = newSeatNo

	return &res, nil
}

func (s *TicketServer) CancelTicket(ctx context.Context, in *CancelTicketRequest) (*CancelTicketResponse, error) {
	ticket, exists := db.Tickets.Map[in.TicketNo]

	fmt.Printf("%#v\n", ticket)

	if !exists || ticket == nil {
		return &CancelTicketResponse{Status: "F"}, errors.New("Invalid Ticket No")
	}

	var alloc *db.Allocations

	if ticket.Section == "A" {
		fmt.Println("comng in section A")
		alloc = &ticket.TrainObj.SectionA

		for k, v := range alloc.M {
			fmt.Printf("k:%s,v:%s\n", k, v)
		}
	} else if ticket.Section == "B" {
		alloc = &ticket.TrainObj.SectionB
	} else {
		return &CancelTicketResponse{Status: "F"}, errors.New(fmt.Sprintf(("Invalid section: %s"), ticket.Section))
	}

	alloc.M[ticket.SeatNo] = "0"
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
