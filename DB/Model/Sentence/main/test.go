package main

import (
	"fmt"

	sentence "github.com/KazuwoKiwame12/bookViewerBackend/DB/Model/Sentence"
)

func main() {
	model := sentence.Sentence{}
	model.PageID = 1
	model.Content = "文章が入ります"
	sentence.Create(model)

	fmt.Println(sentence.Get(1))
	fmt.Println(sentence.GetListByPage(1))
}
