package dao

import "time"

type Exam struct {
	ID          string    `bson:"_id,omitempty"`
	Title       string    `bson:"title"`
	Description string    `bson:"description"`
	StartTime   string    `bson:"start_time"`
	EndTime     string    `bson:"end_time"`
	Duration    int       `bson:"duration"`
	Questions   []string  `bson:"questions"`
	Topic       string    `bson:"topic"`
	SubTopic    string    `bson:"sub_topic"`
	CreatedAt   time.Time `bson:"created_at,omitempty"`
	UpdatedAt   time.Time `bson:"updated_at,omitempty"`
}
