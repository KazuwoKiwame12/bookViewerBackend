package Like

import (
	db "github.com/KazuwoKiwame12/bookViewerBackend/DB"
	chapter "github.com/KazuwoKiwame12/bookViewerBackend/DB/Model/Reply"
	user "github.com/KazuwoKiwame12/bookViewerBackend/DB/Model/User"
)

// Like ...いいねのモデル
type Like struct {
	ID      int
	UserID  int
	ReplyID int
}

//Create ...いいねモデルの新規作成
func Create(like Like) bool {
	db := db.Connect()
	defer db.Close()

	result := db.Create(&like)
	if result.Error != nil {
		return false
	}
	return true
}

//Get ...いいねモデルの取得
func Get(id int) Like {
	db := db.Connect()
	defer db.Close()

	like := Like{}
	db.First(&like, id)

	return like
}

//Delete ...いいねモデルの削除
func Delete(id int) bool {
	db := db.Connect()
	defer db.Close()

	like := Like{}
	like.ID = id
	result := db.Delete(&like)
	if result.Error != nil {
		return false
	}
	return true
}

//GetListByReply ...指定した返信モデルに属するいいねモデルのリスト取得
func GetListByReply(replyID int) []Like {
	db := db.Connect()
	defer db.Close()

	reply := reply.Reply{}
	reply.ID = replyID

	likeList := []Like{}
	db.Model(&reply).Association("Likes").Find(&likeList)

	return likeList
}

//GetListByUser ...指定したユーザーモデルに属するいいねモデルのリスト取得
func GetListByUser(userID int) []Like {
	db := db.Connect()
	defer db.Close()

	user := user.User{}
	user.ID = userID

	likeList := []Like{}
	db.Model(&user).Association("Likes").Find(&likeList)

	return likeList
}
