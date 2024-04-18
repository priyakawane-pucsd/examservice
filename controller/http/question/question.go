package question

import (
	"context"
	"examservice/models/dto"
	"examservice/utils"

	"github.com/bappaapp/goutils/logger"
	"github.com/gin-gonic/gin"
)

type QuestionController struct {
	service Service
}

type Service interface {
	CreateQuestions(ctx context.Context, req *dto.QuestionRequest) (*dto.QuestionResponse, error)
	GetQuestionsList(ctx context.Context) (*dto.ListQuestionResponse, error)
}

func NewQuestionController(ctx context.Context, service Service) *QuestionController {
	return &QuestionController{service: service}
}

func (qc *QuestionController) Register(router gin.IRouter) {
	QuestionRouter := router.Group("/examservice/questions")
	QuestionRouter.POST("/", qc.CreateQuestions)
	QuestionRouter.GET("/", qc.GetQuestionsList)
}

// @Summary Create or update questions
// @Description Create or update questions based on the request
// @Tags Questions
// @Accept json
// @Produce json
// @Param request body dto.QuestionRequest true "Question request body"
// @Success 200 {object} dto.QuestionResponse "Successfully created or updated questions"
// @Failure 400 {object} utils.CustomError "Invalid request body"
// @Failure 500 {object} utils.CustomError "Internal server error"
// @Router /examservice/questions [post]
func (qc *QuestionController) CreateQuestions(ctx *gin.Context) {
	var req dto.QuestionRequest
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

	res, err := qc.service.CreateQuestions(ctx, &req)
	if err != nil {
		utils.WriteError(ctx, err)
		return
	}
	utils.WriteResponse(ctx, res)
}

// GetQuestionsList retrieves a list of all questions.
// @Summary Get all questions
// @Description Retrieve a list of all questions.
// @Tags Questions
// @Accept json
// @Produce json
// @Success 200 {array} dto.ListQuestionResponse "Successful response"
// @Failure 400 {object} utils.CustomError "Invalid request"
// @Failure 500 {object} utils.CustomError "Internal server error"
// @Router /examservice/questions/ [get]
func (qc *QuestionController) GetQuestionsList(ctx *gin.Context) {
	configs, err := qc.service.GetQuestionsList(ctx)
	if err != nil {
		utils.WriteError(ctx, err)
		return
	}
	utils.WriteResponse(ctx, configs)
}