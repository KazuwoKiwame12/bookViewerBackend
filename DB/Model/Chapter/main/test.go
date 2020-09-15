package main

import (
	"fmt"

	chapter "github.com/KazuwoKiwame12/bookViewerBackend/DB/Model/Chapter"
)

func main() {
	model := chapter.Chapter{}
	model.BookID = 1
	model.Number = 1
	chapter.Create(model)
	fmt.Println(chapter.Get(1))
	fmt.Println(chapter.GetListByBook(1))
}
