package sentencecontroller

import (
	"net/http"
	"strconv"

	sentence "github.com/KazuwoKiwame12/bookViewerBackend/DB/Model/Sentence"
	"github.com/labstack/echo/v4"
)

type questionID struct {
	QuestionID int
}

type contentData struct {
	Content string `json:"content"`
}

//GetSentence ...質問に対応するページのセンテンスを取得
func GetSentence(c echo.Context) error {
	questionID, _ := strconv.Atoi(c.Param("id"))
	sentenceList := sentence.GetListByQuestion(questionID)

	contentDataList := []contentData{}
	for _, sentence := range sentenceList {
		contentData := contentData{}
		contentData.Content = sentence.Content

		contentDataList = append(contentDataList, contentData)
	}

	response := map[string][]contentData{"content_list": contentDataList}
	return c.JSON(http.StatusOK, response)
}
