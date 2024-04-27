package dao

type Exam struct {
	ID              string   `bson:"_id,omitempty"`
	Title           string   `bson:"title"`
	Description     string   `bson:"description"`
	StartTime       int64    `bson:"startTime"`
	EndTime         int64    `bson:"endTime"`
	Duration        int      `bson:"duration"`
	Questions       []string `bson:"questions"`
	Topic           string   `bson:"topic"`
	SubTopic        string   `bson:"subTopic"`
	ExamFee         float64  `bson:"examFee"`
	DifficultyLevel string   `bson:"difficultyLevel"`
	CreatedAt       int64    `bson:"createdAt,omitempty"`
	UpdatedAt       int64    `bson:"updatedAt,omitempty"`
}
