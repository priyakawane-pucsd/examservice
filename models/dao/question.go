package dao

type QuestionType string

type AnswerType string

const (
	MULTIPLE_CHOICE QuestionType = "MULTIPLE_CHOICE"
	MULTILINE       QuestionType = "MULTILINE"
	TRUE_OR_FALSE   QuestionType = "TRUE_OR_FALSE"
)

type Question struct {
	ID   string       `json:"id"`
	Type QuestionType `json:"type"`
}
