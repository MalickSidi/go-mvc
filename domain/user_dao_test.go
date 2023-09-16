package domain

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestGetUserNoUserFound(t *testing.T) {
	user, err := GetUser(0)

	assert.Nil(t, user, "User with id: 0 Should return nil")
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.StatusCode)
	assert.EqualValues(t, "user with id: 0 was not found", err.Message)
	assert.EqualValues(t, "not_found", err.Code)
}

func TestGetUserNoError(t *testing.T) {
	id := uint64(1)
	user, err := GetUser(id)

	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.EqualValues(t, id, user.Id)
}

// if user != nil {
// t.Error("There is no User with id: 0")
// }
// if err == nil {
// t.Error("User with id: 0 Should return an Error")
// }
// if err.StatusCode != http.StatusNotFound {
// t.Error("Status code Should be: ( 404 ) When user is not found")
// }
