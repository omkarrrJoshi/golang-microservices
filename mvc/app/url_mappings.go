package app

import (
	"github.com/omkarrrJoshi/golang-microservices/mvc/controller"
)

func mapUrls() {
	router.GET("/users/:user_id", controller.GetUser)
}
