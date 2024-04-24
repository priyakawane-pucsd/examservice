package exam

import (
	"context"
	"examservice/models/dto"
	"examservice/utils"

	"github.com/bappaapp/goutils/logger"
	"github.com/gin-gonic/gin"
)

type ExamController struct {
	service Service
}

type Service interface {
	CreateOrUpdateExam(ctx context.Context, req *dto.ExamRequest) (*dto.ExamResponse, error)
	GetExamsList(ctx context.Context, topic string, subTopic string) (*dto.ListExamsResponse, error)
	GetExamById(ctx context.Context, examId string) (*dto.ExamByIdResponse, error)
	DeleteExamById(ctx context.Context, examId string) (*dto.DeleteExamResponse, error)
}

func NewExamController(ctx context.Context, service Service) *ExamController {
	return &ExamController{service: service}
}

func (ec *ExamController) Register(router gin.IRouter) {
	examRouter := router.Group("/examservice/exams")
	examRouter.POST("/", ec.CreateOrUpdateExam)
	examRouter.GET("/", ec.GetExamsList)
	examRouter.GET("/:id", ec.GetExamById)
	examRouter.DELETE("/:id", ec.DeleteExamById)
}

// CreateOrUpdateExam creates a new exam or updates an existing one.
// If the exam ID is provided in the request body, it will update the existing exam.
// Otherwise, it will create a new exam based on the provided request body.
// @Summary Create or update an exam
// @Description Create a new exam or update an existing one based on the provided request body.
// @Tags Exams
// @Accept json
// @Produce json
// @Param request body dto.ExamRequest true "Exam request body"
// @Success 200 {object} dto.ExamResponse "Successful response"
// @Failure 400 {object} utils.CustomError "Invalid request body"
// @Failure 500 {object} utils.CustomError "Internal server error"
// @Router /examservice/exams [post]
func (ec *ExamController) CreateOrUpdateExam(ctx *gin.Context) {
	var req dto.ExamRequest
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

	res, err := ec.service.CreateOrUpdateExam(ctx, &req)
	if err != nil {
		utils.WriteError(ctx, err)
		return
	}
	utils.WriteResponse(ctx, res)
}

// GetExamsList retrieves a list of exams filtered by topic and subTopic.
// @Summary Get all exams
// @Description Retrieve a list of all exams filtered by topic and subTopic.
// @Tags Exams
// @Accept json
// @Produce json
// @Param topic query string false "Filter by topic"
// @Param subTopic query string false "Filter by subTopic"
// @Success 200 {array} dto.ListExamsResponse "Successful response"
// @Failure 400 {object} utils.CustomError "Invalid request"
// @Failure 500 {object} utils.CustomError "Internal server error"
// @Router /examservice/exams/ [get]
func (ec *ExamController) GetExamsList(ctx *gin.Context) {
	// Get the topic and subTopic query parameters
	topic := ctx.Query("topic")
	subTopic := ctx.Query("subTopic")

	// Call the service function with optional filter parameters
	exams, err := ec.service.GetExamsList(ctx, topic, subTopic)
	if err != nil {
		utils.WriteError(ctx, err)
		return
	}
	utils.WriteResponse(ctx, exams)
}

// @Summary Get an exam by ID
// @Description Retrieve an exam based on the provided exam ID.
// @Tags Exams
// @Accept json
// @Produce json
// @Param id path string true "Exam ID"
// @Success 200 {object} dto.Exam "Successful response"
// @Failure 400 {object} utils.CustomError "Invalid request"
// @Failure 404 {object} utils.CustomError "Exam not found"
// @Failure 500 {object} utils.CustomError "Internal server error"
// @Router /examservice/exams/{id} [get]
func (ec *ExamController) GetExamById(ctx *gin.Context) {
	// Extract the exam ID from the request context
	examId := ctx.Param("id")

	res, err := ec.service.GetExamById(ctx, examId)
	if err != nil {
		utils.WriteError(ctx, err)
		return
	}
	utils.WriteResponse(ctx, res)
}

// DeleteExamById deletes an exam by its ID.
// @Summary Delete an exam by ID
// @Description Deletes an exam by its ID.
// @Tags Exams
// @Param id path string true "Exam ID"
// @Accept json
// @Produce json
// @Success 200 {object} dto.ExamResponse "Successful response"
// @Failure 400 {object} utils.CustomError "Invalid request"
// @Failure 404 {object} utils.CustomError "Exam not found"
// @Failure 500 {object} utils.CustomError "Internal server error"
// @Router /examservice/exams/{id} [delete]
func (ec *ExamController) DeleteExamById(ctx *gin.Context) {
	examId := ctx.Param("id")

	// Call the service function to delete the exam by ID
	res, err := ec.service.DeleteExamById(ctx, examId)
	if err != nil {
		utils.WriteError(ctx, err)
		return
	}
	utils.WriteResponse(ctx, res)
}
