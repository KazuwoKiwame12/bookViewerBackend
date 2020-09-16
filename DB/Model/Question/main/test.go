package main

import (
	"fmt"
	"time"

	question "github.com/KazuwoKiwame12/bookViewerBackend/DB/Model/Question"
)

func main() {
	model := question.Question{}
	model.UserID = 1
	model.SentenceID = 1
	model.Title = "Pythonについて"
	model.Content = "魅力を教えて～"
	jst, _ := time.LoadLocation("Asia/Tokyo")
	model.CreatedAt = time.Now().In(jst)
	question.Create(model)
	fmt.Println(question.Show(1))
	fmt.Println(question.GetQuestionList())
}
