package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Message *string `json:"message,omitempty"`
	Error   *string `json:"error,omitempty"`
	Code    int     `json:"code"`
	Data    any     `json:"data,omitempty"`
}

func WriteError(ctx *gin.Context, err error) {
	if cErr, ok := err.(*CustomError); ok {
		ctx.JSON(cErr.StatusCode, Response{Error: &cErr.Message, Code: cErr.StatusCode})
		return
	}
	errstr := err.Error()
	ctx.JSON(http.StatusInternalServerError, Response{Code: http.StatusInternalServerError, Error: &errstr})
}

func WriteResponse(ctx *gin.Context, res any) {
	if msg, ok := res.(string); ok {
		ctx.JSON(http.StatusOK, Response{Message: &msg, Code: http.StatusOK})
		return
	}
	if data, ok := res.(interface{}); ok {
		// Response is a data object
		ctx.JSON(http.StatusOK, Response{Data: data, Code: http.StatusOK})
		return
	}

	// Default response for unknown types
	ctx.JSON(http.StatusOK, Response{Data: res, Code: http.StatusOK})
	//ctx.JSON(http.StatusOK, Response{Data: http.StatusOK, Code: http.StatusOK})
}
