package dao

type Answer struct {
	ID        string           `bson:"_id,omitempty"`
	UserID    int64            `bson:"userId"`
	ExamID    string           `bson:"examId"`
	Answers   []QuestionAnswer `bson:"answers"`
	Result    Result           `bson:"result"`
	CreatedAt int64            `bson:"createdAt"`
	UpdatedAt int64            `bson:"updatedAt"`
}

type Result struct {
	Attempted int64 `bson:"attempted"`
	Correct   int64 `bson:"correct"`
}

type QuestionAnswer struct {
	QuestionId string `bson:"questionId"`
	Answer     string `bson:"answer"`
}
