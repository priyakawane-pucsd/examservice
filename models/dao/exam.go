package dao

type Exam struct {
	ID              string   `bson:"_id,omitempty"`
	Title           string   `bson:"title"`
	Description     string   `bson:"description"`
	StartTime       int64    `bson:"start_time"`
	EndTime         int64    `bson:"end_time"`
	Duration        int      `bson:"duration"`
	Questions       []string `bson:"questions"`
	Topic           string   `bson:"topic"`
	SubTopic        string   `bson:"sub_topic"`
	ExamFee         float64  `bson:"exam_fee"`
	DifficultyLevel string   `bson:"difficulty_level"`
	CreatedAt       int64    `bson:"created_at,omitempty"`
	UpdatedAt       int64    `bson:"updated_at,omitempty"`
}
