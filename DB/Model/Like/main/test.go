package main

import (
	"fmt"
	like "github.com/KazuwoKiwame12/bookViewerBackend/DB/Model/Like"
)

func main() {
	// ユーザーモデルの作成
	model := like.Like{}
	model.UserID = 1
	model.ReplyID = 1
	like.Create(model)
	model.UserID = 2
	model.ReplyID = 2
	like.Create(model)
	model.UserID = 2
	model.ReplyID = 2
	like.Create(model)

	fmt.Println(like.Get(1))
	fmt.Println(like.GetListByChapter(1))
	fmt.Println(like.GetListByUser(1))
}
