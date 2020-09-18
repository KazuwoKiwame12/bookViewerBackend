package question

import (
	chapter "github.com/KazuwoKiwame12/bookViewerBackend/DB/Model/Chapter"
	page "github.com/KazuwoKiwame12/bookViewerBackend/DB/Model/Page"
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

func GetListByChapter(chapterID int) []Question {
	db := db.Connect()
	defer db.Close()

	chapter := chapter.Chapter{}
	chapter.ID = chapterID

	// 章のidより、ページidを取得
	page := page.Page{}
	db.First(&page, chapterID)

	//ページから、文章一覧を取得
	sentenceList := []sentence.Sentence{}
	db.Model(&page).Association("Sentences").Find(&sentenceList)

	//文章一覧から、idを抜き取る
	sentenceIds := make([]interface{}, len(sentenceList))
	for i, s := range sentenceList {
		sentenceIds[i] = s.ID
	}
	questionList := []Question{}

	//id一覧で、where("sentence_id in ?", id一覧)で検索し、質問一覧取得
	db.Where("sentence_id in (?)", sentenceIds).Find(&questionList)

	return questionList
}

//GetListByTitle ...送られてきたタイトルと部分的にでも合致する質問リストを返す
func GetListByTitle(title string) []Question {
	db := db.Connect()
	defer db.Close()

	questionList := []Question{}
	db.Where("Title LIKE ?", "%"+title+"%").Find(&questionList)

	return questionList
}
