package services

import (
	"github.com/omkarrrJoshi/golang-microservices/mvc/model"
	"github.com/omkarrrJoshi/golang-microservices/mvc/utils"
)

type userService struct{}

var UserService userService

func (u *userService) GetUser(userId int64) (*model.User, *utils.ApplicationError) {
	return model.UserDao.GetUser(userId)
}
