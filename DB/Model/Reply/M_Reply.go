package Reply

import (
	db "github.com/KazuwoKiwame12/bookViewerBackend/DB"
	"time"
)

type like struct {
	ID      int
	UserID  int
	ReplyID int
}

type Reply struct {
	ID         int
	UserID     int
	QuestionID int
	Content    string
	CreatedAt  time.Time
	Likes      []like
}

func Create(reply Reply) bool {
	db := db.Connect()
	defer db.Close()

	result := db.Create(&reply)
	if result.Error != nil {
		return false
	}
	return true
}

//Get ...リプライモデルの取得
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
