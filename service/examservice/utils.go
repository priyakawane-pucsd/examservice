package examservice

import (
	"examservice/models/dao"
	"examservice/models/dto"
)

func ConvertToExamResponseList(exams []*dao.Exam) []dto.Exam {
	var convertedExams []dto.Exam
	for _, exam := range exams {
		convertedExam := ExamsResponse(exam)
		convertedExams = append(convertedExams, *convertedExam)
	}
	return convertedExams
}

func ExamsResponse(exam *dao.Exam) *dto.Exam {
	return &dto.Exam{
		ID:              exam.ID,
		Title:           exam.Title,
		Description:     exam.Description,
		StartTime:       exam.StartTime,
		EndTime:         exam.EndTime,
		Duration:        exam.Duration,
		Questions:       exam.Questions,
		Topic:           exam.Topic,
		ExamFee:         exam.ExamFee,
		DifficultyLevel: exam.DifficultyLevel,
		SubTopic:        exam.SubTopic,
		CreatedAt:       exam.CreatedAt,
		UpdatedAt:       exam.UpdatedAt,
	}
}
