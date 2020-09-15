package genre

import (
	db "github.com/KazuwoKiwame12/bookViewerBackend/DB"
)

type book struct {
	ID      int
	GenreID int
	Title   string
	Image   string
	Price   int
	Author  string
	Bio     string
}

//Genre ...ジャンル
type Genre struct {
	ID    int
	Name  string
	Books []book
}

//Create ...ジャンルモデルの新規作成
func Create(genre Genre) bool {
	db := db.Connect()
	defer db.Close()

	result := db.Create(&genre)
	if result.Error != nil {
		return false
	}
	return true
}

//Get ...ジャンルモデルの取得
func Get(id int) Genre {
	db := db.Connect()
	defer db.Close()

	genre := Genre{}
	db.First(&genre, id)

	return genre
}

//GetList ...全てのジャンルモデルの取得
func GetList() []Genre {
	db := db.Connect()
	defer db.Close()

	genreList := []Genre{}
	db.Find(&genreList)

	return genreList
}
