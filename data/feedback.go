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

func AddFeedback(body string) (err error) {
	feedback := Feedback{
		Body:      body,
		CreatedAt: time.Now(),
	}
	_, err = db.NamedExec(`INSERT INTO feedbacks (body, created_at) VALUES (:body, :created_at)`, feedback)
	if err != nil {
		warning("Error during logQuery:", err)
	}
	return
}

func GetAllFeedbacks() (feedbacks []Feedback) {
	db.Select(&feedbacks, "SELECT * FROM feedbacks ORDER BY created_at DESC")
	return
}
