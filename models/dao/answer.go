package dao

type Answer struct {
	ID        string           `bson:"_id,omitempty"`
	UserID    string           `bson:"userId"`
	ExamID    string           `bson:"examId"`
	Answers   []QuestionAnswer `bson:"answers"`
	CreatedAt int64            `bson:"createdAt"`
	UpdatedAt int64            `bson:"updatedAt"`
}

type QuestionAnswer struct {
	QuestionId string `bson:"questionId"`
	Answer     string `bson:"answer"`
}
