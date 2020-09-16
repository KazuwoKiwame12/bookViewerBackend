package questioncontentcontroller

import (
	"net/http"
	"strconv"
	"time"

	question "github.com/KazuwoKiwame12/bookViewerBackend/DB/Model/Question"
	user "github.com/KazuwoKiwame12/bookViewerBackend/DB/Model/User"
	"github.com/labstack/echo/v4"
)

type questionContent struct {
	UserName  string    `json:"user_name"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	PageNum   int       `json:"page_num"`
	RowNum    int       `json:"row_num"`
	CreatedAt time.Time `json:"created_at"`
}

//GetContent ...質問内容などの取得
func GetContent(c echo.Context) error {
	questionID, _ := strconv.Atoi(c.Param("id"))
	question := question.Show(questionID)

	response := questionContent{}
	response.UserName = user.Get(question.UserID).Name
	response.Title = question.Title
	response.Content = question.Content
	response.PageNum = question.PageNum
	response.RowNum = question.RowNum
	response.CreatedAt = question.CreatedAt

	return c.JSON(http.StatusOK, response)
}
