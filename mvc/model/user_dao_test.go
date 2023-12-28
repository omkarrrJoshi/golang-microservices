package model

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUserNotFound(t *testing.T) {
	user, err := GetUser(0)

	assert.Nil(t, user, "We are not expecting user with id 0")
	assert.NotNil(t, err, "We are expecting error when user id is 0")
	assert.EqualValues(t, err.StatusCode, http.StatusNotFound)
	assert.EqualValues(t, err.Code, "not_found")
	assert.EqualValues(t, err.Msg, "user with userId 0 not found")
}

func TestGetUserNoError(t *testing.T) {
	user, err := GetUser(123)

	assert.Nil(t, err, "We are not expecting an error")
	assert.NotNil(t, user, "We are expecting user with id 123")
	assert.EqualValues(t, user.Id, 123)
	assert.EqualValues(t, user.FirstName, "omkar")
	assert.EqualValues(t, user.LastName, "joshi")
	assert.EqualValues(t, user.Email, "o.j@gmail.com")

}
