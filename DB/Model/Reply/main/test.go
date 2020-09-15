package main

import (
	"fmt"
	reply "github.com/KazuwoKiwame12/bookViewerBackend/DB/Model/Reply"
	"time"
)

func main() {
	model := reply.Reply{}
	model.UserID = 1
	model.QuestionID = 1
	model.Content = "heyhey"
	jst, _ := time.LoadLocation("Asia/Tokyo")
	model.CreatedAt = time.Now().In(jst)
	reply.Create(model)

	fmt.Println(reply.Get(1))
	fmt.Println(reply.GetList())
}
