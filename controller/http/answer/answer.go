package answer

import (
	"context"
	"examservice/models/dto"
	"examservice/utils"

	"github.com/bappaapp/goutils/logger"
	"github.com/gin-gonic/gin"
)

type AnswerController struct {
	service Service
}

type Service interface {
	CreateOrUpdateAnswer(ctx context.Context, req *dto.AnswerRequest, answerId string) (string, error)
}

func NewAnswerController(ctx context.Context, service Service) *AnswerController {
	return &AnswerController{service: service}
}

func (ac *AnswerController) Register(router gin.IRouter) {
	AnswerRouter := router.Group("/examservice/answers")
	AnswerRouter.PUT("/submit/:id", ac.CreateOrUpdateAnswer)
}

// @Summary Create or update answer
// @Description Creates or updates an answer based on the provided request body.
// @Tags Answers
// @Accept json
// @Produce json
// @Param X-USER-ID header string true "User ID"
// @Param id path string false "ID of the answer to update"
// @Param body body dto.AnswerRequest true "Answer request body"
// @Success 200 {object} dto.AnswerResponse "Successful operation"
// @Failure 400 {object} utils.CustomError "Invalid request body"
// @Failure 500 {object} utils.CustomError "Internal server error"
// @Router /examservice/answers/submit/{id} [put]
func (ac *AnswerController) CreateOrUpdateAnswer(ctx *gin.Context) {
	var req dto.AnswerRequest
	err := ctx.BindJSON(&req)
	if err != nil {
		logger.Error(ctx, "Failed to parse request body: %s", err.Error())
		utils.WriteError(ctx, utils.NewBadRequestError("Invalid request body"))
		return
	}

	// Validate the request
	if err := req.Validate(); err != nil {
		logger.Error(ctx, "Validation error: %s", err.Error())
		utils.WriteError(ctx, utils.NewBadRequestError(err.Error()))
		return
	}

	answerId := ctx.Param("id")
	// Access X-USER-ID header
	req.UserID, err = utils.GetUserIdFromContext(ctx)
	if err != nil {
		utils.WriteError(ctx, utils.NewBadRequestError("Invalid userId"))
		return
	}

	res, err := ac.service.CreateOrUpdateAnswer(ctx, &req, answerId)
	if err != nil {
		utils.WriteError(ctx, err)
		return
	}
	utils.WriteResponse(ctx, res)
}
