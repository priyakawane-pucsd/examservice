package ping

import (
	"context"
	"examservice/models/dto"
	"examservice/utils"

	"github.com/bappaapp/goutils/logger"
	"github.com/gin-gonic/gin"
)

type PingController struct {
	service Service
}

type Service interface {
	Ping(ctx context.Context) error
}

func NewPingController(ctx context.Context, service Service) *PingController {
	return &PingController{service: service}
}

func (pc *PingController) Register(router gin.IRouter) {
	pingRouter := router.Group("/examservice/ping")
	pingRouter.GET("/", pc.Ping)
}

// Ping handles the ping endpoint.
// @Summary Pings the server.
// @Description Pings the server and returns "Okay" if successful.
// @Tags Ping
// @Param db query bool false "Flag indicating whether to ping the database"
// @Produce json
// @Success 200 {object} dto.PingResponse
// @Error 500 utils.CustomError
// @Router /examservice/ping/ [get]
func (pc *PingController) Ping(ctx *gin.Context) {
	db := ctx.Query("db")
	if db == "true" {
		logger.Debug(ctx, "pinging db")
		err := pc.service.Ping(ctx)
		if err != nil {
			utils.WriteError(ctx, err)
			return
		}
	}
	utils.WriteResponse(ctx, dto.PingResponse{StatusCode: 200, Message: "Okay"})
}
