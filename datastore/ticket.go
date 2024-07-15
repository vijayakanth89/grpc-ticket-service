package datastore

import (
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
