package ChapterController

import (
	"net/http"

	question "github.com/KazuwoKiwame12/bookViewerBackend/DB/Model/Question"
	user "github.com/KazuwoKiwame12/bookViewerBackend/DB/Model/User"
	"github.com/labstack/echo/v4"
)

type questionsData struct {
	QuestionID int
	Title      string
	UserName   string
	CreatedAt  string
}
type QuestionsList struct {
	Questions []questionsData
}

// 質問一覧を取得
func GetQuestionList(c echo.Context) error {
	response := QuestionsList{}

	questionsDataList := []questionsData{}
	for _, q := range question.GetQuestionList() {
		user := user.Get(q.UserID)
		questionsData := questionsData{}
		questionsData.QuestionID = q.ID
		questionsData.UserName = user.Name
		questionsData.Title = q.Title
		questionsData.CreatedAt = q.CreatedAt.Format("2006-01-02 15:04:05")

		questionsDataList = append(questionsDataList, questionsData)
	}
	response.Questions = questionsDataList
	return c.JSON(http.StatusOK, response)
}
