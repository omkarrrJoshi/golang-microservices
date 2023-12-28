package model

import (
	"fmt"
	"net/http"

	"github.com/omkarrrJoshi/golang-microservices/mvc/utils"
)

var (
	users = map[int64]*User{
		123: {Id: 123, FirstName: "omkar", LastName: "joshi", Email: "o.j@gmail.com"},
	}
)

func GetUser(userId int64) (*User, *utils.ApplicationError) {
	user, ok := users[userId]
	if ok {
		return user, nil
	}
	err := &utils.ApplicationError{
		Msg:        fmt.Sprintf("user with userId %v not found", userId),
		StatusCode: http.StatusNotFound,
		Code:       "not_found",
	}
	return nil, err
}
