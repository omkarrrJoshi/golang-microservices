package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/omkarrrJoshi/golang-microservices/mvc/services"
	"github.com/omkarrrJoshi/golang-microservices/mvc/utils"
)

func GetUser(ctx *gin.Context) {
	userIdParam := ctx.Param("user_id")
	userId, err := strconv.ParseInt(userIdParam, 10, 64)
	if err != nil {
		apiErr := &utils.ApplicationError{
			Msg:        "user id must be a number",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}
		utils.RespondError(ctx, apiErr)
		return
	}

	user, apiErr := services.UserService.GetUser(userId)
	if apiErr != nil {
		utils.RespondError(ctx, apiErr)
		return
	}

	// return user to client
	utils.Respond(ctx, http.StatusOK, user)
}
