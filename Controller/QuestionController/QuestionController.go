package questioncontroller

import (
	"net/http"
	"time"

	question "github.com/KazuwoKiwame12/bookViewerBackend/DB/Model/Question"
	"github.com/labstack/echo/v4"
)

type requestPost struct {
	UserID     int    `json:"user_id"`
	SentenceID int    `json:"sentence_id"`
	Title      string `json:"title"`
	Content    string `json:"content"`
	PageNum    int    `json:"page_num"`
	RowNum     int    `json:"row_num"`
}

//Post ...質問投稿
func Post(c echo.Context) error {
	request := new(requestPost)
	if err := c.Bind(request); err != nil {
		returnValue := map[string]bool{"hasSuccess": false}
		return c.JSON(http.StatusInternalServerError, returnValue)
	}

	var com question.Question
	//構造体に入れる
	com.UserID = request.UserID
	com.SentenceID = request.SentenceID
	com.Title = request.Title
	com.Content = request.Content
	com.Page_num = request.PageNum
	com.Row_num = request.RowNum
	com.Created_At = time.Now()

	//DB処理
	hasSuccess := question.Create(com)
	returnValue := map[string]bool{"hasSuccess": hasSuccess}
	return c.JSON(http.StatusOK, returnValue)
}
