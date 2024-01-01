package utils

import "github.com/gin-gonic/gin"

func Respond(ctx *gin.Context, statusCode int, body interface{}) {
	if ctx.GetHeader("Accept") == "application/xml" {
		ctx.XML(statusCode, body)
		return
	}

	ctx.JSON(statusCode, body)
}

func RespondError(ctx *gin.Context, err *ApplicationError) {
	if ctx.GetHeader("Accept") == "application/xml" {
		ctx.XML(err.StatusCode, err)
		return
	}

	ctx.JSON(err.StatusCode, err)
}
