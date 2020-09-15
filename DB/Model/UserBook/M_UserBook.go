package userbook

import (
	"time"

	db "github.com/KazuwoKiwame12/bookViewerBackend/DB"
)

//UserBook ..ユーザと本の中間モデル
type UserBook struct {
	ID        int
	UserID    int
	BookID    int
	Status    int
	CreatedAt time.Time
}

//Create ...ユーザと本の中間モデルの作成
func Create(userbook UserBook) bool {
	db := db.Connect()
	defer db.Close()

	result := db.Create(&userbook)
	if result.Error != nil {
		return false
	}
	return true
}

//Get ...指定したユーザーと本の中間モデル取得
func Get(userID int, bookID int) UserBook {
	db := db.Connect()
	defer db.Close()

	userbook := UserBook{}
	userbook.UserID = userID
	userbook.BookID = bookID
	db.First(&userbook)

	return userbook
}

//GetList ...指定したユーザーの、ユーザと本の中間モデルの一覧取得
func GetList(userID int) []UserBook {
	db := db.Connect()
	defer db.Close()

	userbookList := []UserBook{}
	db.Where("user_id = ?", userID).Find(&userbookList)

	return userbookList
}

//UpdateStatus ...指定したユーザと本の中間モデルの"status"の更新
func UpdateStatus(userID int, bookID int, status int) bool {
	db := db.Connect()
	defer db.Close()

	userbook := UserBook{}
	result := db.Model(&userbook).Where("user_id = ? AND book_id = ?", userID, bookID).Update("status", status)
	if result.Error != nil {
		return false
	}
	return true
}
