package question

import (
	"time"

	db "github.com/KazuwoKiwame12/bookViewerBackend/DB"
	sentence "github.com/KazuwoKiwame12/bookViewerBackend/DB/Model/Sentence"
)

//Question ...質問モデル
type Question struct {
	ID         int
	UserID     int
	SentenceID int
	Title      string
	Content    string
	PageNum    int
	RowNum     int
	CreatedAt  time.Time
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

//GetListBySentence ...同じ文章に属する質問リストを取得
func GetListBySentence(sentenceID int) []Question {
	db := db.Connect()
	defer db.Close()

	sentence := sentence.Sentence{}
	sentence.ID = sentenceID

	questionList := []Question{}
	db.Model(&sentence).Association("Questions").Find(&questionList)

	return questionList
}
