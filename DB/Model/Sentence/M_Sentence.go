package sentence

import (
	db "github.com/KazuwoKiwame12/bookViewerBackend/DB"
	page "github.com/KazuwoKiwame12/bookViewerBackend/DB/Model/Page"
)

//Sentence ...センテンスモデル
type Sentence struct {
	ID      int
	PageID  int
	Content string
}

//Create ...センテンスの新規作成
func Create(sentence Sentence) bool {
	db := db.Connect()
	defer db.Close()

	result := db.Create(&sentence)
	if result.Error != nil {
		return false
	}
	return true
}

//Get ...センテンスの取得
func Get(sentenceID int) Sentence {
	db := db.Connect()
	defer db.Close()

	sentence := Sentence{}
	db.First(&sentence, sentenceID)

	return sentence
}

//GetListByPage ...同じページに属するセンテンスリストを取得
func GetListByPage(pageID int) []Sentence {
	db := db.Connect()
	defer db.Close()

	page := page.Page{}
	page.ID = pageID

	sentenceList := []Sentence{}
	db.Model(&page).Association("Sentences").Find(&sentenceList)

	return sentenceList
}