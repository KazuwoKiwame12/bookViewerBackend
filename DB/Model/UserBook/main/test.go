package main

import (
	"fmt"
	"time"

	userbook "github.com/KazuwoKiwame12/bookViewerBackend/DB/Model/UserBook"
)

func main() {
	model := userbook.UserBook{}
	model.UserID = 1
	model.BookID = 1
	model.Status = 0
	jst, _ := time.LoadLocation("Asia/Tokyo")
	model.CreatedAt = time.Now().In(jst)
	userbook.Create(model)
	fmt.Println(userbook.Get(1, 1))
	fmt.Println(userbook.GetList(1))

	userbook.UpdateStatus(1, 1, 1)
}
