package questionservice

import (
	"examservice/models/dao"
	"examservice/models/dto"
)

func ConvertToQuestionResponseList(questions []*dao.Question) []dto.Question {
	var convertedQuestions []dto.Question
	for _, question := range questions {
		convertedQue := QuestionsResponse(question)
		convertedQuestions = append(convertedQuestions, *convertedQue)
	}
	return convertedQuestions
}

func QuestionsResponse(question *dao.Question) *dto.Question {
	return &dto.Question{
		ID:          question.ID,
		Text:        question.Text,
		Choices:     question.Choices,
		Correct:     question.Correct,
		Explanation: question.Explanation,
		UserId:      question.UserId,
		CreatedAt:   question.CreatedAt,
		UpdatedAt:   question.UpdatedAt,
	}
}
