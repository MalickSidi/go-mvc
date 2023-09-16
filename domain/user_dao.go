package domain

import (
	"fmt"
	"net/http"

	"github.com/MalickSidi/mvc01/utils"
)

var (
	users = map[uint64]*User{
		1: {Id: 1, FirstName: "Lema", LastName: "Thomas", Email: "lema@email.com"},
		2: {Id: 2, FirstName: "Thamer", LastName: "TJ", Email: "thamer@email.com"},
		3: {Id: 3, FirstName: "Sara", LastName: "Lafy", Email: "sara@email.com"},
	}
)

func GetUser(userId uint64) (*User, *utils.ApplicationError) {
	if user := users[userId]; user != nil {
		return user, nil
	}

	return nil, &utils.ApplicationError{
		Message:    fmt.Sprintf("user with id: %v was not found", userId),
		StatusCode: http.StatusNotFound,
		Code:       "not_found",
	}
}
