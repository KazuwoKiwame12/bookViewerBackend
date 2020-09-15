package main

import (
	"fmt"

	page "github.com/KazuwoKiwame12/bookViewerBackend/DB/Model/Page"
)

func main() {
	model := page.Page{}
	model.ChapterID = 1
	model.Number = 1
	page.Create(model)
	model.Number = 2
	page.Create(model)
	model.Number = 3
	page.Create(model)

	fmt.Println(page.Get(1))
	fmt.Println(page.GetListByChapter(1))
}
