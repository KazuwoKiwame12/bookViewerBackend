package main

import (
	"fmt"
	"time"

	question "github.com/KazuwoKiwame12/bookViewerBackend/DB/Model/Question"
)

func main() {
	model := question.Question{}
	model.UserID = 1
	model.SentenceID = 2
	model.Title = "Pythonとは何だね"
	model.Content = "何に特化してるの?"
	jst, _ := time.LoadLocation("Asia/Tokyo")
	model.CreatedAt = time.Now().In(jst)
	question.Create(model)
	model.Title = "このコードについて"
	model.Content = "具体的にどのような処理をおこなっているの?"
	question.Create(model)
	fmt.Println(question.Show(1))
	fmt.Println(question.GetQuestionList())
}
