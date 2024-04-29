package utils

import (
	"net/http"
	"strconv"

	"github.com/bappaapp/goutils/logger"
	"github.com/gin-gonic/gin"
)

func GetUserIdFromContext(ctx *gin.Context) (int64, error) {
	userId := ctx.GetHeader("X-USER-ID")
	num, err := strconv.ParseInt(userId, 10, 64)
	if err != nil || num == 0 {
		logger.Error(ctx, "Failed to parse header X-USER-ID: %s", err.Error())
		return 0, NewCustomError(http.StatusUnauthorized, "invalid userId")
	}
	return num, nil
}
