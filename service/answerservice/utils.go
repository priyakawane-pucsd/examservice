package answerservice

import (
	"examservice/models/dao"
	"examservice/models/dto"
)

func ConvertToDAOQuestionAnswer(answer *dto.QuestionAnswer) *dao.QuestionAnswer {
	return &dao.QuestionAnswer{
		QuestionId: answer.QuestionId,
		Answer:     answer.Answer,
	}
}
