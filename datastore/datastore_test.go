package datastore

import (
	"errors"
	"fmt"
	"strconv"
	"testing"
)

// TestAdd tests the Add function.
func TestValidateStrings_EMPTY(t *testing.T) {
	result := CheckEmpty("")

	expected := errors.New(ERROR_EMPTY_STRING)

	if result.Error() != expected.Error() {
		t.Errorf("CheckEmpty(\"\") ='%v'; want '%v'", result, expected)
	}
}

func TestValidateStrings_NOT_EMPTY(t *testing.T) {
	result := CheckEmpty("TEST")

	var expected error = nil

	if result != expected {
		t.Errorf("CheckEmpty(\"\") ='%v'; want '%v'", result, expected)
	}
}

func TestIsZero_ZERO(t *testing.T) {
	result := IsZero("0")
	expected := errors.New(ERROR_INVALID_USER_ID)

	if result.Error() != expected.Error() {
		t.Errorf("IsZero(\"0\") ='%v'; want '%v'", result, expected)
	}
}

func TestIsZero_NONZERO(t *testing.T) {
	result := IsZero("1")
	var expected error = nil

	if result != expected {
		t.Errorf("IsZero(\"0\") ='%v'; want '%v'", result, expected)
	}
}

func TestIsValidEmail_EMPTY(t *testing.T) {
	email := ""

	errorMessage := fmt.Sprintf(ERROR_INVALID_EMAIL_TEMPLATE, email)
	result := IsValidEmail(email)
	expected := errors.New(errorMessage)

	if result.Error() != expected.Error() {
		t.Errorf("IsValidEmail(\"\") ='%v'; want '%v'", result, expected)
	}
}

func TestIsValidEmail_InvalidNonEmpty(t *testing.T) {

	tests := []struct {
		name     string
		email    string
		expected string
	}{
		{"without@", "abcgmail.com", fmt.Sprintf(ERROR_INVALID_EMAIL_TEMPLATE, "abcgmail.com")},
		{"without.and@", "abcgmailcom", fmt.Sprintf(ERROR_INVALID_EMAIL_TEMPLATE, "abcgmailcom")},
		{"without first", "@abcgmailcom", fmt.Sprintf(ERROR_INVALID_EMAIL_TEMPLATE, "@abcgmailcom")},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsValidEmail(tt.email)

			if result.Error() != tt.expected {
				t.Errorf("IsValidEmail(\"\") ='%v'; want '%v'", result, tt.expected)
			}

		})
	}

}

func TestIsValidEmail_ValidEmail(t *testing.T) {

	result := IsValidEmail("vijayakanthblog@gmail.com")

	if result != nil {
		t.Errorf("IsValidEmail(\"\") ='%v'; want '%v'", result, nil)
	}
}

func TestPurchase_emptyFirstName(t *testing.T) {
	_, err := PurchaseTicket("", "madhavan", "t@gmail.com", "T001")

	if err.Error() != ERROR_EMPTY_STRING {
		t.Errorf("err.Error(): %s, EMPTY_STRING: %s\n", err.Error(), ERROR_EMPTY_STRING)
	}
}

func TestPurchase_invalidEmail(t *testing.T) {
	invalidEmail := "gmail.com"
	_, err := PurchaseTicket("vijayakanth", "madhavan", invalidEmail, "T001")

	if err.Error() != fmt.Sprintf(ERROR_INVALID_EMAIL_TEMPLATE, invalidEmail) {
		t.Errorf("err.Error(): %s, EMAIL_ERROR_TEMPLATE: %s\n", err.Error(), ERROR_INVALID_EMAIL_TEMPLATE)
	}
}

func TestPurchase_ValidUserDetails(t *testing.T) {

	// will create a fresh user if it does not exists or use the existing user record.

	ticket, err := PurchaseTicket("vijayakanth", "madhavan", "vijayakanthblog@gmail.com", "T001")

	if err != nil && len(ticket.Number) > 0 {
		t.Errorf("PurchaseTicket accepts valid ")
	}
}

func TestAllocationStore_LoadValues(t *testing.T) {
	max := 20
	all := initializeAllocations(max)

	if len(all) == max {
		for seat, value := range all {
			if value != "0" {
				t.Errorf("seatNo: %v, value: %v", seat, value)
			}
		}
	}
}

func TestFindTrain_InvalidTrainNo(t *testing.T) {
	trainNo := "679"
	_, err := findTrain(trainNo)

	if err == nil {
		t.Errorf("trainNo: %s, match with train entries and see", trainNo)
	}

}

func TestFindTrain_validTrainNo(t *testing.T) {
	trainNo := "T004"
	train, err := findTrain(trainNo)

	if err != nil || train == nil {
		t.Errorf("trainNo: %s, match with train entries and see", trainNo)
	}

}

func TestAllocationStore_FindSeat(t *testing.T) {
	all := Allocations{M: initializeAllocations(20), Section: "A"}

	seat, err := all.GetSeatNo()

	if seat == 0 || err != nil {
		t.Errorf("seat: %d, error: %v", seat, err)
	}

}

func TestFindNewSeat(t *testing.T) {
	all := Allocations{M: initializeAllocations(20), Section: "A"}

	user, _ := Users.CreateUser("vijayakanth", "madhavan", "vijayakanthblog@gmail.com")

	for key, _ := range all.M {
		if key != 15 {
			all.M[key] = user.Id
		}
	}

	seat, _, _ := all.FindNewSeat()

	if seat != 15 {
		t.Errorf("seat: %d", seat)
	}
}

func TestFindNewSeat_SectionFull(t *testing.T) {
	all := Allocations{M: initializeAllocations(20), Section: "A"}

	user, _ := Users.CreateUser("vijayakanth", "madhavan", "vijayakanthblog@gmail.com")

	for key, _ := range all.M {
		// if key != 15 {
		all.M[key] = user.Id
		// }
	}

	_, _, err := all.FindNewSeat()

	if err.Error() != ERROR_NO_SEATS {
		t.Errorf("allocations: %v", all)
	}
}

func TestPurchase_SubsequenTicketNo(t *testing.T) {
	t1, _ := PurchaseTicket("vijayakanth", "madhavan", "vijayakanthblog@gmail.com", "T001")

	t2, _ := PurchaseTicket("vijayakanth", "madhavan", "vijayakanthblog@gmail.com", "T001")

	n1, err1 := strconv.Atoi(t1.Number)
	n2, err2 := strconv.Atoi(t2.Number)

	if err1 != nil {
		t.Errorf("err1: %v ", err1)
	}

	if err2 != nil {
		t.Errorf("err1: %v ", err2)
	}

	if n1 == n2 {
		t.Errorf("t1.Number: %s, t2.Number: %s\n", t1.Number, t2.Number)
	}
}

func TestPurchase_MoreThanMax(test *testing.T) {

	for i := 0; i < MAX_SECTION_CAPACITY*2+1; i++ {
		firstName, lastName, email := randomDetails()
		t, err := PurchaseTicket(firstName, lastName, email, "T001")

		if err != nil {
			if err.Error() == ERROR_NO_SEATS {
				return
			} else {
				test.Errorf("error whlie generating ticket: %d, %v", i, err)
				return
			}
		} else {
			fmt.Printf("ticketNo: %s, userId: %s\n", t.Number, t.UserObj.Id)
		}
	}

	test.Errorf("should have caused the MAX_SECTION_CAPACITY error")
}

func TestPurchase_MoreThanMaxAndFindNewSeat(test *testing.T) {

	var arr []*Ticket

	for i := 0; i < MAX_SECTION_CAPACITY*2; i++ {
		firstName, lastName, email := randomDetails()
		ticket, oerr := PurchaseTicket(firstName, lastName, email, "T001")
		if oerr != nil {
			continue
		}

		arr = append(arr, &ticket)
		err := ticket.Cancel()
		if err != nil {
			test.Errorf("err: %v", err)
		}
	}

}

func TestInvalidTicket(test *testing.T) {
	user, _ := Users.CreateUser("vijayakanth", "madhavan", "vijayakanthblog@gmail.com")
	ticket := &Ticket{UserObj: user, Number: "", SeatNo: 0, TrainObj: &AvailableTrains[1], Section: ""}

	err := ticket.Cancel()

	if err == nil {
		test.Errorf("This should 've caused error, since the ticket details are invalid %v", ticket)
	}
}

func TestRemoveUser(t *testing.T) {
	all := Allocations{M: initializeAllocations(20), Section: "A"}

	user, _ := Users.CreateUser("f1", "l1", "vijayakanthblog@gmail.com")

	all.M[15] = user.Id

	err := all.RemoveUser(15)

	if err != nil {
		t.Errorf("err: %v", err)
	}
}

func TestCreateUserDuplicate(t *testing.T) {

	user1, _ := Users.CreateUser("f1", "l1", "vijayakanthblog@gmail.com")
	user2, _ := Users.CreateUser("f1", "l1", "vijayakanthblog@gmail.com")

	if user1.Id != user2.Id {
		t.Errorf("user1: %v, user2: %v", user1, user2)
	}

}

func TestCreateUserDifferent(t *testing.T) {

	user1, _ := Users.CreateUser("f1", "l1", "vijayakanthblog@gmail.com")
	user2, _ := Users.CreateUser("f1", "l1", "rabhavan@gmail.com")

	if user1.Id == user2.Id {
		t.Errorf("user1: %v, user2: %v", user1, user2)
	}

}

func TestListTickets(t *testing.T) {
	err := Tickets.ListTickets()

	if err != nil {
		t.Errorf("This shoudl not geneate an error")
	}
}

func TestTicketExists(t *testing.T) {
	ticket, err := Tickets.Exists("25")

	fmt.Printf("t:%v, err: %v", ticket, err)

	if err != nil {
		t.Errorf("t:%v, err: %v", ticket, err)
	}
}

func TestTicketExistsInvalidTicket(t *testing.T) {
	ticket, err := Tickets.Exists("XKDK25")

	fmt.Printf("t:%v, err: %v", ticket, err)

	if err == nil {
		t.Errorf("t:%v, err: %v", ticket, err)
	}
}

func TestTestPurchaseDummy(t *testing.T) {
	TestPurchase(2)
}

func TestNextIntTicketCollections(test *testing.T) {
	ticketId := Tickets.getNextId()
	ticketId2 := Tickets.getNextId()

	if ticketId == ticketId2 {
		test.Errorf("ticket1: %v, ticket2: %v", ticketId, ticketId2)
	}
}

func TestPurchaseAndReallocate(test *testing.T) {
	// chanding the train here to creat fresh ticket

	train := "T001"

	for _, value := range Tickets.Map {
		if value.TrainObj.Number == train {
			value.Cancel()
		}
	}

	TestPurchase(MAX_SECTION_CAPACITY)

	t, e := PurchaseTicket("sachin", "tendulkar", "sachin@yahoo.co.in", train)

	oldSeat := t.GetSeatNumber()

	if e == nil {
		err := t.SeatReallocate()

		if err != nil && err.Error() != ERROR_NO_SEATS {
			test.Errorf("e.error():%s", err.Error())
		} else if err != nil && err.Error() == ERROR_NO_SEATS {
			fmt.Println("coming here..")
			if oldSeat != t.GetSeatNumber() {
				test.Errorf("realloc failed..e.error()")
			}
		} else {

			if oldSeat == t.GetSeatNumber() {
				test.Errorf("realloc failed..e.error()")
			}
		}
	}
}

func TestPurchaseAtMaxCapacity(test *testing.T) {
	// chanding the train here to creat fresh ticket

	train := "T001"

	for _, value := range Tickets.Map {
		if value.TrainObj.Number == train {
			value.Cancel()
		}
	}

	TestPurchase(MAX_SECTION_CAPACITY*2 - 1)

	t, e := PurchaseTicket("sachin", "tendulkar", "sachin@yahoo.co.in", train)

	oldSeat := t.GetSeatNumber()

	if e == nil {
		err := t.SeatReallocate()

		if err != nil && err.Error() != ERROR_NO_SEATS {
			test.Errorf("e.error():%s", err.Error())
		} else if err != nil && err.Error() == ERROR_NO_SEATS {
			fmt.Println("coming here..")
			if oldSeat != t.GetSeatNumber() {
				test.Errorf("realloc failed..e.error()")
			}
		} else {

			if oldSeat == t.GetSeatNumber() {
				test.Errorf("realloc failed..e.error()")
			}
		}
	}
}

func TestPurchaseAtMaxCapacityAndTwo(test *testing.T) {
	// chanding the train here to creat fresh ticket

	train := "T001"

	for _, value := range Tickets.Map {
		if value.TrainObj.Number == train {
			value.Cancel()
		}
	}

	TestPurchase(MAX_SECTION_CAPACITY*2 - 2)

	t, e := PurchaseTicket("sachin", "tendulkar", "sachin@yahoo.co.in", train)

	oldSeat := t.GetSeatNumber()

	if e == nil {
		err := t.SeatReallocate()

		if err != nil && err.Error() != ERROR_NO_SEATS {
			test.Errorf("e.error():%s", err.Error())
		} else if err != nil && err.Error() == ERROR_NO_SEATS {
			if oldSeat != t.GetSeatNumber() {
				test.Errorf("realloc failed..e.error()")
			}
		} else {

			if oldSeat == t.GetSeatNumber() {
				test.Errorf("realloc failed..e.error()")
			}
		}
	}
}
