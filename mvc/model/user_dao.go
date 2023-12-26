package model

import (
	"errors"
	"fmt"
)

var (
	users = map[int64]*User{
		123: &User{Id: 1, FirstName: "omkar", LastName: "joshi", Email: "o.j@gmail.com"},
	}
)

func GetUser(userId int64) (*User, error) {
	user, ok := users[userId]
	if ok {
		return user, nil
	}
	return nil, errors.New(fmt.Sprintf("user with userId %v not found", userId))
}
