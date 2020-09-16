package main

import (
	"fmt"
	"time"

	reply "github.com/KazuwoKiwame12/bookViewerBackend/DB/Model/Reply"
)

func main() {
	model := reply.Reply{}
	model.UserID = 1
	model.QuestionID = 1
	model.Content = "Pythonでは機械学習に特化したライブラリがたくさん存在します"
	jst, _ := time.LoadLocation("Asia/Tokyo")
	model.CreatedAt = time.Now().In(jst)
	reply.Create(model)
	model.UserID = 2
	model.Content = "Pythonは配列操作が柔軟に行うことが出来ます"
	reply.Create(model)

	fmt.Println(reply.Get(1))
	fmt.Println(reply.GetList())
}
