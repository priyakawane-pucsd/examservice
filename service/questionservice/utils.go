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
	var choices []dto.Choice
	for _, choice := range question.Choices {
		choices = append(choices, dto.Choice{
			Key:   choice.Key,
			Value: choice.Value,
		})
	}
	return &dto.Question{
		ID:          question.ID,
		Text:        question.Text,
		Choices:     choices,
		Correct:     question.Correct,
		Explanation: question.Explanation,
		CreatedBy:   question.CreatedBy,
		Topic:       question.Topic,
		SubTopic:    question.SubTopic,
		CreatedAt:   question.CreatedAt,
		UpdatedAt:   question.UpdatedAt,
	}
}
