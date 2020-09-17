package replycontroller

import (
	"net/http"
	"time"

	reply "github.com/KazuwoKiwame12/bookViewerBackend/DB/Model/Reply"
	"github.com/labstack/echo/v4"
)

type requestPost struct {
	UserID     int    `json:"userId"`
	QuestionID int    `json:"questionId"`
	Content    string `json:"content"`
}

//Post ...質問回答投稿
func Post(c echo.Context) error {
	request := new(requestPost)
	if err := c.Bind(request); err != nil {
		returnValue := map[string]bool{"hasSuccess": false}
		return c.JSON(http.StatusInternalServerError, returnValue)
	}

	var rep reply.Reply
	//構造体に入れる
	rep.UserID = request.UserID
	rep.QuestionID = request.QuestionID
	rep.Content = request.Content
	jst, _ := time.LoadLocation("Asia/Tokyo")
	rep.CreatedAt = time.Now().In(jst)

	//DB処理
	hasSuccess := reply.Create(rep)
	returnValue := map[string]bool{"hasSuccess": hasSuccess}
	return c.JSON(http.StatusOK, returnValue)
}
