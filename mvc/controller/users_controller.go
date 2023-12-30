package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/omkarrrJoshi/golang-microservices/mvc/services"
	"github.com/omkarrrJoshi/golang-microservices/mvc/utils"
)

func GetUser(resp http.ResponseWriter, req *http.Request) {
	userIdParam := req.URL.Query().Get("user_id")
	userId, err := strconv.ParseInt(userIdParam, 10, 64)
	if err != nil {
		apiErr := utils.ApplicationError{
			Msg:        "user id must be a number",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}
		jsonErr, _ := json.Marshal(apiErr)
		resp.WriteHeader(apiErr.StatusCode)
		resp.Write(jsonErr)
		//return bad request to user
		return
	}
	user, apiErr := services.UserService.GetUser(userId)
	if apiErr != nil {
		jsonErr, _ := json.Marshal(apiErr)
		resp.WriteHeader(apiErr.StatusCode)
		resp.Write(jsonErr)
		return
	}

	// return user to client
	jsonVal, _ := json.Marshal(user)
	resp.Write(jsonVal)
}
