package data

import "time"

type Feedback struct {
	Id        int
	Body      string
	CreatedAt time.Time `db:"created_at"`
}

func (feedback *Feedback) Date() string {
	return feedback.CreatedAt.Format("06/01/02 3:04pm")
}

func GetAllFeedbacks() (feedbacks []Feedback) {
	db.Select(&feedbacks, "SELECT * FROM feedbacks ORDER BY created_at DESC")
	return
}
