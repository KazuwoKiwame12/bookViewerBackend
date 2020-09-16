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
	model.Content = "Pythonではコードが書きやすいようになっています"
	jst, _ := time.LoadLocation("Asia/Tokyo")
	model.CreatedAt = time.Now().In(jst)
	reply.Create(model)
	model.UserID = 2
	model.Content = "PythonはIoT等にも活用できます"
	reply.Create(model)

	fmt.Println(reply.Get(1))
	fmt.Println(reply.GetList())
}
