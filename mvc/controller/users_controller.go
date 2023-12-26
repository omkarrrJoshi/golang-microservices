package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/omkarrrJoshi/golang-microservices/mvc/services"
)

func GetUser(resp http.ResponseWriter, req *http.Request) {
	userIdParam := req.URL.Query().Get("user_id")
	userId, err := strconv.ParseInt(userIdParam, 10, 64)
	if err != nil {
		resp.WriteHeader(http.StatusBadRequest)
		resp.Write([]byte("user id must be a number"))
		//return bad request to user
		return
	}
	user, err := services.GetUser(userId)
	if err != nil {
		resp.WriteHeader(http.StatusNotFound)
		resp.Write([]byte(err.Error()))
		// handle the error and return to the client
		return
	}

	// return user to client
	jsonVal, _ := json.Marshal(user)
	resp.Write(jsonVal)
}
