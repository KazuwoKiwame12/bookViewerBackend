package questionsearchcontroller

import (
	"net/http"
	"strconv"

	question "github.com/KazuwoKiwame12/bookViewerBackend/DB/Model/Question"
	questionservice "github.com/KazuwoKiwame12/bookViewerBackend/Service/QuestionService"
	"github.com/labstack/echo/v4"
)

//GetListByTitle ...タイトルと部分的にも合致する質問一覧を取得
func GetListByTitle(c echo.Context) error {
	title := c.Param("title")
	//TODO 取得する質問は章を絞れるようにする
	response := questionservice.GetListFC(question.GetListByTitle(title))
	return c.JSON(http.StatusOK, response)
}

//GetListBySentence ...文章ごとの質問一覧を取得
func GetListBySentence(c echo.Context) error {
	sentenceID, _ := strconv.Atoi(c.Param("id"))
	//TODO 取得する質問は章を絞れるようにする
	response := questionservice.GetListFC(question.GetListBySentence(sentenceID))
	return c.JSON(http.StatusOK, response)
}
