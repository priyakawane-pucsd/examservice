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
	CreateExam(ctx context.Context, req *dto.ExamRequest) (*dto.ExamResponse, error)
	GetExamsList(ctx context.Context) (*dto.ListExamsResponse, error)
}

func NewExamController(ctx context.Context, service Service) *ExamController {
	return &ExamController{service: service}
}

func (ec *ExamController) Register(router gin.IRouter) {
	examRouter := router.Group("/examservice/exams")
	examRouter.POST("/", ec.CreateExam)
	examRouter.GET("/", ec.GetExamsList)
}

// CreateExam creates a new exam.
// @Summary Create a new exam
// @Description Create a new exam based on the provided request body.
// @Tags Exams
// @Accept json
// @Produce json
// @Param request body dto.ExamRequest true "Exam request body"
// @Success 200 {object} dto.ExamResponse "Successful response"
// @Failure 400 {object} utils.CustomError "Invalid request body"
// @Failure 500 {object} utils.CustomError "Internal server error"
// @Router /examservice/exams [post]
func (ec *ExamController) CreateExam(ctx *gin.Context) {
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

	res, err := ec.service.CreateExam(ctx, &req)
	if err != nil {
		utils.WriteError(ctx, err)
		return
	}
	utils.WriteResponse(ctx, res)
}

// GetExamsList retrieves a list of exams.
// @Summary Get all exams
// @Description Retrieve a list of all exams.
// @Tags Exams
// @Accept json
// @Produce json
// @Success 200 {array} dto.ListExamsResponse "Successful response"
// @Failure 400 {object} utils.CustomError "Invalid request"
// @Failure 500 {object} utils.CustomError "Internal server error"
// @Router /examservice/exams/ [get]
func (ec *ExamController) GetExamsList(ctx *gin.Context) {
	exams, err := ec.service.GetExamsList(ctx)
	if err != nil {
		utils.WriteError(ctx, err)
		return
	}
	utils.WriteResponse(ctx, exams)
}
