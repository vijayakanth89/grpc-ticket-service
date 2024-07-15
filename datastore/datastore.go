package datastore

import (
	"errors"
)

// return seat No, section code and error if applicable

func seatsAvailable(t *Train, user *User) (int, string, error) {

	seatNo, error := t.SectionA.GetSeatNo()

	if error == nil {
		// marking the seat as allocated.
		t.SectionA.M[seatNo] = user.Id
		return seatNo, "A", nil
	}

	seatNo, error = t.SectionB.GetSeatNo()

	if error == nil {
		// marking the seat as allocated.
		t.SectionB.M[seatNo] = user.Id
		return seatNo, "B", nil
	}

	return 0, "", errors.New(ERROR_NO_SEATS)

}

func (section *Allocations) GetSeatNo() (int, error) {
	for key, value := range section.M {
		if value == "0" {
			return key, nil
		}
	}
	return 0, errors.New(ERROR_NO_SEATS)
}

func PurchaseTicket(firstName string, lastName string, email string, trainNo string) (Ticket, error) {

	user, createUserError := Users.CreateUser(firstName, lastName, email)
	if createUserError != nil {
		return Ticket{}, createUserError
	}

	train, error := findTrain(trainNo)

	if error == nil {
		seatNo, section, error := seatsAvailable(train, user)

		if error == nil {

			newTicketNo := Tickets.getNextId()

			newTicketObj := Ticket{SeatNo: seatNo, Section: section, UserObj: user, Number: newTicketNo, TrainObj: train}

			Tickets.Map[newTicketNo] = &newTicketObj

			return newTicketObj, nil

		} else {
			return Ticket{}, error
		}
	}

	return Ticket{}, errors.New("Could not create Ticket")

}

func initializeAllocations(maxSeats int) map[int]string {
	allocations := make(map[int]string, MAX_SECTION_CAPACITY)

	for seat := 1; seat <= MAX_SECTION_CAPACITY; seat++ {
		allocations[seat] = "0" // Assuming "0" as the default value
	}
	return allocations
}

func findTrain(trainNo string) (*Train, error) {
	for _, train := range AvailableTrains {
		if train.Number == trainNo {
			return &train, nil
		}
	}
	return nil, errors.New("train not found")
}
