package datastore

import (
	"fmt"
	"math/rand"
	"time"
)

var AvailableTrains = []Train{
	{Number: "T001", From: "Chennai", To: "Bangalore", Fare: "$20", SectionA: Allocations{M: initializeAllocations(MAX_SECTION_CAPACITY), Section: "A"}, SectionB: Allocations{M: initializeAllocations(MAX_SECTION_CAPACITY), Section: "B"}},
	{Number: "T002", From: "Mumbai", To: "Delhi", Fare: "$20", SectionA: Allocations{M: initializeAllocations(MAX_SECTION_CAPACITY), Section: "A"}, SectionB: Allocations{M: initializeAllocations(MAX_SECTION_CAPACITY), Section: "B"}},
	{Number: "T003", From: "Hyderabad", To: "Pune", Fare: "$20", SectionA: Allocations{M: initializeAllocations(MAX_SECTION_CAPACITY), Section: "A"}, SectionB: Allocations{M: initializeAllocations(MAX_SECTION_CAPACITY), Section: "B"}},
	{Number: "T004", From: "Kolkata", To: "Guwahati", Fare: "$20", SectionA: Allocations{M: initializeAllocations(MAX_SECTION_CAPACITY), Section: "A"}, SectionB: Allocations{M: initializeAllocations(MAX_SECTION_CAPACITY), Section: "B"}},
	{Number: "T005", From: "Ahmedabad", To: "Surat", Fare: "$20", SectionA: Allocations{M: initializeAllocations(MAX_SECTION_CAPACITY), Section: "A"}, SectionB: Allocations{M: initializeAllocations(MAX_SECTION_CAPACITY), Section: "B"}},
	{Number: "T006", From: "Chennai", To: "Coimbatore", Fare: "$20", SectionA: Allocations{M: initializeAllocations(MAX_SECTION_CAPACITY), Section: "A"}, SectionB: Allocations{M: initializeAllocations(MAX_SECTION_CAPACITY), Section: "B"}},
	{Number: "T007", From: "Bangalore", To: "Mysore", Fare: "$20", SectionA: Allocations{M: initializeAllocations(MAX_SECTION_CAPACITY), Section: "A"}, SectionB: Allocations{M: initializeAllocations(MAX_SECTION_CAPACITY), Section: "B"}},
	{Number: "T008", From: "Jaipur", To: "Udaipur", Fare: "$20", SectionA: Allocations{M: initializeAllocations(MAX_SECTION_CAPACITY), Section: "A"}, SectionB: Allocations{M: initializeAllocations(MAX_SECTION_CAPACITY), Section: "B"}},
	{Number: "T009", From: "Lucknow", To: "Varanasi", Fare: "$20", SectionA: Allocations{M: initializeAllocations(MAX_SECTION_CAPACITY), Section: "A"}, SectionB: Allocations{M: initializeAllocations(MAX_SECTION_CAPACITY), Section: "B"}},
	{Number: "T010", From: "Patna", To: "Ranchi", Fare: "$20", SectionA: Allocations{M: initializeAllocations(MAX_SECTION_CAPACITY), Section: "A"}, SectionB: Allocations{M: initializeAllocations(MAX_SECTION_CAPACITY), Section: "B"}},
}

const (
	ERROR_NO_SEATS       = "no seats available"
	INVALID_USER_ID      = "invalid user id"
	EMPTY_STRING         = "parameter cannot be empty"
	EMAIL_ERROR_TEMPLATE = "email:'%s', is invalid"
)

const (
	MAX_SECTION_CAPACITY = 20
)

func randomDetails() (string, string, string) {
	firstNames := []string{"Alice", "Bob", "Charlie", "David", "Emma", "Frank", "Grace"}
	lastNames := []string{"Smith", "Johnson", "Williams", "Jones", "Brown", "Davis"}

	rand.Seed(time.Now().UnixNano())
	firstName := firstNames[rand.Intn(len(firstNames))]
	lastName := lastNames[rand.Intn(len(lastNames))]
	email := fmt.Sprintf("%s.%s@example.com", firstName, lastName)

	return firstName, lastName, email
}

func TestPurchase() {
	for i := 0; i < 10; i++ {
		firstName, lastName, email := randomDetails()
		t, err := PurchaseTicket(firstName, lastName, email, "T001")

		if err == nil {
			fmt.Printf("ticketNo: %s, userId: %s\n", t.Number, t.UserObj.Id)
		} else {
			fmt.Println(err)
		}
	}

	for i := 0; i < 10; i++ {
		firstName, lastName, email := randomDetails()
		t, err := PurchaseTicket(firstName, lastName, email, "T002")

		if err == nil {
			fmt.Printf("ticketNo: %s, userId: %s\n", t.Number, t.UserObj.Id)
		} else {
			fmt.Println(err)
		}
	}
}
