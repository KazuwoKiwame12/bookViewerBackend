package questionservice

import (
	"time"

	question "github.com/KazuwoKiwame12/bookViewerBackend/DB/Model/Question"
	user "github.com/KazuwoKiwame12/bookViewerBackend/DB/Model/User"
)

type questionFC struct {
	QuestionID int       `json:"question_id"`
	Title      string    `json:"title"`
	UserName   string    `json:"user_nama"`
	CreatedAt  time.Time `json:"created_at"`
}

//QuestionListFC ..クライアントに返す質問一覧
type QuestionListFC struct {
	Questions []questionFC
}

//GetListFC ...タイトルと部分的にも合致する質問一覧を取得
func GetListFC(questionList []question.Question) QuestionListFC {
	questionListFC := []questionFC{}
	for _, q := range questionList {
		user := user.Get(q.UserID)
		qFC := questionFC{}
		qFC.QuestionID = q.ID
		qFC.UserName = user.Name
		qFC.Title = q.Title
		qFC.CreatedAt = q.CreatedAt

		questionListFC = append(questionListFC, qFC)
	}
	return QuestionListFC{Questions: questionListFC}
}
