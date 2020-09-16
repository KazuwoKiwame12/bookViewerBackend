package reply

import (
	"time"

	db "github.com/KazuwoKiwame12/bookViewerBackend/DB"
	book "github.com/KazuwoKiwame12/bookViewerBackend/DB/Model/Book"
	chapter "github.com/KazuwoKiwame12/bookViewerBackend/DB/Model/Chapter"
	page "github.com/KazuwoKiwame12/bookViewerBackend/DB/Model/Page"
	question "github.com/KazuwoKiwame12/bookViewerBackend/DB/Model/Question"
	sentence "github.com/KazuwoKiwame12/bookViewerBackend/DB/Model/Sentence"
)

//Reply ...質問に対する返信のモデル
type Reply struct {
	ID         int
	UserID     int
	QuestionID int
	Content    string
	CreatedAt  time.Time
}

//Create ...返信モデルの作成
func Create(reply Reply) bool {
	db := db.Connect()
	defer db.Close()

	result := db.Create(&reply)
	if result.Error != nil {
		return false
	}
	return true
}

//Get ...リプライ詳細の取得
func Get(id int) Reply {
	db := db.Connect()
	defer db.Close()

	reply := Reply{}
	db.First(&reply, id)

	return reply
}

//GetList ...全てのリプライモデルの取得
func GetList() []Reply {
	db := db.Connect()
	defer db.Close()

	replyList := []Reply{}
	db.Find(&replyList)

	return replyList
}

//GetListByReader ...読者のリプライモデル一覧を取得
func GetListByReader(questionID int) []Reply {
	db := db.Connect()
	defer db.Close()

	replyList := []Reply{}
	db.Where("question_id = ?", questionID).Not("user_id = ?", getAuthorID(questionID)).Find(&replyList)
	return replyList
}

//GetListByAuthor ...著者のリプライモデル一覧を取得
func GetListByAuthor(questionID int) []Reply {
	db := db.Connect()
	defer db.Close()

	replyList := []Reply{}
	db.Where("user_id = ? AND question_id = ?", getAuthorID(questionID), questionID).Find(&replyList)
	return replyList
}

func getAuthorID(questionID int) int {
	q := question.Show(questionID)
	s := sentence.Get(q.SentenceID)
	p := page.Get(s.PageID)
	c := chapter.Get(p.ChapterID)
	b := book.Get(c.BookID)
	return b.AuthorID
}
