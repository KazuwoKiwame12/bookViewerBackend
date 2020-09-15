package book

import (
	db "github.com/KazuwoKiwame12/bookViewerBackend/DB"
	genre "github.com/KazuwoKiwame12/bookViewerBackend/DB/Model/Genre"
)

//Book ...本のモデル
type Book struct {
	ID      int
	GenreID int
	Title   string
	Image   string
	Price   int
	Author  string
	Bio     string
}

//Create ...本モデルの新規作成
func Create(book Book) bool {
	db := db.Connect()
	defer db.Close()

	result := db.Create(&book)
	if result.Error != nil {
		return false
	}
	return true
}

//Get ...本モデルの取得
func Get(id int) Book {
	db := db.Connect()
	defer db.Close()

	book := Book{}
	db.First(&book, id)

	return genre
}

//GetListByGenre ...指定したジャンルに属する本モデルのリスト取得
func GetListByGenre(genreID int) []Book {
	db := db.Connect()
	defer db.Close()

	genre := genre.Genre{}
	genre.ID = genreID

	bookList := []Book{}
	db.Model(&genre).Association("Books").Find(&bookList)

	return bookList
}
