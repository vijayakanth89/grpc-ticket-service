package datastore

import "errors"

type Train struct {
	Number   string
	From     string
	To       string
	Fare     string
	SectionA Allocations
	SectionB Allocations
}

type Allocations struct {
	// seatNo to User.Id tomap here.
	M       map[int]string
	Section string
}

// returns new seat number if found
func (alloc *Allocations) FindNewSeat() (int, string, error) {
	for seatNo, userId := range alloc.M {
		if userId == "0" {
			return seatNo, alloc.Section, nil
		}
	}
	return 0, "", errors.New(ERROR_NO_SEATS)
}
