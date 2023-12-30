package services

import (
	"net/http"
	"testing"

	"github.com/omkarrrJoshi/golang-microservices/mvc/model"
	"github.com/omkarrrJoshi/golang-microservices/mvc/utils"
	"github.com/stretchr/testify/assert"
)

var (
	getUserFunction func(userId int64) (*model.User, *utils.ApplicationError)
)

func init() {
	model.UserDao = &userDaoMock{}
}

type userDaoMock struct{}

func (u *userDaoMock) GetUser(userId int64) (*model.User, *utils.ApplicationError) {
	return getUserFunction(userId)
}

func TestGetUserNotFoundInDatabase(t *testing.T) {
	getUserFunction = func(userId int64) (*model.User, *utils.ApplicationError) {
		return nil, &utils.ApplicationError{
			Msg:        "user not found",
			StatusCode: http.StatusNotFound,
		}
	}
	user, err := UserService.GetUser(0)
	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, err.StatusCode, http.StatusNotFound)
}

func TestGetUserNoError(t *testing.T) {
	getUserFunction = func(userId int64) (*model.User, *utils.ApplicationError) {
		return &model.User{Id: 123, FirstName: "omkar", LastName: "joshi"}, nil
	}
	user, err := UserService.GetUser(123)
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.EqualValues(t, user.FirstName, "omkar")
	assert.EqualValues(t, user.LastName, "joshi")
}
