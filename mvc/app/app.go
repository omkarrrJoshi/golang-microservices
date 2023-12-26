package app

import (
	"net/http"

	"github.com/omkarrrJoshi/golang-microservices/mvc/controller"
)

func StartApp() {
	http.HandleFunc("/users", controller.GetUser)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
