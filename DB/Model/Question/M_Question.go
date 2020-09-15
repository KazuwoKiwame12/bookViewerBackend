package question

import (
	"time"

	db "github.com/KazuwoKiwame12/bookViewerBackend/DB"
)

//Question ...質問モデル
type Question struct {
	ID         int
	UserID     int
	SentenceID int
	Title      string
	Content    string
	Page_num   int
	Row_num    int
	Created_At time.Time
}

//Create ...質問投稿
func Create(question Question) bool {
	db := db.Connect()
	defer db.Close()

	result := db.Create(&question)
	if result.Error != nil {
		return false
	}
	return true
}

//GetQuestionList ...質問一覧取得
func GetQuestionList() []Question {
	db := db.Connect()
	defer db.Close()

	questionList := []Question{}
	db.Find(&questionList)

	return questionList
}

//Show ...質問詳細取得
func Show(questionID int) Question {
	db := db.Connect()
	defer db.Close()

	question := Question{}
	db.First(&question, questionID)

	return question
}
