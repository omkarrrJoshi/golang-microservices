package services

import (
	"github.com/omkarrrJoshi/golang-microservices/mvc/model"
	"github.com/omkarrrJoshi/golang-microservices/mvc/utils"
)

func GetUser(userId int64) (*model.User, *utils.ApplicationError) {
	return model.GetUser(userId)
}
