package questioncontentcontroller

import (
	"net/http"
	"strconv"

	question "github.com/KazuwoKiwame12/bookViewerBackend/DB/Model/Question"
	user "github.com/KazuwoKiwame12/bookViewerBackend/DB/Model/User"
	"github.com/labstack/echo/v4"
)

type questionContent struct {
	UserName  string `json:"userName"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	PageNum   int    `json:"pageNum"`
	RowNum    int    `json:"rowNum"`
	CreatedAt string `json:"createdAt"`
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
	response.CreatedAt = question.CreatedAt.Format("2006-01-02 15:04:05")

	return c.JSON(http.StatusOK, response)
}
