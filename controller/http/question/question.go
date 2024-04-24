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
	CreateOrUpdateQuestions(ctx context.Context, req *dto.QuestionRequest) (*dto.QuestionResponse, error)

	//todo use filters to fetch questions
	GetQuestionsList(ctx context.Context) (*dto.ListQuestionResponse, error)
	DeleteQuestionById(ctx context.Context, questionId string) (*dto.DeleteQuestionResponse, error)
	GetQuestionById(ctx context.Context, questionId string) (*dto.QuestionByIdResponse, error)
}

func NewQuestionController(ctx context.Context, service Service) *QuestionController {
	return &QuestionController{service: service}
}

func (qc *QuestionController) Register(router gin.IRouter) {
	QuestionRouter := router.Group("/examservice/questions")
	QuestionRouter.POST("/", qc.CreateOrUpdateQuestions)
	QuestionRouter.GET("/", qc.GetQuestionsList)
	QuestionRouter.GET("/:id", qc.GetQuestionById)
	QuestionRouter.DELETE("/:id", qc.DeleteQuestionById)
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
func (qc *QuestionController) CreateOrUpdateQuestions(ctx *gin.Context) {
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

	res, err := qc.service.CreateOrUpdateQuestions(ctx, &req)
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

// GetQuestionById retrieves a question by its ID.
// @Summary Retrieve a question by ID
// @Description Retrieves a question based on the provided ID.
// @Tags Questions
// @Accept json
// @Produce json
// @Param id path string true "Question ID"
// @Success 200 {object} dto.QuestionByIdResponse "Successful response"
// @Failure 400 {object} utils.CustomError "Invalid request"
// @Failure 404 {object} utils.CustomError "Question not found"
// @Failure 500 {object} utils.CustomError "Internal server error"
// @Router /examservice/questions/{id} [get]
func (qc *QuestionController) GetQuestionById(ctx *gin.Context) {
	questionId := ctx.Param("id")

	// Call the service function to delete the question by ID
	res, err := qc.service.GetQuestionById(ctx, questionId)
	if err != nil {
		utils.WriteError(ctx, err)
		return
	}
	utils.WriteResponse(ctx, res)
}

// DeleteQuestionById deletes a question by its ID.
// @Summary Delete a question by ID
// @Description Deletes a question by its ID.
// @Tags Questions
// @Param id path string true "Question ID"
// @Accept json
// @Produce json
// @Success 200 {object} dto.QuestionResponse "Successful response"
// @Failure 400 {object} utils.CustomError "Invalid request"
// @Failure 404 {object} utils.CustomError "Question not found"
// @Failure 500 {object} utils.CustomError "Internal server error"
// @Router /examservice/questions/{id} [delete]
func (qc *QuestionController) DeleteQuestionById(ctx *gin.Context) {
	// Extract the question ID from the request context
	questionId := ctx.Param("id")

	// Call the service function to delete the question by ID
	res, err := qc.service.DeleteQuestionById(ctx, questionId)
	if err != nil {
		utils.WriteError(ctx, err)
		return
	}
	utils.WriteResponse(ctx, res)
}
