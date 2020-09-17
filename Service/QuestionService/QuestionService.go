package questionservice

import (
	question "github.com/KazuwoKiwame12/bookViewerBackend/DB/Model/Question"
	user "github.com/KazuwoKiwame12/bookViewerBackend/DB/Model/User"
)

type questionFC struct {
	QuestionID int    `json:"questionId"`
	Title      string `jsond:"title"`
	UserName   string `json:"userName"`
	CreatedAt  string `json:"createdAt"`
}

//QuestionListFC ..クライアントに返す質問一覧
type QuestionListFC struct {
	Questions []questionFC `json:"questions"`
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
		qFC.CreatedAt = q.CreatedAt.Format("2006-01-02 15:04:05")

		questionListFC = append(questionListFC, qFC)
	}
	return QuestionListFC{Questions: questionListFC}
}
