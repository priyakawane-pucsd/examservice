package dao

type QuestionType string

type AnswerType string

const (
	MULTIPLE_CHOICE QuestionType = "MULTIPLE_CHOICE"
	MULTILINE       QuestionType = "MULTILINE"
	TRUE_OR_FALSE   QuestionType = "TRUE_OR_FALSE"
)

type Choice struct {
	Key   string `bson:"key"`
	Value string `bson:"value"`
}

type Question struct {
	ID          string   `bson:"_id,omitempty"`
	Text        string   `bson:"text"`
	Choices     []Choice `bson:"choices"`
	Correct     string   `bson:"correct"`
	Explanation string   `bson:"explanation"`
	Topic       string   `bson:"topic"`
	SubTopic    string   `bson:"subTopic"`
	CreatedBy   int64    `bson:"createdBy"`
	IsDeleted   bool     `bson:"isDeleted"`
	CreatedAt   int64    `bson:"createdAt"`
	UpdatedAt   int64    `bson:"updatedAt"`
}
