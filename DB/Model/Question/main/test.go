package main

import (
	"fmt"

	question "github.com/KazuwoKiwame12/bookViewerBackend/DB/Model/Question"
)

func main() {
	model := question.Question{}
	model.UserID = 1
	model.SentenceID = 2
	model.Title = "Pythonについて"
	model.Content = "魅力を教えて～"
	question.Create(model)
	fmt.Println(question.Show(1))
	fmt.Println(question.GetQuestionList())
}
