package datastore

import (
	"errors"
	"fmt"
	"strconv"
)

type User struct {
	FirstName string
	LastName  string
	EmailId   string
	Id        string
}

type UsersCollection struct {
	Map    map[string]*User
	LastId int
}

// getNextId generates the next user ID
func (c *UsersCollection) getNextId() string {
	c.LastId++
	nextId := strconv.Itoa(c.LastId)
	return nextId
}

func (users *UsersCollection) GetUser(userId string) (*User, error) {

	error := CheckEmpty(userId)

	if error != nil {
		return &User{}, error
	}

	error = IsZero(userId)

	if error != nil {
		return &User{}, error
	}

	for _, value := range Users.Map {
		if value.Id == userId {

			return value, nil
		}
	}
	return &User{}, errors.New(fmt.Sprintf("user with id: %s is not found", userId))
}

func (users *UsersCollection) CreateUser(firstName string, lastName string, emailId string) (*User, error) {

	err := CheckEmpty(firstName, lastName, emailId)

	if err != nil {
		return &User{}, err
	}

	err = IsValidEmail(emailId)

	if err != nil {
		return &User{}, err
	}

	user, exists := users.Map[emailId]

	if exists {
		// already registered.
		return user, nil
	} else {
		nextId := users.getNextId()

		uNext := User{
			FirstName: firstName,
			LastName:  lastName,
			EmailId:   emailId,
			Id:        nextId,
		}

		users.Map[emailId] = &uNext
		return &uNext, nil
	}
}

var Users = UsersCollection{LastId: 0, Map: make(map[string]*User)}
