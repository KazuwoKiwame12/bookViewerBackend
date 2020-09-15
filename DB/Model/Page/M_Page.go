package page

import (
	db "github.com/KazuwoKiwame12/bookViewerBackend/DB"
	chapter "github.com/KazuwoKiwame12/bookViewerBackend/DB/Model/Chapter"
)

type sentence struct {
	ID      int
	PageID  int
	Content string
}

//Page ...ページのモデル
type Page struct {
	ID        int
	ChapterID int
	Number    int
	Sentences []sentence
}

//Create ...ページモデルの新規作成
func Create(page Page) bool {
	db := db.Connect()
	defer db.Close()

	result := db.Create(&page)
	if result.Error != nil {
		return false
	}
	return true
}

//Get ...ページモデルの取得
func Get(id int) Page {
	db := db.Connect()
	defer db.Close()

	page := Page{}
	db.First(&page, id)

	return page
}

//GetListByChapter ...指定した章モデルに属するページモデルのリスト取得
func GetListByChapter(chapterID int) []Page {
	db := db.Connect()
	defer db.Close()

	chapter := chapter.Chapter{}
	chapter.ID = chapterID

	pageList := []Page{}
	db.Model(&chapter).Association("Pages").Find(&pageList)

	return pageList
}
