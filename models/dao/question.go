package dao

type QuestionType string

type AnswerType string

const (
	MULTIPLE_CHOICE QuestionType = "MULTIPLE_CHOICE"
	MULTILINE       QuestionType = "MULTILINE"
	TRUE_OR_FALSE   QuestionType = "TRUE_OR_FALSE"
)

type Question struct {
	ID          string   `bson:"_id,omitempty"`
	Text        string   `bson:"text"`
	Choices     []string `bson:"choices"`
	Correct     string   `bson:"correct"`
	Explanation string   `bson:"explanation"`
	UserId      string   `bson:"userId"`
	CreatedAt   string   `bson:"createdAt"`
	UpdatedAt   string   `bson:"updatedAt"`
}
