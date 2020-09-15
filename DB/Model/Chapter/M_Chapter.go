package chapter

import (
	db "github.com/KazuwoKiwame12/bookViewerBackend/DB"
	book "github.com/KazuwoKiwame12/bookViewerBackend/DB/Model/Book"
)

type page struct {
	ID        int
	ChapterID int
	Number    int
}

//Chapter ...章のモデル
type Chapter struct {
	ID     int
	BookID int
	Number int
	Pages  []page
}

//Create ...章モデルの新規作成
func Create(chapter Chapter) bool {
	db := db.Connect()
	defer db.Close()

	result := db.Create(&chapter)
	if result.Error != nil {
		return false
	}
	return true
}

//Get ...章モデルの取得
func Get(id int) Chapter {
	db := db.Connect()
	defer db.Close()

	chapter := Chapter{}
	db.First(&chapter, id)

	return chapter
}

//GetListByBook ...指定した本モデルに属する章モデルのリスト取得
func GetListByBook(bookID int) []Chapter {
	db := db.Connect()
	defer db.Close()

	book := book.Book{}
	book.ID = bookID

	chapterList := []Chapter{}
	db.Model(&book).Association("Chapters").Find(&chapterList)

	return chapterList
}
