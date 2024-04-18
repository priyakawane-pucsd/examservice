package dao

type Exam struct {
	ID          string   `bson:"_id,omitempty"`
	Title       string   `bson:"title"`
	Description string   `bson:"description"`
	StartTime   string   `bson:"start_time"`
	EndTime     string   `bson:"end_time"`
	Duration    int      `bson:"duration"`
	Questions   []string `bson:"questions"`
	CreatedAt   string   `bson:"created_at,omitempty"`
	UpdatedAt   string   `bson:"updated_at,omitempty"`
}
