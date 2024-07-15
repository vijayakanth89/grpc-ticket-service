package datastore

import (
	"errors"
	"fmt"
	"strconv"
)

type Ticket struct {
	UserObj  *User
	Number   string
	TrainObj *Train
	SeatNo   int
	Section  string
}

type TicketCollection struct {
	Map    map[string]*Ticket
	LastId int
}

func (collection *TicketCollection) Exists(TicketNo string) (t *Ticket, err error) {
	ticket, exists := Tickets.Map[TicketNo]

	fmt.Printf("%#v\n", ticket)

	if !exists || ticket == nil {
		return &Ticket{}, errors.New(ERROR_INVALID_TICKET_NO)
	}
	return ticket, nil
}

func (ticket *Ticket) GetSeatNumber() string {
	return strconv.Itoa(ticket.SeatNo)
}

func (ticket *Ticket) SeatReallocate() (err error) {
	var allo1, allo2 *Allocations
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
		return errors.New(fmt.Sprintf("invalid section: %s", ticket.Section))
	}

	newSeatNo, newSection, err1 = allo1.FindNewSeat()

	if newSeatNo == 0 || err1 != nil {
		newSeatNo, newSection, err1 = allo2.FindNewSeat()
	} else {
		if newSeatNo == 0 || err1 != nil {
			return errors.New(ERROR_NO_SEATS)
		}
	}

	ticket.Section = newSection
	ticket.SeatNo = newSeatNo

	return nil
}

func (ticket *Ticket) Cancel() (err error) {
	if ticket.Section == "A" {
		err = ticket.TrainObj.SectionA.RemoveUser(ticket.SeatNo)
	} else if ticket.Section == "B" {
		err = ticket.TrainObj.SectionB.RemoveUser(ticket.SeatNo)
	} else {
		err = errors.New(fmt.Sprintf(ERROR_INVALID_SECTION_TEMPLATE, ticket.Section))
	}
	return err
}

// getNextId generates the next user ID
func (c *TicketCollection) getNextId() string {
	c.LastId++
	nextId := strconv.Itoa(c.LastId)
	return nextId
}

func (c *TicketCollection) ListTickets() error {
	for _, v := range c.Map {
		fmt.Printf("%#v\n", *v)
	}

	return nil
}

var Tickets = TicketCollection{LastId: 0, Map: make(map[string]*Ticket)}
