package main

import (
	"fmt"

	book "github.com/KazuwoKiwame12/bookViewerBackend/DB/Model/Book"
)

func main() {
	model := book.Book{}
	model.GenreID = 1
	model.Title = "Python入門"
	model.Image = "hogehoge"
	model.Price = 2000
	model.Author = "Me"
	model.Bio = "この本は初心者のための本です"
	book.Create(model)
	fmt.Println(book.Get(1))
	fmt.Println(book.GetListByGenre(1))
}
