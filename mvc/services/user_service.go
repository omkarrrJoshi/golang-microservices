package services

import "github.com/omkarrrJoshi/golang-microservices/mvc/model"

func GetUser(userId int64) (*model.User, error) {
	return model.GetUser(userId)
}
